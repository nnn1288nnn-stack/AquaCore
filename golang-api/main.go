package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/penghu-digital-captain/golang-api/handlers"
	"github.com/penghu-digital-captain/golang-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局資料庫連線
var db *gorm.DB

// 初始化環境變數
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env 文件未找到，使用環境變數")
	}
}

// 初始化資料庫
func initDatabase() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "appuser"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "apppass"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "aquaculture_db"
	}

	// DSN 連接字符串
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	log.Printf("🔗 連接資料庫: %s:%s/%s", dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("資料庫連接失敗: %v", err)
	}

	log.Println("✅ 資料庫連接成功")
	return database, nil
}

// 主函數
func main() {
	// 初始化資料庫
	var err error
	db, err = initDatabase()
	if err != nil {
		log.Fatalf("❌ 資料庫初始化失敗: %v", err)
	}
	log.Printf("✅ 資料庫連接成功，db = %v", db != nil)

	// 自動遷移模型
	log.Println("🔄 開始自動遷移...")
	db.AutoMigrate(&models.User{}, &models.EnvironmentalData{}, &models.Asset{}, &models.Task{}, &models.OperationLog{})
	log.Println("✅ 自動遷移完成")

	// 設置 Gin 引擎
	router := gin.Default()

	// 中間件
	router.Use(corsMiddleware())

	// 健康檢查
	router.GET("/health", getHealth)

	// 根路由
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "澎湖數位老船長 - Golang API",
			"version": "1.0.0",
		})
	})

	// API 路由群組
	api := router.Group("/api")
	log.Printf("📋 註冊 API 路由群組: /api")
	{
		// 日誌服務
		logHandler := handlers.NewLogHandler()
		api.POST("/logs", logHandler.LogRequest)
		api.GET("/logs", logHandler.GetLogs)

		// 儀表板
		api.GET("/dashboard", getDashboard)

		// 環境數據
		log.Printf("🔗 註冊環境數據路由: GET /api/environmental-data")
		api.GET("/environmental-data", getEnvironmentalData)
		api.POST("/environmental-data", createEnvironmentalData)

// ============================================
// 資產/庫存處理
// ============================================

func getAssets(c *gin.Context) {
	var assets []models.Asset
	if err := db.Preload("User").Find(&assets).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "獲取資產數據失敗",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "資產列表",
		"data":    assets,
	})
}

func createAsset(c *gin.Context) {
	var asset models.Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(400, gin.H{
			"message": "請求數據格式錯誤",
			"error":   err.Error(),
		})
		return
	}

	if err := db.Create(&asset).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "創建資產失敗",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "資產已創建",
		"data":    asset,
	})
}

func updateAsset(c *gin.Context) {
	id := c.Param("id")
	var asset models.Asset

	if err := db.First(&asset, id).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "資產不存在",
		})
		return
	}

	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(400, gin.H{
			"message": "請求數據格式錯誤",
			"error":   err.Error(),
		})
		return
	}

	if err := db.Save(&asset).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "更新資產失敗",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "資產已更新",
		"data":    asset,
	})
}

func deleteAsset(c *gin.Context) {
	id := c.Param("id")
	var asset models.Asset

	if err := db.First(&asset, id).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "資產不存在",
		})
		return
	}

	if err := db.Delete(&asset).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "刪除資產失敗",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "資產已刪除",
	})
}

		// 資產/庫存
		api.GET("/assets", getAssets)
		api.POST("/assets", createAsset)
		api.PUT("/assets/:id", updateAsset)
		api.DELETE("/assets/:id", deleteAsset)

		// 用戶管理
		api.GET("/users", getUsers)
		api.POST("/users", createUser)
		api.PUT("/users/:id", updateUser)
		api.DELETE("/users/:id", deleteUser)
	}

	// 啟動伺服器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Golang API 在 http://0.0.0.0:%s 啟動", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ 伺服器啟動失敗: %v", err)
	}
}

// CORS 中間件
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// ============================================
// Handler 函數
// ============================================

// getHealth - 健康檢查
func getHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "healthy",
		"service": "golang-api",
	})
}

// ============================================
// 儀表板
// ============================================

func getDashboard(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "儀表板數據",
		"data": gin.H{
			"status": "OK",
		},
	})
}

// ============================================
// 環境數據處理
// ============================================

func getEnvironmentalData(c *gin.Context) {
	log.Printf("🔍 查詢環境數據... 函數被調用")
	c.JSON(500, gin.H{
		"message": "測試錯誤響應",
		"error":   "函數被調用",
	})
	return

func createEnvironmentalData(c *gin.Context) {
	var data models.EnvironmentalData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"message": "請求數據格式錯誤",
			"error":   err.Error(),
		})
		return
	}

	if err := db.Create(&data).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "創建環境數據失敗",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "環境數據已記錄",
		"data":    data,
	})
}

// ============================================
// 資產/庫存處理
// ============================================

func getAssets(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "資產列表",
		"data": []gin.H{},
	})
}

func createAsset(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "資產已建立",
		"data": gin.H{
			"id": 1,
		},
	})
}

func updateAsset(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "資產已更新",
	})
}

func deleteAsset(c *gin.Context) {
	c.JSON(204, nil)
}

// ============================================
// 任務處理
// ============================================

func getTasks(c *gin.Context) {
	var tasks []models.Task
	if err := db.Preload("User").Find(&tasks).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "獲取任務數據失敗",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "任務列表",
		"data":    tasks,
	})
}

func createTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{
			"message": "請求數據格式錯誤",
			"error":   err.Error(),
		})
		return
	}

	if err := db.Create(&task).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "創建任務失敗",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "任務已創建",
		"data":    task,
	})
}

func updateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := db.First(&task, id).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "任務不存在",
		})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{
			"message": "請求數據格式錯誤",
			"error":   err.Error(),
		})
		return
	}

	if err := db.Save(&task).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "更新任務失敗",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "任務已更新",
		"data":    task,
	})
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	if err := db.First(&task, id).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "任務不存在",
		})
		return
	}

	if err := db.Delete(&task).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "刪除任務失敗",
			"error":   err.Error(),
		})
		return
	}

// ============================================
// 用戶管理
// ============================================

func getUsers(c *gin.Context) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "獲取用戶數據失敗",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "用戶列表",
		"data":    users,
	})
}

func createUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "請求數據格式錯誤",
			"error":   err.Error(),
		})
		return
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "創建用戶失敗",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "用戶已創建",
		"data":    user,
	})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "用戶不存在",
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "請求數據格式錯誤",
			"error":   err.Error(),
		})
		return
	}

	if err := db.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "更新用戶失敗",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "用戶已更新",
		"data":    user,
	})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "用戶不存在",
		})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "刪除用戶失敗",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "用戶已刪除",
	})
}

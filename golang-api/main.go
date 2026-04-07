package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	// 自動遷移模型
	// db.AutoMigrate(&models.User{}, &models.EnvironmentalData{}, &models.Asset{}, &models.Task{})

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
	{
		// 儀表板
		api.GET("/dashboard", getDashboard)

		// 環境數據
		api.GET("/environmental-data", getEnvironmentalData)
		api.POST("/environmental-data", createEnvironmentalData)

		// 資產/庫存
		api.GET("/assets", getAssets)
		api.POST("/assets", createAsset)
		api.PUT("/assets/:id", updateAsset)
		api.DELETE("/assets/:id", deleteAsset)

		// 任務
		api.GET("/tasks", getTasks)
		api.POST("/tasks", createTask)
		api.PUT("/tasks/:id", updateTask)
		api.DELETE("/tasks/:id", deleteTask)
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
	c.JSON(200, gin.H{
		"message": "環境數據列表",
		"data": []gin.H{},
	})
}

func createEnvironmentalData(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "環境數據已記錄",
		"data": gin.H{
			"id": 1,
		},
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
	c.JSON(200, gin.H{
		"message": "任務列表",
		"data": []gin.H{},
	})
}

func createTask(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "任務已建立",
		"data": gin.H{
			"id": 1,
		},
	})
}

func updateTask(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "任務已更新",
	})
}

func deleteTask(c *gin.Context) {
	c.JSON(204, nil)
}

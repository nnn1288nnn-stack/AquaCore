package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// LogEntry - 日志条目结构
type LogEntry struct {
	Type      string `json:"type" binding:"required"`      // 日志类型: login, nginx, mariadb, system
	Message   string `json:"message" binding:"required"`   // 日志消息
	Level     string `json:"level,omitempty"`              // 日志级别: info, warn, error
	User      string `json:"user,omitempty"`               // 用户名 (登录日志)
	IP        string `json:"ip,omitempty"`                 // IP地址
	Timestamp string `json:"timestamp,omitempty"`          // 时间戳
	Service   string `json:"service,omitempty"`            // 服务名
}

// LogHandler - 日志处理器
type LogHandler struct {
	logFile *os.File
}

// NewLogHandler - 创建新的日志处理器
func NewLogHandler() *LogHandler {
	// 确保 logs 目录存在
	os.MkdirAll("logs", 0755)

	// 打开日志文件
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("❌ 无法创建日志文件: %v\n", err)
		return &LogHandler{logFile: nil}
	}

	return &LogHandler{
		logFile: logFile,
	}
}

// LogRequest - 记录日志请求
func (h *LogHandler) LogRequest(c *gin.Context) {
	var entry LogEntry

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// 设置默认值
	if entry.Timestamp == "" {
		entry.Timestamp = time.Now().Format(time.RFC3339)
	}
	if entry.Level == "" {
		entry.Level = "info"
	}

	// 写入文件
	if h.logFile != nil {
		logLine := fmt.Sprintf("[%s] %s %s: %s", entry.Timestamp, entry.Level, entry.Type, entry.Message)
		if entry.User != "" {
			logLine += fmt.Sprintf(" (用户: %s)", entry.User)
		}
		if entry.IP != "" {
			logLine += fmt.Sprintf(" (IP: %s)", entry.IP)
		}
		if entry.Service != "" {
			logLine += fmt.Sprintf(" (服务: %s)", entry.Service)
		}
		logLine += "\n"

		if _, err := h.logFile.WriteString(logLine); err != nil {
			fmt.Printf("❌ 写入日志失败: %v\n", err)
		}
	}

	// 同时输出到控制台
	fmt.Printf("📝 日志记录: %s - %s\n", entry.Type, entry.Message)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "日志已记录",
	})
}

// GetLogs - 获取日志内容
func (h *LogHandler) GetLogs(c *gin.Context) {
	logType := c.Query("type")
	limit := c.DefaultQuery("limit", "100")

	// 这里可以实现日志查询逻辑
	// 暂时返回简单的响应

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"logs": []string{
			fmt.Sprintf("日志类型: %s, 限制: %s", logType, limit),
		},
	})
}

// HealthCheck - 健康检查
func (h *LogHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "log-service",
		"time":    time.Now().Format(time.RFC3339),
	})
}
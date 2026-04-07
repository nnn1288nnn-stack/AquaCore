package models

import (
	"time"

	"gorm.io/gorm"
)

// User - 用戶模型
type User struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name" gorm:"size:100;not null"`
	Phone          string         `json:"phone" gorm:"size:20"`
	Email          string         `json:"email" gorm:"size:100"`
	PreferredLang  string         `json:"preferred_language" gorm:"size:10;default:'zh-TW'"`
	Role           string         `json:"role" gorm:"size:20;default:'operator'"`
	Status         string         `json:"status" gorm:"size:20;default:'active'"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

// EnvironmentalData - 環境數據模型
type EnvironmentalData struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	WaterTemp      float64   `json:"water_temperature" gorm:"column:water_temperature"`
	Salinity       float64   `json:"salinity"`
	DissolvedOxygen float64  `json:"dissolved_oxygen" gorm:"column:dissolved_oxygen"`
	PHLevel        float64   `json:"ph_level" gorm:"column:ph_level"`
	Ammonia        float64   `json:"ammonia"`
	RecordedAt     time.Time `json:"recorded_at"`
	Notes          string    `json:"notes" gorm:"type:text"`
}

// Asset - 資產/庫存模型
type Asset struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"size:100;not null"`
	Category     string    `json:"category" gorm:"size:50"`
	Quantity     int       `json:"quantity" gorm:"default:0"`
	Unit         string    `json:"unit" gorm:"size:20"`
	ReorderLevel int       `json:"reorder_level" gorm:"default:10"`
	UnitCost     float64   `json:"unit_cost" gorm:"type:decimal(10,2)"`
	Supplier     string    `json:"supplier" gorm:"size:100"`
	LastUpdatedBy *uint    `json:"last_updated_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
	User         User      `json:"user,omitempty" gorm:"foreignKey:LastUpdatedBy"`
}

// Task - 任務模型
type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"size:200;not null"`
	Description string    `json:"description" gorm:"type:text"`
	AssignedTo  *uint     `json:"assigned_to"`
	Status      string    `json:"status" gorm:"size:20;default:'pending'"`
	Priority    string    `json:"priority" gorm:"size:20;default:'medium'"`
	DueDate     *time.Time `json:"due_date"`
	CompletedAt *time.Time `json:"completed_at"`
	CreatedBy   *uint     `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Assignee    User      `json:"assignee,omitempty" gorm:"foreignKey:AssignedTo"`
	Creator     User      `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
}

// OperationLog - 操作日誌模型
type OperationLog struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	UserID       *uint     `json:"user_id"`
	OperationType string   `json:"operation_type" gorm:"size:50"`
	ResourceType string    `json:"resource_type" gorm:"size:50"`
	ResourceID   uint      `json:"resource_id"`
	Details      string    `json:"details" gorm:"type:json"`
	IPAddress    string    `json:"ip_address" gorm:"size:45"`
	CreatedAt    time.Time `json:"created_at"`
	User         User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
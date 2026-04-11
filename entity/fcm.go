package entity

import "time"

type FCMToken struct {
	ID         string    `gorm:"primaryKey"`           // ID unik untuk setiap device
	UserID     string    `gorm:"index;not null"`       // UserID dengan index
	FCMToken   string    `gorm:"size:255;unique;not null"` // Token unique
	DeviceType string    `gorm:"size:50"`              // android, ios, web
	DeviceName string    `gorm:"size:100"`             // "Samsung S21", "iPhone 14"
	IsActive   bool      `gorm:"default:true"`         // Token masih aktif?
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
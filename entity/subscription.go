package entity

import "time"

type Subscription struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    string    `gorm:"type:varchar(50);not null;index"`
	IsActive  bool      `gorm:"default:true"`
	ExpiredAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

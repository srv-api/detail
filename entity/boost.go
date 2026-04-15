package entity

import "time"

type Boost struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    string    `gorm:"type:varchar(50);not null;index"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

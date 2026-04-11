package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserDetail struct {
	ID           string         `gorm:"primary_key" json:"id"`
	UserID       string         `gorm:"type:varchar(36);index" json:"user_id"`
	Longitude    float64        `gorm:"type:decimal(10,8);index" json:"longitude"`
	Latitude     float64        `gorm:"type:decimal(10,8);index" json:"latitude"`
	Radius       int            `json:"radius"`
	MaxAge       int            `gorm:"max_age" json:"max_age"`
	MinAge       int            `gorm:"min_age" json:"min_age"`
	GenderTarget string         `gorm:"type:varchar(5)" json:"gender_target"`
	UpdatedBy    string         `gorm:"column:updated_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

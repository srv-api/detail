package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserDetail struct {
	ID           string         `gorm:"primary_key" json:"id"`
	UserID       string         `gorm:"type:varchar(36);index" json:"user_id"`
	IsPremium    bool           `gorm:"is_premium" json:"is_premium"`
	IsBoosted    bool           `gorm:"is_boosted" json:"is_boosted"`
	IsStarLike   bool           `gorm:"is_star_like" json:"is_star_like"`
	IsSee        bool           `gorm:"is_see" json:"is_see"`
	Bio          string         `gorm:"type:varchar(255);index" json:"bio"`
	Longitude    float64        `gorm:"type:decimal(11,8);index" json:"longitude"`
	Latitude     float64        `gorm:"type:decimal(11,8);index" json:"latitude"`
	Radius       int            `json:"radius"`
	MaxAge       int            `gorm:"max_age" json:"max_age"`
	MinAge       int            `gorm:"min_age" json:"min_age"`
	GenderTarget string         `gorm:"type:varchar(5)" json:"gender_target"`
	UpdatedBy    string         `gorm:"column:updated_by"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

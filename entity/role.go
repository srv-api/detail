package entity

import "time"

type Role struct {
	ID         string    `gorm:"type:varchar(36);primaryKey" json:"id"`
	Role       string    `gorm:"type:varchar(100)" json:"role"`
	UserID     string    `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID string    `gorm:"type:varchar(36);index" json:"merchant_id"`
	CreatedAt  time.Time `json:"created_at"`
}

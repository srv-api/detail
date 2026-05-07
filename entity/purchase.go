package entity

import "time"

type PurchasePremium struct {
	ID        string    `gorm:"type:varchar(36);primaryKey"`
	UserID    string    `gorm:"type:varchar(50);index;not null"`
	DetailID  string    `gorm:"type:varchar(36);index" json:"detail_id"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

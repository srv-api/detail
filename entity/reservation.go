package entity

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	ID          uint           `gorm:"primary_key,omitempty" json:"id"`
	UserID      string         `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID  string         `gorm:"type:varchar(36);index" json:"merchant_id"`
	Table       []byte         `gorm:"type:json" json:"table"`
	Name        string         `gorm:"name,omitempty" json:"name"`
	Whatsapp    string         `gorm:"whatsapp,omitempty" json:"whatsapp"`
	Date        string         `gorm:"date,omitempty" json:"date"`
	Time        string         `gorm:"time,omitempty" json:"time"`
	Status      string         `gorm:"status" json:"status"`
	Description string         `gorm:"description" json:"description"`
	CreatedBy   string         `gorm:"created_by" json:"created_by"`
	DeletedBy   string         `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

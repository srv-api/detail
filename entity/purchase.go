package entity

import "time"

type PurchaseHistory struct {
	ID        string    `gorm:"type:varchar(36);primaryKey"`
	UserID    string    `gorm:"type:varchar(50);index;not null"`
	ProductID string    `gorm:"type:varchar(50);not null"`
	Token     string    `gorm:"type:text;uniqueIndex"`
	Amount    float64   `gorm:"type:decimal(10,2)"`
	Status    string    `gorm:"type:varchar(20);default:'pending'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

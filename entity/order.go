package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        string         `gorm:"type:varchar(36);index" json:"user_id"`
	Title         string         `gorm:"title" json:"title"`
	Body          string         `gorm:"body" json:"body"`
	MerchantID    string         `gorm:"type:varchar(36);index;not null" json:"merchant_id"`
	Product       string         `gorm:"product,omitempty" json:"product"`
	ProductParsed []ProductItem  `gorm:"-" json:"product_parsed"` // Tambahan untuk response
	OrderName     string         `gorm:"order_name,omitempty" json:"order_name"`
	Status        int            `gorm:"status" json:"status"`
	CreatedBy     string         `gorm:"created_by" json:"created_by"`
	UpdatedBy     string         `gorm:"updated_by" json:"updated_by"`
	DeletedBy     string         `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type ProductItem struct {
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

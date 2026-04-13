package entity

import "time"

type RoleUser struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	MerchantID   string    `gorm:"type:varchar(36);index" json:"merchant_id"`
	UserID       string    `gorm:"type:varchar(36);index" json:"user_id"`
	RoleID       string    `gorm:"type:varchar(36);index" json:"role_id"`
	PermissionID []byte    `gorm:"type:jsonb" json:"permission_id"`
	CreatedAt    time.Time `json:"created_at"`
}

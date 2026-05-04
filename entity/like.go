package entity

import "time"

type Like struct {
	ID           uint      `gorm:"primaryKey"`
	UserID       string    `gorm:"type:varchar(50);not null;index:idx_user_target,unique"`
	TargetUserID string    `gorm:"type:varchar(50);not null;index:idx_user_target,unique"`
	IsSuperLike  bool      `gorm:"default:false"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

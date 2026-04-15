package entity

import "time"

type Match struct {
	ID        uint      `gorm:"primaryKey"`
	User1ID   string    `gorm:"type:varchar(50);index"`
	User2ID   string    `gorm:"type:varchar(50);index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

package entity

type Currency struct {
	ID       string `gorm:"primary_key,omitempty" json:"id"`
	UserID   string `gorm:"type:varchar(36);index" json:"user_id"`
	DetailID string `gorm:"type:varchar(36);index" json:"detail_id"`
	Currency string `gorm:"currency,omitempty" json:"currency"`
}

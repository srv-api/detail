package entity

type Currency struct {
	ID         string `gorm:"primary_key,omitempty" json:"id"`
	UserID     string `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID string `gorm:"type:varchar(36);index" json:"merchant_id"`
	Currency   string `gorm:"currency,omitempty" json:"currency"`
}

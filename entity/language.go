package entity

type Language struct {
	ID         string `gorm:"primary_key,omitempty" json:"id"`
	UserID     string `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID string `gorm:"type:varchar(36);index" json:"merchant_id"`
	Language   string `gorm:"language,omitempty" json:"language"`
}

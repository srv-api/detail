package dto

import "time"

type GetByContentSettingIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type ContentSettingRequest struct {
	ID          string `gorm:"primary_key,omitempty" json:"id"`
	UserID      string `gorm:"type:varchar(36);index" json:"user_id"`
	MerchantID  string `gorm:"type:varchar(36);index" json:"merchant_id"`
	Logo        string `gorm:"logo,omitempty" json:"logo"`
	Title       string `gorm:"title,omitempty" json:"title"`
	Description string `gorm:"description" json:"description"`
	LinkAndroid string `gorm:"link_android" json:"link_android"`
	LinkApple   string `gorm:"link_apple" json:"link_apple"`
}

type ContentSettingResponse struct {
	ButtonHeader []ButtonHeader `json:"button_header"`
	TopHeader    []TopHeader    `json:"top_header"`
	Feature      []Feature      `json:"feature"`
	Footer       []Footer       `json:"footer"`
	UpdatedBy    string         `json:"updated_by"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type UpdateContentSettingRequest struct {
	ID           string         `json:"id"`
	UserID       string         `json:"user_id"`
	MerchantID   string         `json:"merchant_id"`
	TopHeader    []TopHeader    `json:"top_header"`
	ButtonHeader []ButtonHeader `json:"button_header"`
	Feature      []Feature      `json:"feature"`
	Footer       []Footer       `json:"footer"`
	UpdatedBy    string         `json:"updated_by"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type UpdateContentSettingResponse struct {
	ID           string         `json:"id"`
	UserID       string         `json:"user_id"`
	MerchantID   string         `json:"merchant_id"`
	TopHeader    []TopHeader    `json:"top_header"`
	ButtonHeader []ButtonHeader `json:"button_header"`
	Feature      []Feature      `json:"feature"`
	Footer       []Footer       `json:"footer"`
	UpdatedBy    string         `json:"updated_by"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

// Sub-sections
type TopHeader struct {
	Logo        string `json:"logo"`
	Title       string `json:"title"`
	Description string `json:"description"`
	LinkAndroid string `json:"link_android"`
	LinkApple   string `json:"link_apple"`
}

type ButtonHeader struct {
	ButtonName string `json:"button_name"`
	Link       string `json:"link"`
}

type Feature struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Footer struct {
	Company       string `json:"company"`
	InstagramLink string `json:"instagram_link"`
	FacebookLink  string `json:"facebook_link"`
}

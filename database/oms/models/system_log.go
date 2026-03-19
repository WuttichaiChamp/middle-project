package models

import "time"

type PlatformLog struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ShopCode     string    `gorm:"size:100" json:"shop_code"`
	Platform     string    `gorm:"size:100" json:"platform"`
	Action       string    `gorm:"size:100" json:"action"`
	Request      string    `gorm:"type:jsonb" json:"request"`
	Response     string    `gorm:"type:jsonb" json:"response"`
	Status       string    `gorm:"size:50" json:"status"`
	ErrorMessage string    `gorm:"type:text" json:"error_message"`
	CreatedAt    time.Time `json:"created_at"`
}

func (PlatformLog) TableName() string {
	return "platform_log"
}

package models

import "github.com/google/uuid"

type Platform struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PlatformCode string    `json:"platform_code"`
	PlatformName string    `json:"platform_name"`
	Active       bool      `json:"active"`
}

func (Platform) TableName() string { return "platform" }

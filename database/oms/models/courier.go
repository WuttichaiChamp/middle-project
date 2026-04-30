package models

import (
	"time"

	"github.com/google/uuid"
)

type Courier struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"type:text" json:"name"`
	Code        string    `gorm:"type:text" json:"code"`
	AuthType    string    `gorm:"type:text" json:"auth_type"`
	Credentials string    `gorm:"type:text" json:"credentials"`
	IsActive    bool      `json:"is_active"`
	CreateBy    string    `gorm:"type:text" json:"create_by"`
	CreateDtm   time.Time `json:"create_dtm"`
	UpdateBy    string    `gorm:"type:text" json:"update_by"`
	UpdateDtm   time.Time `json:"update_dtm"`
}

func (Courier) TableName() string { return "courier" }

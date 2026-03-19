package models

import (
	"time"

	"github.com/google/uuid"
)

type Courier struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CourierType string    `json:"courier_type"`
	CourierCode string    `json:"courier_code"`
	CourierName string    `json:"courier_name"`
	CreateBy    string    `json:"create_by"`
	CreateDtm   time.Time `json:"create_dtm"`
	UpdateBy    string    `json:"update_by"`
	UpdateDtm   time.Time `json:"update_dtm"`
}

func (Courier) TableName() string { return "courier" }

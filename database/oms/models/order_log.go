package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderLog struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null;index" json:"order_id"`
	Action    string    `gorm:"type:varchar(100);not null" json:"action"`
	OldStatus string    `gorm:"type:varchar(50)" json:"old_status"`
	NewStatus string    `gorm:"type:varchar(50)" json:"new_status"`
	Remark    string    `gorm:"type:varchar(255)" json:"remark"`
	CreateBy  string    `gorm:"type:varchar(255)" json:"create_by"`
	CreateDtm time.Time `json:"create_dtm"`
}

func (OrderLog) TableName() string { return "order_log" }

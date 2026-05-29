package models

import (
	"github.com/google/uuid"
)

type PlatformStatement struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ShopCode string    `gorm:"type:text" json:"shop_code"`

	// 	shop_code
	// statement_code
	// platform_code
	// type
	// doc_ref
	// period_form_dtm
	// period_to_dtm
	// total_order_amount
	// total_fee_amount
	// total_settlement_amount
	// currency
	// status
	// raw_payload
	// create_dtm
	// update_dtm
}

func (PlatformStatement) TableName() string { return "platform_statement" }

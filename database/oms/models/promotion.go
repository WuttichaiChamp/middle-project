package models

import (
	"time"

	"github.com/google/uuid"
)

type PromotionType string
type PromotionStatus string

const (
	PromoTypeTradeDiscount PromotionType = "TRADE_DISCOUNT"
	PromoTypeAGetB         PromotionType = "A_GET_B"
	PromoTypeFirstOrder    PromotionType = "FIRST_ORDER"

	PromoStatusPending   PromotionStatus = "PENDING"
	PromoStatusCompleted PromotionStatus = "COMPLETED"
	PromoStatusCancelled PromotionStatus = "CANCELLED"
)

// Promotion master table - IMMUTABLE after usage
type Promotion struct {
	ID           uuid.UUID     `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	PromoCode    string        `gorm:"type:varchar(50);unique;not null" json:"promo_code"`
	PromoType    PromotionType `gorm:"type:varchar(50);not null" json:"promo_type"`
	EffectiveDtm time.Time     `gorm:"type:timestamp;not null" json:"effective_dtm"`
	ExpireDtm    time.Time     `gorm:"type:timestamp;not null" json:"expire_dtm"`
	Status       string        `gorm:"type:varchar(50);not null" json:"status"`
	CreateBy     string        `gorm:"type:varchar(255);not null" json:"create_by"`
	CreateDtm    time.Time     `gorm:"autoCreateTime" json:"create_dtm"`
	UpdateBy     *string       `gorm:"type:varchar(255)" json:"update_by,omitempty"`
	UpdateDtm    *time.Time    `gorm:"autoUpdateTime" json:"update_dtm,omitempty"`
	IsNotAuto    *bool         `gorm:"type:boolean;default:false" json:"is_not_auto,omitempty"`
}

// Promotion items (source items that trigger the promotion)
type PromotionItem struct {
	ID         uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	PromoID    uuid.UUID  `gorm:"type:uuid;not null" json:"promo_id"`
	PromoItem  string     `gorm:"type:varchar(255);not null" json:"promo_item"`
	Qty        float64    `gorm:"type:decimal(10,4);not null" json:"qty"`
	CreateDtm  time.Time  `gorm:"autoCreateTime" json:"create_dtm"`
	UpdateDtm  *time.Time `gorm:"autoUpdateTime" json:"update_dtm,omitempty"`
	CreateBy   string     `gorm:"type:varchar(255)" json:"create_by"`
	UpdateBy   *string    `gorm:"type:varchar(255)" json:"update_by,omitempty"`
	QtyBenefit *float64   `gorm:"type:decimal(10,4)" json:"qty_benefit,omitempty"`
	Seq        *int       `gorm:"type:int" json:"seq,omitempty"`
	QtyLimit   *int       `gorm:"type:int" json:"qty_limit,omitempty"`
}

// Promotion item source (hierarchy mapping for source items)
type PromotionItemSource struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	PromoItemID   uuid.UUID `gorm:"type:uuid;not null" json:"promo_item_id"`
	HierarchyType string    `gorm:"type:varchar(50);not null" json:"hierarchy_type"` // PRODUCT, PRODUCT_GROUP, etc.
	HierarchyCode string    `gorm:"type:varchar(255);not null" json:"hierarchy_code"`
	Qty           float64   `gorm:"type:decimal(10,4);not null" json:"qty"`
}

// Promotion item destination (benefit items)
type PromotionItemDestination struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	PromoItemID   uuid.UUID `gorm:"type:uuid;not null" json:"promo_item_id"`
	HierarchyType string    `gorm:"type:varchar(50);not null" json:"hierarchy_type"`
	HierarchyCode string    `gorm:"type:varchar(255);not null" json:"hierarchy_code"`
	Qty           float64   `gorm:"type:decimal(10,4);not null" json:"qty"`
}

// Promotion usage - tracks when promotion is applied to sales order
// Creates IMMUTABLE snapshot of promotion rules
type PromotionItemUsage struct {
	ID                  uuid.UUID       `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	PromotionCode       string          `gorm:"type:varchar(50);not null" json:"promotion_code"`
	PromotionItem       string          `gorm:"type:varchar(255);not null" json:"promotion_item"`
	Status              PromotionStatus `gorm:"type:varchar(50);not null" json:"status"`
	CreateBy            string          `gorm:"type:varchar(255);not null" json:"create_by"`
	CreateDtm           time.Time       `gorm:"autoCreateTime" json:"create_dtm"`
	UpdateDtm           *time.Time      `gorm:"autoUpdateTime" json:"update_dtm,omitempty"`
	QtyPromo            *float64        `gorm:"type:decimal(10,4)" json:"qty_promo,omitempty"`
	SaleCode            *string         `gorm:"type:varchar(50)" json:"sale_code,omitempty"`
	QtyBenefit          *float64        `gorm:"type:decimal(10,4)" json:"qty_benefit,omitempty"`
	PromoType           PromotionType   `gorm:"type:varchar(50);not null" json:"promo_type"`
	SnapshotSource      string          `gorm:"type:jsonb;not null" json:"snapshot_source"`      // IMMUTABLE snapshot
	SnapshotDestination string          `gorm:"type:jsonb;not null" json:"snapshot_destination"` // IMMUTABLE snapshot
}

// Promotion usage destination - tracks benefit items given
type PromotionItemUsageDestination struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	PromoItemUsageID  uuid.UUID `gorm:"type:uuid;not null" json:"promo_item_usage_id"`
	HierarchyType     string    `gorm:"type:varchar(50);not null" json:"hierarchy_type"`
	HierarchyCode     string    `gorm:"type:varchar(255);not null" json:"hierarchy_code"`
	Qty               float64   `gorm:"type:decimal(10,4);not null" json:"qty"`
}

// Approval table for promotions requiring approval
type Approval struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	ApproveCode  string    `gorm:"type:varchar(50);unique;not null" json:"approve_code"`
}

func (Promotion) TableName() string {
	return "promotion"
}

func (PromotionItem) TableName() string {
	return "promotion_item"
}

func (PromotionItemSource) TableName() string {
	return "promotion_item_source"
}

func (PromotionItemDestination) TableName() string {
	return "promotion_item_destination"
}

func (PromotionItemUsage) TableName() string {
	return "promotion_item_usage"
}

func (PromotionItemUsageDestination) TableName() string {
	return "promotion_item_usage_destination"
}

func (Approval) TableName() string {
	return "approval"
}

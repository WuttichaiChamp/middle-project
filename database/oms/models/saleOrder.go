package models

import (
	"time"

	"github.com/google/uuid"
)

// Sales Order Status Constants
const (
	SOStatusTemp     = "TEMP"
	SOStatusNew      = "NEW"
	SOStatusPending  = "PENDING"
	SOStatusApproved = "APPROVED"
)

type SalesOrder struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	SaleCode          string     `gorm:"type:varchar(50);uniqueIndex;not null" json:"sale_code"`
	QuotationID       *uuid.UUID `gorm:"type:uuid" json:"quotation_id"`
	Quotation         *Quotation `gorm:"foreignKey:QuotationID" json:"quotation,omitempty"`
	Source            *string    `gorm:"type:varchar(50)" json:"source"`             // MANUAL, CONVERT_FROM_QUOTATION, INTERFACE
	InterfacePlatform *string    `gorm:"type:varchar(50)" json:"interface_platform"` // SHOPEE, LAZADA, TIKTOK, LINE
	CustomerOrderRef  *string    `gorm:"type:varchar(100)" json:"customer_order_ref"`

	CompanyCode  string `gorm:"type:varchar(50)" json:"company_code"`
	SiteCode     string `gorm:"type:varchar(50)" json:"site_code"`
	CustomerCode string `gorm:"type:varchar(50)" json:"customer_code"`
	CustomerName string `gorm:"type:varchar(255)" json:"customer_name"`

	DeliveryDate *time.Time `json:"delivery_date"`

	SoldToCode string `gorm:"type:varchar(50)" json:"sold_to_code"`
	BillToCode string `gorm:"type:varchar(50);column:bill_to_code" json:"bill_to_code"`
	BillToAddr string `gorm:"type:varchar(255);column:bill_to_address" json:"bill_to_address"`
	ShipToCode string `gorm:"type:varchar(50);column:ship_to_code" json:"ship_to_code"`
	ShipToType string `gorm:"type:varchar(50);column:ship_to_type" json:"ship_to_type"`
	ShipToAddr string `gorm:"type:varchar(255);column:ship_to_address" json:"ship_to_address"`

	PickingType        string     `gorm:"type:varchar(50)" json:"picking_type"`
	ShipDate           *time.Time `json:"ship_date"`
	TransportZone      string     `gorm:"type:varchar(50)" json:"transport_zone"`
	TransportVendor    string     `gorm:"type:varchar(100)" json:"transport_vendor"`
	ShippingLeadTime   int        `gorm:"default:0" json:"shipping_lead_time"`
	MinimumShelfLife   int        `gorm:"default:0" json:"minimum_shelf_life"`
	InvoiceUom         string     `gorm:"type:varchar(50)" json:"invoice_uom"`
	SalesChannel       string     `gorm:"type:varchar(50)" json:"sales_channel"`
	BatchCannotReverse bool       `gorm:"default:true" json:"batch_cannot_reverse"`

	DeliveryMethod     string  `gorm:"type:varchar(50)" json:"delivery_method"`
	TransportCostType  string  `gorm:"type:varchar(50)" json:"transport_cost_type"`
	TotalTransportCost float64 `gorm:"type:decimal(10,4);default:0" json:"total_transport_cost"`

	TotalAmount   float64 `gorm:"type:decimal(10,4);default:0" json:"total_amount"`
	TotalWeight   float64 `gorm:"type:decimal(10,4);default:0" json:"total_weight"`
	TotalVat      float64 `gorm:"type:decimal(10,4);default:0" json:"total_vat"`
	TotalDiscount float64 `gorm:"type:decimal(10,4);default:0" json:"total_discount"`

	SubtotalExclVat             float64 `gorm:"type:decimal(10,4);default:0" json:"subtotal_excl_vat"`
	SubtotalExclTransport       float64 `gorm:"type:decimal(10,4);default:0" json:"subtotal_excl_transport"`
	SubtotalWeightExclTransport float64 `gorm:"type:decimal(10,4);default:0" json:"subtotal_weight_excl_transport"`

	PaymentMethod   string `gorm:"type:varchar(50)" json:"payment_method"`
	PaymentTermCode string `gorm:"type:varchar(50)" json:"payment_term_code"`

	SalePersonCode string `gorm:"type:varchar(50)" json:"sale_person_code"`

	EffectiveDatePrice *time.Time `json:"effective_date_price"`
	ExpirePriceDay     int        `json:"expire_price_day"`
	ExpirePriceDate    *time.Time `json:"expire_price_date"`

	PassPriceList   string `gorm:"type:varchar(50)" json:"pass_price_list"`
	PassAtpCheck    string `gorm:"type:varchar(50)" json:"pass_atp_check"`
	PassCreditLimit string `gorm:"type:varchar(50)" json:"pass_credit_limit"`
	PassPriceExpire string `gorm:"type:varchar(50)" json:"pass_price_expire"`

	Status              string `gorm:"type:varchar(20);not null" json:"status"`
	StatusPaymentAction string `gorm:"type:varchar(50)" json:"status_payment_action"`
	StatusPayment       string `gorm:"type:varchar(50)" json:"status_payment"`

	Remark         string `gorm:"type:varchar(255)" json:"remark"`
	RemarkApproval string `gorm:"type:varchar(255)" json:"remark_approval"`

	IsApproved    bool   `gorm:"default:false" json:"is_approved"`
	StatusApprove string `gorm:"type:varchar(50)" json:"status_approve"`

	CreatedAt time.Time `gorm:"column:create_date" json:"create_date"`
	CreatedBy string    `gorm:"column:create_by" json:"create_by"`
	UpdatedAt time.Time `gorm:"column:update_date" json:"update_date"`
	UpdatedBy string    `gorm:"column:update_by" json:"update_by"`

	Items []SalesOrderItem `gorm:"foreignKey:SaleID;constraint:OnDelete:CASCADE" json:"items"`
}

type SalesOrderItem struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	SaleID uuid.UUID `gorm:"type:uuid;not null;index" json:"sale_id"`

	ProductCode string `gorm:"type:varchar(50)" json:"product_code"`
	ProductDesc string `gorm:"type:varchar(255)" json:"product_desc"`

	Qty       float64 `gorm:"type:decimal(10,4)" json:"qty"`
	OriginQty float64 `gorm:"type:decimal(10,4)" json:"origin_qty"`
	Unit      string  `gorm:"type:varchar(20)" json:"unit"`

	PriceListUnit float64 `gorm:"type:decimal(10,4)" json:"price_list_unit"`
	SaleQty       float64 `gorm:"type:decimal(10,4)" json:"sale_qty"`
	SaleUnit      string  `gorm:"type:varchar(20)" json:"sale_unit"`
	SaleUnitType  string  `gorm:"type:varchar(50)" json:"sale_unit_type"`

	PriceUnit       float64 `gorm:"type:decimal(10,4)" json:"price_unit"`
	PassPriceWeight string  `gorm:"-" json:"pass_price_weight"`
	PriceLimit      string  `gorm:"-" json:"price_limit"`

	TotalAmount               float64 `gorm:"type:decimal(10,4);default:0" json:"total_amount"`
	SubtotalExclVat           float64 `gorm:"type:decimal(10,4);default:0" json:"subtotal_excl_vat"`
	TotalVat                  float64 `gorm:"type:decimal(10,4);default:0" json:"total_vat"`
	TransportCostUnit         float64 `gorm:"type:decimal(10,4);default:0" json:"transport_cost_unit"`
	SubtotalExclTransport     float64 `gorm:"type:decimal(10,4);default:0" json:"subtotal_excl_transport"`
	NetPriceUnitExclTransport float64 `gorm:"type:decimal(10,4);default:0" json:"net_price_unit_excl_transport"`

	WeightUnit    float64 `gorm:"type:decimal(10,4);default:0" json:"weight_unit"`
	AvgWeightUnit float64 `gorm:"type:decimal(10,4);default:0" json:"avg_weight_unit"`
	TotalWeight   float64 `gorm:"type:decimal(10,4);default:0" json:"total_weight"`

	TransportCostWeightUnit        float64 `gorm:"type:decimal(10,4);default:0" json:"transport_cost_weight_unit"`
	SubtotalWeightExclTransport    float64 `gorm:"type:decimal(10,4);default:0" json:"subtotal_weight_excl_transport"`
	NetPricePerWeightExclTransport float64 `gorm:"type:decimal(10,4);default:0" json:"net_price_per_weight_excl_transport"`

	Status  string `gorm:"type:varchar(20)" json:"status"`
	Remark  string `gorm:"type:varchar(255)" json:"remark"`
	UnitUom string `gorm:"type:varchar(50)" json:"unit_uom"`

	TotalDiscount        float64 `gorm:"type:decimal(10,4);default:0" json:"total_discount"`
	TotalDiscountPercent float64 `gorm:"type:decimal(10,4);default:0" json:"total_discount_percent"`

	CreatedAt time.Time `gorm:"column:create_date" json:"create_date"`
	CreatedBy string    `gorm:"column:create_by" json:"create_by"`
	UpdatedAt time.Time `gorm:"column:update_date" json:"update_date"`
	UpdatedBy string    `gorm:"column:update_by" json:"update_by"`

	PromotionID  *uuid.UUID `gorm:"type:uuid" json:"promotion_id"`
	ItemType     string     `gorm:"type:varchar(20)" json:"item_type"`
	ParentItemID *uuid.UUID `gorm:"type:uuid" json:"parent_item_id"`
}

// TableName overrides the table name used by SalesOrder to `sale`
func (SalesOrder) TableName() string {
	return "sale"
}

// TableName overrides the table name used by SalesOrderItem to `sale_item`
func (SalesOrderItem) TableName() string {
	return "sale_item"
}

// Request/Response Types for Sales Order API

type CreateSalesOrderRequest struct {
	Source             string                        `json:"source"`
	InterfacePlatform  *string                       `json:"interface_platform,omitempty"`
	CustomerOrderRef   *string                       `json:"customer_order_ref,omitempty"`
	CompanyCode        string                        `json:"company_code"`
	SiteCode           string                        `json:"site_code"`
	CustomerCode       string                        `json:"customer_code"`
	CustomerName       string                        `json:"customer_name"`
	DeliveryDate       *string                       `json:"delivery_date,omitempty"`
	SoldToCode         string                        `json:"sold_to_code"`
	BillToCode         *string                       `json:"bill_to_code,omitempty"`
	BillToAddr         *string                       `json:"bill_to_address,omitempty"`
	ShipToCode         *string                       `json:"ship_to_code,omitempty"`
	ShipToType         *string                       `json:"ship_to_type,omitempty"`
	ShipToAddr         *string                       `json:"ship_to_address,omitempty"`
	ShipDate           *string                       `json:"ship_date,omitempty"`
	PickingType        *string                       `json:"picking_type,omitempty"`
	TransportZone      *string                       `json:"transport_zone,omitempty"`
	TransportVendor    *string                       `json:"transport_vendor,omitempty"`
	ShippingLeadTime   *int                          `json:"shipping_lead_time,omitempty"`
	MinimumShelfLife   *int                          `json:"minimum_shelf_life,omitempty"`
	InvoiceUom         *string                       `json:"invoice_uom,omitempty"`
	SalesChannel       *string                       `json:"sales_channel,omitempty"`
	BatchCannotReverse *bool                         `json:"batch_cannot_reverse,omitempty"`
	DeliveryMethod     *string                       `json:"delivery_method,omitempty"`
	TransportCostType  *string                       `json:"transport_cost_type,omitempty"`
	PaymentMethod      *string                       `json:"payment_method,omitempty"`
	PaymentTermCode    *string                       `json:"payment_term_code,omitempty"`
	SalePersonCode     *string                       `json:"sale_person_code,omitempty"`
	EffectiveDatePrice *string                       `json:"effective_date_price,omitempty"`
	ExpirePriceDate    *string                       `json:"expire_price_date,omitempty"`
	ExpirePriceDay     *int                          `json:"expire_price_day,omitempty"`
	Remark             *string                       `json:"remark,omitempty"`
	Status             *string                       `json:"status,omitempty"`
	Items              []CreateSalesOrderItemRequest `json:"items"`
}

type CreateSalesOrderItemRequest struct {
	ProductCode    string   `json:"product_code"`
	ProductDesc    string   `json:"product_desc"`
	Qty            float64  `json:"qty"`
	Unit           string   `json:"unit"`
	ItemType       *string  `json:"item_type,omitempty"`
	PassPrice      bool     `json:"pass_price"`
	PriceListUnit  float64  `json:"price_list_unit"`
	SaleQty        *float64 `json:"sale_qty,omitempty"`
	SaleUnit       *string  `json:"sale_unit,omitempty"`
	PriceUnit      float64  `json:"price_unit"`
	TotalAmount    *float64 `json:"total_amount,omitempty"`
	DiscountByLine *float64 `json:"discount_by_line,omitempty"`
	PromotionID    *string  `json:"promotion_id,omitempty"`
	ParentItemID   *string  `json:"parent_item_id,omitempty"`
}

type SalesOrderPromotionRequest struct {
	PromotionID    string  `json:"promotion_id"`
	PromotionCode  string  `json:"promotion_code"`
	PromotionType  string  `json:"promotion_type"`
	DiscountAmount float64 `json:"discount_amount"`
}

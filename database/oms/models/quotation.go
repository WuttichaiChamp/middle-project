package models

import (
	"time"

	"github.com/google/uuid"
)

type Quotation struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	QuotationCode string `gorm:"column:quotation_code;type:varchar(50);uniqueIndex;not null" json:"quotation_code"`
	CompanyCode   string `gorm:"type:varchar(50);not null" json:"company_code"`
	SiteCode      string `gorm:"type:varchar(50);not null" json:"site_code"`

	CustomerCode string `gorm:"type:varchar(50);index" json:"customer_code"`
	CustomerName string `gorm:"type:varchar(255);not null" json:"customer_name"`

	DeliveryDate *Date `gorm:"type:timestamp;index" json:"delivery_date"`

	Status        string `gorm:"type:varchar(20);not null;default:'TEMP';index" json:"status"`
	IsApproved    bool   `gorm:"not null;default:false" json:"is_approved"`
	StatusApprove string `gorm:"type:varchar(20)" json:"status_approve"`

	TotalAmount   float64 `gorm:"type:decimal(15,4);not null;default:0" json:"total_amount"`
	TotalVat      float64 `gorm:"type:decimal(15,4);not null;default:0" json:"total_vat"`
	TotalDiscount float64 `gorm:"type:decimal(15,4);not null;default:0" json:"total_discount"`
	TotalWeight   float64 `gorm:"type:decimal(15,4);not null;default:0" json:"total_weight"`

	Remark         string `gorm:"type:varchar(255)" json:"remark"`
	RemarkApproval string `gorm:"type:varchar(255)" json:"remark_approval"`

	CreatedAt      time.Time `gorm:"column:create_date;index" json:"created_at"`
	CreatedBy      string    `gorm:"column:create_by" json:"created_by"`
	UpdatedAt      time.Time `gorm:"column:update_date" json:"updated_at"`
	UpdatedBy      string    `gorm:"column:update_by" json:"updated_by"`
	SoldToCode     string    `gorm:"type:varchar(50)" json:"sold_to_code"`
	BillToCode     string    `gorm:"type:varchar(50)" json:"bill_to_code"`
	BillToAddress  string    `gorm:"type:text" json:"bill_to_address"`
	ShipToCode     string    `gorm:"type:varchar(50)" json:"shipping_code"`
	ShipToType     string    `gorm:"type:varchar(50)" json:"ship_to_type"`
	ShipToAddress  string    `gorm:"type:text" json:"shipping_address"`
	ContractPerson string    `gorm:"type:varchar(150)" json:"contract_person"`
	PhoneNumber    string    `gorm:"type:varchar(50)" json:"phone_number"`
	Email          string    `gorm:"type:varchar(150)" json:"email"`

	DeliveryMethod    string `gorm:"type:varchar(50)" json:"delivery_method"`
	TransportCostType string `gorm:"type:varchar(50)" json:"transport_cost_type"`

	TotalTransportCost    float64 `gorm:"type:decimal(15,4);default:0"`
	SubtotalExclVat       float64 `gorm:"type:decimal(15,4);default:0"`
	SubtotalExclTransport float64 `gorm:"type:decimal(15,4);default:0"`

	PaymentTermCode    string `gorm:"type:varchar(50)" json:"payment_term_code"`
	EffectiveDatePrice *Date  `gorm:"column:effective_date_price;type:timestamp" json:"valid_from"`
	ExpirePriceDate    *Date  `gorm:"column:expire_price_date;type:timestamp" json:"valid_to"`
	ExpirePriceDay     *int   `gorm:"column:expire_price_day;type:integer" json:"expire_price_day"`

	PassPrice       bool   `gorm:"not null;default:false"`
	PassPriceList   string `gorm:"type:varchar(50)"`
	PassAtpCheck    string `gorm:"type:varchar(50)"`
	PassCreditLimit string `gorm:"type:varchar(50)"`
	PassPriceExpire string `gorm:"type:varchar(50)"`

	Items []QuotationItem `gorm:"foreignKey:QuotationID;constraint:OnDelete:CASCADE"`
}

type QuotationItem struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	QuotationID uuid.UUID `gorm:"type:uuid;not null;index"`

	ProductCode string  `gorm:"type:varchar(50);not null" json:"product_code"`
	ProductDesc string  `gorm:"type:varchar(255);not null" json:"product_desc"`
	Qty         float64 `gorm:"type:decimal(15,4);not null" json:"qty"`
	Unit        string  `gorm:"type:varchar(20);not null" json:"unit"`

	PriceUnit   float64 `gorm:"type:decimal(15,4);not null" json:"price_unit"`
	TotalAmount float64 `gorm:"type:decimal(15,4);not null;default:0" json:"total_amount"`

	TotalVat      float64 `gorm:"type:decimal(15,4);not null;default:0" json:"total_vat"`
	TotalDiscount float64 `gorm:"type:decimal(15,4);not null;default:0" json:"total_discount"`

	ItemType     string     `gorm:"type:varchar(20);not null;default:'NORMAL'" json:"item_type"`
	PromotionID  *uuid.UUID `gorm:"type:uuid" json:"promotion_id"`
	ParentItemID *uuid.UUID `gorm:"type:uuid" json:"parent_item_id"`

	CreatedAt     time.Time `gorm:"column:create_date" json:"created_at"`
	CreatedBy     string    `gorm:"column:create_by" json:"created_by"`
	UpdatedAt     time.Time `gorm:"column:update_date" json:"updated_at"`
	UpdatedBy     string    `gorm:"column:update_by" json:"updated_by"`
	PriceListUnit float64   `gorm:"type:decimal(15,4);default:0" json:"price_list_unit"`
	SaleQty       float64   `gorm:"type:decimal(15,4);default:0" json:"sale_qty"`
	SaleUnit      string    `gorm:"type:varchar(20)" json:"sale_unit"`

	PassPrice  bool    `gorm:"type:boolean;default:false" json:"pass_price"`
	PassWeight float64 `gorm:"type:decimal(15,4);default:0" json:"pass_weight"`

	SubtotalExclVat           float64 `gorm:"type:decimal(15,4);default:0"`
	SubtotalExclTransport     float64 `gorm:"type:decimal(15,4);default:0"`
	NetPriceUnitExclTransport float64 `gorm:"type:decimal(15,4);default:0"`

	WeightUnit    float64 `gorm:"type:decimal(15,4);default:0"`
	AvgWeightUnit float64 `gorm:"type:decimal(15,4);default:0"`
	TotalWeight   float64 `gorm:"type:decimal(15,4);default:0"`

	TotalDiscountPercent float64 `gorm:"type:decimal(15,4);default:0"`

	Status string `gorm:"type:varchar(20);not null"`
}

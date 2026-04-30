package models

import (
	"time"

	"github.com/google/uuid"
)

type Shipment struct {
	ID                 uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	ShipmentCode       string         `gorm:"type:text" json:"shipment_code"`
	CourierCode        string         `gorm:"type:text" json:"courier_code"`
	ShipmentType       string         `gorm:"type:text" json:"shipment_type"`
	TrackingNo         string         `gorm:"type:text" json:"tracking_no"`
	Status             string         `gorm:"type:text" json:"status"`
	DeliveryOptionID   string         `gorm:"type:text" json:"delivery_option_id"`
	ShippingProviderID string         `gorm:"type:text" json:"shipping_provider_id"`
	AssignedBy         string         `gorm:"type:text" json:"assigned_by"`
	LabelURL           string         `gorm:"type:text" json:"label_url"`
	CreateBy           string         `gorm:"type:text" json:"create_by"`
	CreateDtm          time.Time      `json:"create_dtm"`
	UpdateBy           string         `gorm:"type:text" json:"update_by"`
	UpdateDtm          time.Time      `json:"update_dtm"`
	ShipmentItems      []ShipmentItem `gorm:"foreignKey:ShipmentID" json:"shipment_items"`
}

func (Shipment) TableName() string { return "shipment" }

type ShipmentItem struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ShipmentID      string    `gorm:"type:text" json:"shipment_id"`
	ShipmentItem    string    `gorm:"type:text" json:"shipment_item"`
	DocumentRef     string    `gorm:"type:text" json:"document_ref"`
	DocumentItemRef string    `gorm:"type:text" json:"document_item_ref"`
	SaleCode        string    `gorm:"type:text" json:"sale_code"`
	SaleItem        string    `gorm:"type:text" json:"sale_item"`
	ProductCode     string    `gorm:"type:text" json:"product_code"`
	Qty             string    `gorm:"type:text" json:"qty"`
	Status          string    `gorm:"type:text" json:"status"`
	AssignedBy      string    `gorm:"type:text" json:"assigned_by"`
	CreateBy        string    `gorm:"type:text" json:"create_by"`
	CreateDtm       time.Time `json:"create_dtm"`
	UpdateBy        string    `gorm:"type:text" json:"update_by"`
	UpdateDtm       time.Time `json:"update_dtm"`
}

func (ShipmentItem) TableName() string { return "shipment_item" }

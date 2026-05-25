package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderPackage struct {
	ID                uuid.UUID          `gorm:"type:uuid;primaryKey" json:"id"`
	PackageCode       string             `json:"package_code"`
	ShopCode          string             `json:"shop_code"`
	IsCod             bool               `json:"is_cod"`
	DocumentRef       string             `json:"document_ref"`
	CourierCode       string             `json:"courier_code"`
	TrackingNo        string             `json:"tracking_no"`
	Status            string             `json:"status"`
	StatusShipment    string             `json:"status_shipment"`
	Weight            float64            `json:"weight"`
	Dimension         float64            `json:"dimension"`
	DimensionDesc     string             `json:"dimension_desc"`
	LabelURL          string             `json:"label_url"`
	ShipDate          time.Time          `json:"ship_date"`
	ExternalRef       string             `json:"external_ref"`
	CreateBy          string             `json:"create_by"`
	CreateDtm         time.Time          `json:"create_dtm"`
	UpdateBy          string             `json:"update_by"`
	UpdateDtm         time.Time          `json:"update_dtm"`
	OrderPackageItems []OrderPackageItem `gorm:"foreignKey:PackageID" json:"order_package_items"`
}

func (OrderPackage) TableName() string { return "order_package" }

type OrderPackageItem struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PackageID       uuid.UUID `gorm:"type:uuid" json:"package_id"`
	ProductCode     string    `json:"product_code"`
	Qty             float64   `json:"qty"`
	UnitCode        string    `json:"unit_code"`
	DocumentRefItem string    `json:"document_ref_item"`
	CreateBy        string    `json:"create_by"`
	CreateDtm       time.Time `json:"create_dtm"`
	UpdateBy        string    `json:"update_by"`
	UpdateDtm       time.Time `json:"update_dtm"`
}

func (OrderPackageItem) TableName() string { return "order_package_item" }

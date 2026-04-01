package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID                  uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
	OrderCode           string      `json:"order_code"`
	OrderType           string      `json:"order_type"`
	OrderDate           *time.Time  `json:"order_date"`
	ShopCode            string      `json:"shop_code"`
	TenantID            uuid.UUID   `gorm:"type:uuid" json:"tenant_id"`
	CustomerCode        string      `json:"customer_code"`
	ShipToCode          string      `json:"ship_to_code"`
	BillToCode          string      `json:"bill_to_code"`
	TransportZone       string      `json:"transport_zone"`
	InterfaceQty        float64     `json:"interface_qty"`
	InterfaceUnitCode   string      `json:"interface_unit_code"`
	Qty                 float64     `json:"qty"`
	UnitCode            string      `json:"unit_code"`
	EstimatePickingDate *time.Time  `json:"estimate_picking_date"`
	DeliveryDate        *time.Time  `json:"delivery_date"`
	SubmitDate          *time.Time  `json:"submit_date"`
	Status              string      `json:"status"`
	StatusShipment      string      `json:"status_shipment"`
	DocumentRefType     string      `json:"document_ref_type"`
	DocumentRef         string      `json:"document_ref"`
	Remark              string      `json:"remark"`
	ExternalRef         string      `json:"external_ref"` //order id
	CreateBy            string      `json:"create_by"`
	CreateDtm           time.Time   `json:"create_dtm"`
	UpdateBy            string      `json:"update_by"`
	UpdateDtm           time.Time   `json:"update_dtm"`
	OrderItems          []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
}

func (Order) TableName() string { return "order" }

type OrderItem struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OrderID           uuid.UUID `gorm:"type:uuid" json:"order_id"`
	OrderItem         string    `json:"order_item"`
	DocumentRefItem   string    `json:"document_ref_item"`
	ProductCode       string    `json:"product_code"`
	ProductType       string    `json:"product_type"`
	InterfaceOrderQty float64   `json:"interface_order_qty"`
	Qty               float64   `json:"qty"`
	UnitCode          string    `json:"unit_code"`
	IsFocGwp          bool      `json:"is_foc_gwp"`
	WarehouseCode     string    `json:"warehouse_code"`
	BatchNo           string    `json:"batch_no"`
	SerialCode        string    `json:"serial_code"`
	Remark            string    `json:"remark"`
	ExternalItemRef   string    `json:"external_item_ref"`
	Status            string    `json:"status"`
	CreateBy          string    `json:"create_by"`
	CreateDtm         time.Time `json:"create_dtm"`
	UpdateBy          string    `json:"update_by"`
	UpdateDtm         time.Time `json:"update_dtm"`
}

func (OrderItem) TableName() string { return "order_item" }

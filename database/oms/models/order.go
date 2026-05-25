package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID                          uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
	OrderCode                   string      `gorm:"uniqueIndex" json:"order_code"`
	OrderType                   string      `json:"order_type"`
	OrderDate                   *time.Time  `json:"order_date"`
	ShopCode                    string      `json:"shop_code"`
	TenantID                    uuid.UUID   `gorm:"type:uuid" json:"tenant_id"`
	CustomerCode                string      `json:"customer_code"`
	ShipToCode                  string      `json:"ship_to_code"`
	BillToCode                  string      `json:"bill_to_code"`
	TransportZone               string      `json:"transport_zone"`
	InterfaceQty                float64     `json:"interface_qty"`
	InterfaceUnitCode           string      `json:"interface_unit_code"`
	Qty                         float64     `json:"qty"`
	UnitCode                    string      `json:"unit_code"`
	EstimatePickingDate         *time.Time  `json:"estimate_picking_date"`
	DeliveryDate                *time.Time  `json:"delivery_date"`
	SubmitDate                  *time.Time  `json:"submit_date"`
	Status                      string      `json:"status"`
	StatusShipment              string      `json:"status_shipment"`
	DocumentRefType             string      `json:"document_ref_type"`
	DocumentRef                 string      `json:"document_ref"`
	DocumentRefType2            string      `json:"document_ref_type2"`
	DocumentRef2                string      `json:"document_ref2"`
	PaymentMethod               string      `json:"payment_method"`
	TotalDiscount               float64     `json:"total_discount"`
	Discount                    float64     `json:"discount"`
	TotalVoucherAmount          float64     `json:"total_voucher_amount"`
	PromisedShippingTimes       string      `json:"promised_shipping_times"`
	PromisedShippingType        string      `json:"promised_shipping_type"`
	TotalAmount                 float64     `json:"total_amount"`
	ShippingFee                 float64     `json:"shipping_fee"`
	Currency                    string      `json:"currency"`
	ShippingFeeOriginal         float64     `json:"shipping_fee_original"`
	ShippingFeeDiscountSeller   float64     `json:"shipping_fee_discount_seller"`
	ShippingFeeDiscountPlatform float64     `json:"shipping_fee_discount_platform"`
	PhoneNo                     string      `json:"phone_no"`
	Address                     string      `json:"address"`
	Country                     string      `json:"country"`
	District                    string      `json:"district"`
	Email                       string      `json:"email"`
	PostalCode                  string      `json:"postal_code"`
	Province                    string      `json:"province"`
	RecipientName               string      `json:"recipient_name"`
	SubDistrict                 string      `json:"sub_district"`
	PromisedRtsDtm              *time.Time  `json:"promised_rts_dtm"`
	PromisedDeliveryDtm         *time.Time  `json:"promised_delivery_dtm"`
	ShipAddress                 string      `json:"ship_address"`
	InterfaceCreateDtm          *time.Time  `json:"interface_create_dtm"`
	InterfaceUpdateDtm          *time.Time  `json:"interface_update_dtm"`
	InvoiceNo                   string      `json:"invoice_no"`
	Remark                      string      `json:"remark"`
	ExternalRef                 string      `json:"external_ref"`
	CreateBy                    string      `json:"create_by"`
	CreateDtm                   time.Time   `json:"create_dtm"`
	UpdateBy                    string      `json:"update_by"`
	UpdateDtm                   time.Time   `json:"update_dtm"`
	OrderItems                  []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
}

func (Order) TableName() string { return "order" }

type OrderItem struct {
	ID                          uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	OrderID                     uuid.UUID  `gorm:"type:uuid" json:"order_id"`
	OrderItem                   string     `json:"order_item"`
	DocumentRefItem             string     `json:"document_ref_item"`
	ProductCode                 string     `json:"product_code"`
	ProductType                 string     `json:"product_type"`
	InterfaceOrderQty           float64    `json:"interface_order_qty"`
	Qty                         float64    `json:"qty"`
	UnitCode                    string     `json:"unit_code"`
	PriceUnit                   float64    `json:"price_unit"`
	InterfaceQty                float64    `json:"interface_qty"`
	VatAmount                   float64    `json:"vat_amount"`
	VoucherPlatform             float64    `json:"voucher_platform"`
	VoucherSeller               float64    `json:"voucher_seller"`
	TotalVoucher                float64    `json:"total_voucher"`
	ShippingFeeOriginal         float64    `json:"shipping_fee_original"`
	ShippingFeeDiscountSeller   float64    `json:"shipping_fee_discount_seller"`
	ShippingFeeDiscountPlatform float64    `json:"shipping_fee_discount_platform"`
	ShippingMethod              string     `json:"shipping_method"`
	ShippingCourier             string     `json:"shipping_courier"`
	PlatformCarrierID           string     `json:"platform_carrier_id"`
	ShippingFee                 float64    `json:"shipping_fee"`
	ProductDesc                 string     `json:"product_desc"`
	ShippingDesc                string     `json:"shipping_desc"`
	RtsDtm                      *time.Time `json:"rts_dtm"`
	IsCod                       bool       `json:"is_cod"`
	TotalAmount                 float64    `json:"total_amount"`
	IsFocGwp                    bool       `json:"is_foc_gwp"`
	WarehouseCode               string     `json:"warehouse_code"`
	BatchNo                     string     `json:"batch_no"`
	SerialCode                  string     `json:"serial_code"`
	Remark                      string     `json:"remark"`
	ExternalItemRef             string     `json:"external_item_ref"`
	Status                      string     `json:"status"`
	CreateBy                    string     `json:"create_by"`
	CreateDtm                   time.Time  `json:"create_dtm"`
	UpdateBy                    string     `json:"update_by"`
	UpdateDtm                   time.Time  `json:"update_dtm"`
}

func (OrderItem) TableName() string { return "order_item" }

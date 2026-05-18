package models

import (
	"time"

	"github.com/google/uuid"
)

type Shop struct {
	ID                  uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	ShopCode            string     `json:"shop_code"`
	ShopName            string     `json:"shop_name"`
	Email               string     `json:"email"`
	PlatformCode        string     `json:"platform_code"`
	ExternalShopID      string     `gorm:"type:text" json:"external_shop_id"`
	IsAutoReconnectAuth bool       `json:"is_auto_reconnect_auth"`
	IsActive            bool       `json:"is_active"`
	AutoDownloadDataFlg bool       `json:"auto_download_data_flg"`
	LastSyncedDtm       *time.Time `json:"last_synced_dtm"`
	SyncStatusEnabled   bool       `json:"sync_status_enabled"`
	IsSelfFulfilled     bool       `json:"is_self_fulfilled"`
	IsDefaultCourier    string     `json:"is_default_courier"`
	AccessToken         string     `json:"access_token"`
	RefreshToken        string     `json:"refresh_token"`
	APIKey              string     `json:"api_key"`
	CreateBy            string     `json:"create_by"`
	CreateDtm           time.Time  `json:"create_dtm"`
	UpdateBy            string     `json:"update_by"`
	UpdateDtm           time.Time  `json:"update_dtm"`
	ShopAuth            []ShopAuth `json:"shop_auth"`
}

func (Shop) TableName() string { return "shop" }

type ShopAuth struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ShopID         uuid.UUID `gorm:"type:uuid" json:"shop_id"`
	ExpireDtm      time.Time `json:"expire_dtm"`
	TokenType      string    `json:"token_type"`
	AccessToken    string    `json:"access_token"`
	RefreshToken   string    `json:"refresh_token"`
	LastRefreshDtm time.Time `json:"last_refresh_dtm"`
	Scope          string    `json:"scope"`
	ExpireIn       int64     `json:"expire_in"`
	CreateDtm      time.Time `json:"create_dtm"`
	UpdateDtm      time.Time `json:"update_dtm"`
}

func (ShopAuth) TableName() string { return "shop_auth" }

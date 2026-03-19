package models

import "github.com/google/uuid"

type SystemConfig struct {
	TopicCode  string    `gorm:"primaryKey;column:topic_code" json:"topic_code"`
	ConfigCode string    `gorm:"primaryKey;column:config_code" json:"config_code"`
	TenantID   uuid.UUID `json:"tenant_id"`
	ConfigName string    `json:"config_name"`
	Cond1      string    `json:"cond1"`
	Cond2      string    `json:"cond2"`
	Value      string    `json:"value"`
	Sequence   int       `json:"sequence"`
	Remark     string    `json:"remark"`
}

func (SystemConfig) TableName() string {
	return "system_config"
}

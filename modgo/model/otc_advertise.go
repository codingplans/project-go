package model

import "time"

type OtcAdvertise struct {
	CoinID       int64     `gorm:"column:coin_id" json:"coin_id"`
	CurrencyId   int64     `gorm:"column:currency_id" json:"currency_id"`
	Type         int64     `gorm:"column:type" json:"type"`
	DealAmount   float64   `gorm:"column:deal_amount" json:"deal_amount"`
	ID           int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	MaxLimit     float64   `gorm:"column:max_limit" json:"max_limit"`
	MinLimit     float64   `gorm:"column:min_limit" json:"min_limit"`
	OwnerID      int64     `gorm:"column:owner_id" json:"owner_id"`
	PayMode      string    `gorm:"column:pay_mode" json:"pay_mode"`
	Remark       string    `gorm:"column:remark" json:"remark"`
	AutoReply    string    `gorm:"column:auto_reply" json:"auto_reply"`
	Price        float64   `gorm:"column:price" json:"price"`
	RemainAmount float64   `gorm:"column:remain_amount" json:"remain_amount"`
	Status       int64     `gorm:"column:status" json:"status"`
	Timeout      int64     `gorm:"column:timeout" json:"timeout"`
	CreateAt     time.Time `gorm:"column:create_at" json:"create_at"`
	TotalAmount  float64   `gorm:"column:total_amount" json:"total_amount"`
	Version      int64     `gorm:"column:version" json:"version"`
	IsTrade      int64     `gorm:"column:is_trade" json:"is_trade"`
}

// TableName sets the insert table name for this struct type
func (c *OtcAdvertise) TableName() string {
	return "otc_advertise"
}

package model

import "time"

type Coin struct {
	ID                int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	ColdWalletAddress string    `gorm:"column:cold_wallet_address" json:"cold_wallet_address"`
	CreatedTime       time.Time `gorm:"column:created_time" json:"created_time"`
	FullName          string    `gorm:"column:full_name" json:"full_name"`
	HasLegal          int64     `gorm:"column:has_legal" json:"has_legal"`
	MaxTxFee          float64   `gorm:"column:max_tx_fee" json:"max_tx_fee"`
	MaxWithdrawAmount float64   `gorm:"column:max_withdraw_amount" json:"max_withdraw_amount"`
	MinTxFee          float64   `gorm:"column:min_tx_fee" json:"min_tx_fee"`
	MinWithdrawAmount float64   `gorm:"column:min_withdraw_amount" json:"min_withdraw_amount"`
	MinerFee          float64   `gorm:"column:miner_fee" json:"miner_fee"`
	Sort              int64     `gorm:"column:sort" json:"sort"`
	Status            int64     `gorm:"column:status" json:"status"`
	Symbol            string    `gorm:"column:symbol" json:"symbol"`
	UpdatedTime       time.Time `gorm:"column:updated_time" json:"updated_time"`
	UsdRate           float64   `gorm:"column:usd_rate" json:"usd_rate"`
	WithdrawScale     int64     `gorm:"column:withdraw_scale" json:"withdraw_scale"`
	WithdrawThreshold float64   `gorm:"column:withdraw_threshold" json:"withdraw_threshold"`
}

// TableName sets the insert table name for this struct type
func (c *Coin) TableName() string {
	return "coin"
}

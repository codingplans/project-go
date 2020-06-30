package model

import "time"

type OtcAsset struct {
	Balance       float64   `gorm:"column:balance" json:"balance"`
	CoinID        int       `gorm:"column:coin_id" json:"coin_id"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
	FrozenBalance float64   `gorm:"column:frozen_balance" json:"frozen_balance"`
	ID            int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	IsLock        int       `gorm:"column:is_lock" json:"is_lock"`
	MemberID      int       `gorm:"column:member_id" json:"member_id"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
	Version       int       `gorm:"column:version" json:"version"`
}

// TableName sets the insert table name for this struct type
func (c *OtcAsset) TableName() string {
	return "otc_asset"
}

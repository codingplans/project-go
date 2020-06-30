package model

import "time"

type MemberTransaction struct {
	Amount        float64   `gorm:"column:amount" json:"amount"`
	HistoryAmount float64   `gorm:"column:history_amount" json:"history_amount"`
	CoinID        int       `gorm:"column:coin_id" json:"coin_id"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
	Fee           float64   `gorm:"column:fee" json:"fee"`
	ID            int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	MemberID      int64     `gorm:"column:member_id" json:"member_id"`
	RealFee       float64   `gorm:"column:real_fee" json:"real_fee"`
	Type          int       `gorm:"column:type" json:"type"`
}

// 用户交易
func (c *MemberTransaction) TableName() string {
	return "member_transaction"
}

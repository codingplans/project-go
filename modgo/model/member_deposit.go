package model

import (
	"time"
)

type MemberDeposit struct {
	Address    string    `gorm:"column:address" json:"address"`
	Amount     float64   `gorm:"column:amount" json:"amount"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	ID         int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	MemberID   int       `gorm:"column:member_id" json:"member_id"`
	Symbol     string    `gorm:"column:symbol" json:"symbol"`
	Txid       string    `gorm:"column:txid" json:"txid"`
}

// 用户提现
func (m *MemberDeposit) TableName() string {
	return "member_deposit"
}

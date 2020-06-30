package model

import "time"

type MemberAccount struct {
	Balance      float64   `gorm:"column:balance" json:"balance"`
	CoinID       int       `gorm:"column:coin_id" json:"coin_id"`
	CreateAt     time.Time `gorm:"column:create_at" json:"create_at"`
	DealAmount   float64   `gorm:"column:deal_amount" json:"deal_amount"`
	Fee          float64   `gorm:"column:fee" json:"fee"`
	ID           int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	MemberID     int64     `gorm:"column:member_id" json:"member_id"`
	RemainAmount float64   `gorm:"column:remain_amount" json:"remain_amount"`
	Remark       string    `gorm:"column:remark" json:"remark"`
	Status       int       `gorm:"column:status" json:"status"`
	UpdateAt     time.Time `gorm:"column:update_at" json:"update_at"`
}

// TableName sets the insert table name for this struct type
func (m *MemberAccount) TableName() string {
	return "member_account"
}

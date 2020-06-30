package model

import "time"

type WalletAddress struct {
	ID         int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	Address    string    `gorm:"column:address" json:"address"`
	CoinID     int64     `gorm:"column:coin_id" json:"coin_id"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	MemberID   int64     `gorm:"column:member_id" json:"member_id"`
	ProviderID string    `gorm:"column:provider_id" json:"provider_id"`
}

// TableName sets the insert table name for this struct type
func (w *WalletAddress) TableName() string {
	return "wallet_address"
}

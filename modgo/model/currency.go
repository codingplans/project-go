package model

import "time"

type Currency struct {
	Code          string    `gorm:"column:code" json:"code"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	FullName      string    `gorm:"column:full_name" json:"full_name"`
	ID            int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	Precision     int64     `gorm:"column:precision" json:"precision"`
	Sort          int64     `gorm:"column:sort" json:"sort"`
	Status        int64     `gorm:"column:status" json:"status"`
	Symbol        string    `gorm:"column:symbol" json:"symbol"`
	UsdToThisRate float64   `gorm:"column:usd_to_this_rate" json:"usd_to_this_rate"`
}

// TableName sets the insert table name for this struct type
func (c *Currency) TableName() string {
	return "currency"
}

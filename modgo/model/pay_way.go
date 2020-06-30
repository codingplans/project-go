package model

type PayWay struct {
	ID     int64  `gorm:"column:id;primary_key" json:"id;primary_key"`
	Name   string `gorm:"column:name" json:"name"`
	EnName string `gorm:"column:en_name" json:"en_name"`
	Sort   int    `gorm:"column:sort" json:"sort"`
	Status int    `gorm:"column:status" json:"status"`
}

// TableName sets the insert table name for this struct type
func (p *PayWay) TableName() string {
	return "pay_way"
}

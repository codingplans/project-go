package model

type Country struct {
	// Id       int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	ID       int64  `gorm:"auto_increment;primary_key;column:id" json:"id" `
	AreaCode int64  `gorm:"column:area_code" json:"area_code"`
	Currency string `gorm:"column:currency" json:"currency"`
	Language string `gorm:"column:language" json:"language"`
	Name     string `gorm:"column:name" json:"name"`
	Sort     int64  `gorm:"column:sort" json:"sort"`
}

// TableName sets the insert table name for this struct type
func (c *Country) TableName() string {
	return "country"
}

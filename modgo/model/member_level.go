package model

type MemberLevel struct {
	BeforeID int64  `gorm:"column:before_id" json:"before_id"`
	ID       int64  `gorm:"column:id;primary_key" json:"id;primary_key"`
	MemberID int64  `gorm:"column:member_id" json:"member_id"`
	Records  string `gorm:"column:records" json:"records"`
}

// TableName sets the insert table name for this struct type
func (m *MemberLevel) TableName() string {
	return "member_level"
}

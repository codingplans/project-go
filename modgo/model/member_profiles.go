package model

import (
	"time"
)

type MemberProfiles struct {
	BackImage    string    `gorm:"column:back_image" json:"back_image"`
	BankCardAuth string    `gorm:"column:bank_card_auth" json:"bank_card_auth"`
	BoundNum     string    `gorm:"column:bound_num" json:"bound_num"`
	CountryID    int64     `gorm:"column:country_id" json:"country_id"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	FontImage    string    `gorm:"column:font_image" json:"font_image"`
	ID           int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	Kyc1Auth     int64     `gorm:"column:kyc1_auth" json:"kyc1_auth"`
	Kyc2Auth     int64     `gorm:"column:kyc2_auth" json:"kyc2_auth"`
	Phone        string    `gorm:"column:phone" json:"phone"`
	RealName     string    `gorm:"column:real_name" json:"real_name"`
	Remark       string    `gorm:"column:remark" json:"remark"`
	Status       int64     `gorm:"column:status" json:"status"`
	Type         int64     `gorm:"column:type" json:"type"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID       int64     `gorm:"column:user_id" json:"user_id"`
}

// 用户交易方式
func (m *MemberProfiles) TableName() string {
	return "member_profiles"
}

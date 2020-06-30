package model

import (
	"time"
)

type MemberCapital struct {
	BackImage string    `gorm:"column:back_image" json:"back_image"`
	CreateAt  time.Time `gorm:"column:create_at" json:"create_at"`
	FontImage string    `gorm:"column:font_image" json:"font_image"`
	HandImage string    `gorm:"column:hand_image" json:"hand_image"`
	ID        int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	Password  string    `gorm:"column:password" json:"password"`
	RealName  string    `gorm:"column:real_name" json:"real_name"`
	BoundNum  string    `gorm:"column:bound_num" json:"bound_num"`
	Type      int64     `gorm:"column:type" json:"type"`
	UserID    int64     `gorm:"column:user_id" json:"user_id"`
}

// 用户实名制
func (m *MemberCapital) TableName() string {
	return "member_capital"
}

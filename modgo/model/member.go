package model

import "time"

type Member struct {
	ID        int64     `gorm:"column:id;primary_key" json:"id;primary_key"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	Country   int64     `gorm:"column:country" json:"country"`
	CreateAt  time.Time `gorm:"column:create_at" json:"create_at"`
	Email     string    `gorm:"column:email" json:"email"`
	Password  string    `gorm:"column:password" json:"password"`
	Token     string    `gorm:"column:token" json:"token"`
	EmailAuth int64     `gorm:"column:email_auth" json:"email_auth"`
	Mobile    string    `gorm:"column:mobile" json:"mobile"`
	Sex       int64     `gorm:"column:sex" json:"sex"`
	Status    int64     `gorm:"column:status" json:"status"`
	UpdateAt  time.Time `gorm:"column:update_at" json:"update_at"`
	Username  string    `gorm:"column:username" json:"username"`
}

// TableName sets the insert table name for this struct type
func (c *Member) TableName() string {
	return "member"
}

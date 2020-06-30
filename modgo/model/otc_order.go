package model

type OtcOrder struct {
	AdvertiseType int64   `gorm:"column:advertise_type" json:"advertise_type"`
	Amount        float64 `gorm:"column:amount" json:"amount"`
	Commission    float64 `gorm:"column:commission" json:"commission"`
	CreateAt      int64   `gorm:"column:create_at" json:"create_at"`
	OrderID       string  `gorm:"column:order_id;primary_key" json:"order_id;primary_key"`
	PayImage      string  `gorm:"column:pay_image" json:"pay_image"`
	PayInfo       string  `gorm:"column:pay_info" json:"pay_info"`
	PayTime       int64   `gorm:"column:pay_time" json:"pay_time"`
	FinishTime    int64   `gorm:"column:finish_time" json:"finish_time"`
	Price         float64 `gorm:"column:price" json:"price"`
	SponsorID     int64   `gorm:"column:sponsor_id" json:"sponsor_id"`
	Status        int64   `gorm:"column:status" json:"status"`
	Total         float64 `gorm:"column:total" json:"total"`
	TraderID      int64   `gorm:"column:trader_id" json:"trader_id"`
	AdvertiseId   int64   `gorm:"column:advertise_id" json:"advertise_id"`
}

// TableName sets the insert table name for this struct type
func (o *OtcOrder) TableName() string {
	return "otc_order"
}

// order status: 1 创建订单 2 已打款  3确认收到  4未收到，申诉   5结束订单 6完成订单

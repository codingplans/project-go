package main

import (
	"context"
	"flag"
	"fmt"
	sa "github.com/sensorsdata/sa-sdk-go"
	"iceberg/frame/icelog"
	"laoyuegou.com/shence"
	"plorder/config"
	"time"
)

type PLOrder struct {
	ShenceClient *shence.Client
	shence       sa.SensorsAnalytics
	ctx          context.Context
	cancel       context.CancelFunc
}

func main() {
	// TestRedis()
	testshence()
}

func testshence() {
	sc := connShence()
	// sc.ShenceClient.EventChan
	// fmt.Printf("query cost %d millisecond.\n", sc)
	icelog.Infof("%+v,&&&&&& %+v", sc)
	err := sc.shence.Track(fmt.Sprint(12312), "complete", map[string]interface{}{
		"ID":              123123,
		"ordertype":       123123,
		"class":           123123,
		"directional":     1,
		"placetime":       time.Unix(int64(123123), 0).Format("2006-01-02 15:04:05"),
		"OrderAmount":     int(123123),
		"CouponAmount":    12,
		"order_coupon_id": int(123),
		"payment":         int(123123 - 123123),
		"ordernumber":     int(123123),
		"auto":            false,
		"isLogin":         true,
		"godID":           int(123123),
		"room_id":         int(123123),
		// "pay_way":         paymentconst.Agent(oResp.GetData().GetPaymentType()).Name(),
		"godCouponok": 12,
		"SkillLv":     int(12),
		"PotentialLv": int(12),
		"godTag":      123,
		"GrabMode":    12312,
	}, true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(err)

}

// func connEs() *CacheStore {
func connShence() *PLOrder {
	cfgFile := flag.String("config-path", "/Users/darren/go/src/testgo/config.json", "config file")

	cfg := config.Init(*cfgFile)
	// fmt.Printf("%+v", cfg)
	Shence := new(PLOrder)
	Shence.ctx, Shence.cancel = context.WithCancel(context.Background())
	c, _ := sa.InitDefaultConsumer(cfg.ShenceCfg.Host, cfg.ShenceCfg.Timeout)
	Shence.shence = sa.InitSensorsAnalytics(c, cfg.ShenceCfg.Project, false)
	Shence.ShenceClient = shence.NewClient(cfg.ShenceCfg.Host,
		cfg.ShenceCfg.Project,
		cfg.ShenceCfg.Timeout, Shence.ctx)
	Shence.ShenceClient.Init()

	return Shence
}

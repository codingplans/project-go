package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var str1 = `SELECT
    sum(ad_space_req_cnt)  as '广告位请求数' , 
    sum(ad_render_cnt) as  '渲染数' , 
    sum(ad_exp_cnt) as '曝光数',
    sum(ad_clk_cnt) as '点击数',
    sum(ad_clk_cnt)/sum(ad_req_cnt) as '点击率'
    sum(ad_settlement_income) as '预估收入' ,
    sum(ad_settlement_income)/sum(ad_exp_cnt)  as 'ecpm',
    sum(ad_settlement_income)/count(distinct(source_uid)) as 'arpu',
    COUNT(DISTINCT(ad_bid_succ_uid))AS '竞胜UV' 
    COUNT(DISTINCT(ad_req_uid)) AS '请求UV'
    COUNT(DISTINCT(ad_exp_uid)) AS '曝光UV'
    COUNT(DISTINCT(ad_clk_uid)) AS '点击UV'
    SUM(ad_resp_cnt)/SUM(ad_req_cnt) as '广告位填充率'
    SUM(ad_bid_succ_uid) as '客户端竞胜数'
    SUM(ad_exp_cnt)/SUM(ad_resp_cnt) as '广告位曝光率'
    SUM(ad_inchapter_strength) as '插页广告强度'
    SUM(ad_prerolls_strength) as '强插广告强度'
    SUM(ad_bottom_empty_window_time) as '底部空窗期'
    SUM(ad_video_reward) AS '激励视频奖励获取'
`
var str = `FROM
    viw.viw_cc_ad_user_action_report_inc_d
where`
var str3 = `dt = '2022-11-20'
group by`
var str4 = ` dt,
    hour,
    project,
    ad_unit_id, 
    ad_format,
    ad_scene,
    os,
    app_version_code,
    source_channel,
    ad_layout,
    creative_type
`

func BenchmarkAddStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = str + str1 + str3 + str4
	}
}

func BenchmarkSprintStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s%s", str, str1, str3, str4)
	}
}

func BenchmarkJoinStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{str, str1, str3, str4}, "")
	}
}

func BenchmarkBufferStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var by bytes.Buffer
		by.WriteString(str1)
		by.WriteString(str)
		by.WriteString(str3)
		by.WriteString(str4)
		_ = by.String()
	}
}

// BenchmarkAddStringWithOperator-8            50000000             30.3 ns/op
// BenchmarkAddStringWithSprintf-8             5000000              261  ns/op
// BenchmarkAddStringWithJoin-8                30000000             58.7 ns/op
// BenchmarkAddStringWithBuffer-8              2000000000           0.00 ns/op

// go test -bench='Str$' -run=none -cpu=8
// goos: darwin
// goarch: amd64
// pkg: testgo/modgo/test
// cpu: Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz
// BenchmarkAddStr-8      	172992486	         7.022 ns/op
// BenchmarkSprintStr-8   	 8747796	       133.8 ns/op
// BenchmarkJoinStr-8     	33202748	        36.98 ns/op
// BenchmarkBufferStr-8   	23436049	        50.49 ns/op

// go test -bench='Str$' -run=none -cpu=8
// goos: darwin
// goarch: arm64
// pkg: testgo/modgo/test
// BenchmarkAddStr-8        5818599               188.2 ns/op
// BenchmarkSprintStr-8     4361670               278.8 ns/op
// BenchmarkJoinStr-8       6267175               193.4 ns/op
// BenchmarkBufferStr-8     1880974               635.9 ns/op
// PASS
// ok      testgo/modgo/test       6.643s

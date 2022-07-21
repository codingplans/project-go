package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var engine *xorm.Engine
var rsEngine *xorm.Engine

type rsUser struct {
	EmpNo        string `json:"emp_no"`
	Name         string `json:"name"`
	Mobile       string `json:"mobile"`
	IdCard       string `json:"id_card"`
	PositionName string `json:"position_name,omitempty"`
	Nature       int    `json:"nature"` // 网点类型1：加盟，2：直营

}

type user struct {
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	NodeCode    string `json:"node_code"`
	MobilePhone string `json:"mobile_phone"`
	NodeId      string `json:"node_id"`
	Nature      int    `json:"nature"` // 网点类型1：加盟，2：直营
	FullName    string `json:"full_name"`
	NodeCode2   string `json:"node_code_2"`
	Name        string `json:"name"`
	IdNumber    string `json:"id_number"`
	EmpNumber   string `json:"emp_number"`
}

type mapUser struct {
	Users              []user   `json:"users"`
	RsUsers            []rsUser `json:"rs_users"`
	MobUsers           []rsUser `json:"mobile_users"`
	JiamengUserCount   int      `json:"jiameng_user_count"`    // 加盟数量
	UserCount          int      `json:"user_count"`            // 直营数量
	RsJiamengUserCount int      `json:"rs_jiameng_user_count"` // 加盟人事数量
	RsUserCount        int      `json:"rs_user_count"`         // 直营人事数量
}

// 人事表用户集合
var RsUsByIdCard = make(map[string][]rsUser, 0)

// 用户表 集合
var UsByIdCardMobile = make(map[string]mapUser, 0)

// 用户表 集合
var UsByMobile = make(map[string]mapUser, 0)

// 用户表 集合
var UsByIdCard = make(map[string]mapUser, 0)

func main() {
	// WriteOneCsv()
	// WriteTwoCsv()
	// WriteThreeCsv()

}

func WriteOneCsv() {

	// 不存在则创建;存在则清空;读写模式;
	file, err := os.Create("user_idCard_list.csv")
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	// 延迟关闭
	defer file.Close()
	// 写入UTF-8 BOM，防止中文乱码
	file.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(file)
	w.Write([]string{"身份证号", "是否有人事账号", "人事直营数量", "人事加盟数量", "直营数量", "加盟数量"})
	for idCard, us := range UsByIdCard {

		hasRs := "否"
		if us.RsJiamengUserCount+us.RsUserCount > 0 {
			hasRs = "是"
		}
		w.Write([]string{
			idCard,
			hasRs,
			fmt.Sprintf("%d", us.RsUserCount),
			fmt.Sprintf("%d", us.RsJiamengUserCount),
			fmt.Sprintf("%d", us.UserCount),
			fmt.Sprintf("%d", us.JiamengUserCount),
		})
		// 刷新缓冲
		w.Flush()
	}

	fmt.Println("1写入数据完毕")

}
func WriteTwoCsv() {
	// 不存在则创建;存在则清空;读写模式;
	file, err := os.Create("idCard_mobile_list.csv")
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	// 延迟关闭
	defer file.Close()
	// 写入UTF-8 BOM，防止中文乱码
	file.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(file)
	w.Write([]string{"身份证号", "手机号", "人事直营数量", "人事加盟数量", "直营数量", "加盟数量"})
	for idCardMob, us := range UsByIdCardMobile {
		s := strings.Split(idCardMob, "_")
		idcard, mob := s[0], s[1]
		w.Write([]string{
			idcard,
			mob,
			fmt.Sprintf("%d", us.RsUserCount),
			fmt.Sprintf("%d", us.RsJiamengUserCount),
			fmt.Sprintf("%d", us.UserCount),
			fmt.Sprintf("%d", us.JiamengUserCount),
		})
		// 刷新缓冲
		w.Flush()
	}

	fmt.Println("2写入数据完毕")

}
func WriteThreeCsv() {

	// 不存在则创建;存在则清空;读写模式;
	file, err := os.Create("mobile_list.csv")
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	// 延迟关闭
	defer file.Close()
	// 写入UTF-8 BOM，防止中文乱码
	file.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(file)
	w.Write([]string{"手机号", "是否有人事账号", "人事直营数量", "人事加盟数量", "直营数量", "加盟数量"})
	for mobile, us := range UsByMobile {

		hasRs := "否"
		if us.RsJiamengUserCount+us.RsUserCount > 0 {
			hasRs = "是"
		}
		w.Write([]string{
			mobile,
			hasRs,
			fmt.Sprintf("%d", us.RsUserCount),
			fmt.Sprintf("%d", us.RsJiamengUserCount),
			fmt.Sprintf("%d", us.UserCount),
			fmt.Sprintf("%d", us.JiamengUserCount),
		})
		// 刷新缓冲
		w.Flush()
	}

	fmt.Println("3写入数据完毕")

}

func init() {
	var err error

	// 用户表
	engine, err = xorm.NewEngine("mysql", "******?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	engine.ShowSQL(true)
	fmt.Println(engine.DB().Ping(), "mysql 连接成功")

	// 人事表
	pgsource := "******?sslmode=disable"
	rsEngine, err = xorm.NewEngine("postgres", pgsource)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	rsEngine.ShowSQL(true)
	fmt.Println(rsEngine.DB().Ping(), "psql 连接成功")

}

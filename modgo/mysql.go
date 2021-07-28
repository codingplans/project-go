package main

import (
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"os"
	"strings"
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
	GetRenshiData()
	GetUserData()
	WriteOneCsv()
	WriteTwoCsv()
	WriteThreeCsv()

}

// 读取用户数据
func GetUserData() {
	var err error

	sql := "select t_user.username, t_user.nickname, t_user.code as node_code, t_user.mobile_phone,node_id, t_node_new.nature, t_node_new.full_name, " +
		"t_node_new.code as node_code2,t_idcard.name, t_idcard.id_number, emp_number from t_user left join t_user_cert on t_user.id = t_user_cert." +
		"id left join t_idcard on t_user_cert.idcard_id = t_idcard.id left join t_node_new on t_user.node_id = t_node_new.id where t_user.username is not null and t_user." +
		"disabled = 0 limit 10000;"
	us := make([]user, 0)

	err = rsEngine.SQL(sql).Find(&us)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if len(us) == 0 {
		fmt.Println("遍历结束 当前数量：")
		return
	}

	for _, v := range us {
		if v.IdNumber != "" {
			tUser := mapUser{}
			if vv, ok := UsByIdCard[v.IdNumber]; ok {
				tUser = vv
			}
			tUser.Users = append(tUser.Users, v)
			if v.Nature == 2 {
				tUser.UserCount += 1
			}
			if v.Nature == 1 {
				tUser.JiamengUserCount += 1
			}

			if vUsers, ok := RsUsByIdCard[v.IdNumber]; ok {
				for _, vUser := range vUsers {
					if vUser.Nature == 2 {
						tUser.RsUserCount += 1
					}
					if vUser.Nature == 1 {
						tUser.RsJiamengUserCount += 1
					}
				}
			}
			UsByIdCard[v.IdNumber] = tUser

		}

		// 第二条
		if v.MobilePhone != "" && v.IdNumber != "" {
			k := fmt.Sprintf("%s_%s", v.IdNumber, v.MobilePhone)
			tmep2 := mapUser{}
			if vv, ok := UsByIdCardMobile[k]; ok {
				tmep2 = vv
				tmep2.Users = append(tmep2.Users, v)
			} else {
				tmep2.Users = []user{v}
			}

			if v.Nature == 1 {
				tmep2.JiamengUserCount += 1
			}
			if v.Nature == 2 {
				tmep2.UserCount += 1
			}
			UsByIdCardMobile[k] = tmep2
		}

		// 第三条
		if v.MobilePhone != "" {
			tUser := mapUser{}
			if vv, ok := UsByMobile[v.MobilePhone]; ok {
				tUser = vv
			}
			tUser.Users = append(tUser.Users, v)
			if v.Nature == 2 {
				tUser.UserCount += 1
			}
			if v.Nature == 1 {
				tUser.JiamengUserCount += 1
			}
			UsByMobile[v.MobilePhone] = tUser
		}
	}
	fmt.Println("用户数据读取完毕")
}

func GetRenshiData() {
	var err error
	sql := "select emp_no, base_ec_employee.name, mobile, id_card," +
		"base_ec_position.name as position_name from base_ec_employee left join base_ec_position on position_id = base_ec_position.id where (base_ec_employee.status != 0 and base_ec_employee.status != 4);"

	sss := make([]rsUser, 0)
	err = engine.SQL(sql).Find(&sss)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// /NCxxxxx 为直营
	for _, v := range sss {
		if strings.HasPrefix(v.EmpNo, "NC") {
			// fmt.Println("直营")
			v.Nature = 2
		} else {
			// fmt.Println("加盟")
			v.Nature = 1
		}
		// 走第一条
		if v.IdCard != "" {
			tmp := []rsUser{v}
			if vv, ok := RsUsByIdCard[v.IdCard]; ok {
				tmp = append(vv, v)
			}
			RsUsByIdCard[v.IdCard] = tmp

		}

		// 走第二条
		if v.Mobile != "" && v.IdCard != "" {
			tmep2 := mapUser{}
			k := fmt.Sprintf("%s_%s", v.IdCard, v.Mobile)
			if vv, ok := UsByIdCardMobile[k]; ok {
				tmep2 = vv
				tmep2.RsUsers = append(tmep2.RsUsers, v)
			} else {
				tmep2.RsUsers = []rsUser{v}
			}

			if v.Nature == 1 {
				tmep2.RsJiamengUserCount++
			}
			if v.Nature == 2 {
				tmep2.RsUserCount++
			}
			UsByIdCardMobile[k] = tmep2
		}

		// 走第三条
		if v.Mobile != "" {
			tmep2 := mapUser{}
			if vv, ok := UsByMobile[v.Mobile]; ok {
				tmep2 = vv
				tmep2.RsUsers = append(tmep2.RsUsers, v)
			} else {
				tmep2.RsUsers = []rsUser{v}
			}

			if v.Nature == 1 {
				tmep2.RsJiamengUserCount++
			}
			if v.Nature == 2 {
				tmep2.RsUserCount++
			}
			UsByMobile[v.Mobile] = tmep2
		}

	}

	fmt.Println("人事信息读取完毕 !!!!!!!!")

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
	engine, err = xorm.NewEngine("mysql", "ztoconnect_rw:ztoconnect_rw123$Z@tcp(10.9.15.250:3306)/ztoconnect?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	engine.ShowSQL(true)
	fmt.Println(engine.DB().Ping(), "mysql 连接成功")

	// 人事表
	pgsource := "postgres://postgres:Postgres_secv5@10.9.15.164:5432/ztometa?sslmode=disable"
	rsEngine, err = xorm.NewEngine("postgres", pgsource)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	rsEngine.ShowSQL(true)
	fmt.Println(rsEngine.DB().Ping(), "psql 连接成功")

}

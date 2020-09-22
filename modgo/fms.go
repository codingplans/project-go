package main

import (
	"fmt"
	"github.com/looplab/fsm"
)

func main() {
	myFsm := fsm.NewFSM(
		"closed", // 初始化
		fsm.Events{ // 事件
			{Name: "开", Src: []string{"closed", "unlocked"}, Dst: "open"},
			{Name: "关", Src: []string{"open"}, Dst: "closed"},
			{Name: "锁", Src: []string{"unlocked", "open"}, Dst: "locked"},
			{Name: "解锁", Src: []string{"locked"}, Dst: "unlocked"},
		},
		fsm.Callbacks{}, // 回调
	)
	fmt.Println("初始化的当前状态为：", myFsm.Current())
	err := myFsm.Event("开")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("执行了“开”，当前状态为：", myFsm.Current())
	err = myFsm.Event("关")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("执行了“关”，当前状态为：", myFsm.Current())

	err = myFsm.Event("开")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("执行了“开”，当前状态为：", myFsm.Current())

	err = myFsm.Event("锁")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("执行了“锁”，当前状态为：", myFsm.Current())

	err = myFsm.Event("解锁")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("执行了“解锁”，当前状态为：", myFsm.Current())

	err = myFsm.Event("开")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("执行了“开”，当前状态为：", myFsm.Current())

	fmt.Println()
}

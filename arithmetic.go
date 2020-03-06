package main

import "fmt"

func main() {
	findJiInArr()
}

// 查找计数 func
func findJiInArr() {

	arr := []int{
		2, 2,
		3,
		4, 4, 4, 4, 4,
		5, 5, 5,
		6, 6, 6, 6}

	// temp := []
	for k, v := range arr {

		fmt.Println(k, v)
	}

}

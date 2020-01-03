package main

import (
	"fmt"
)

type kk struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

func Testmaps() {
	// temp := make([]int, 2, 122)
	// temp = append(temp, 123)

	// aa := int64(40011212)
	// bb := float32(0.12) * 100

	// fmt.Println(temp, len(temp), cap(temp))
	// discounts := map[float32]interface{}{
	// 	1:   "保持原价",
	// 	0.9: "9折优惠",
	// 	0.8: "8折优惠",
	// }
	// discounts := 90
	// var discounts map[int32]interface{}

	kk := float32(0.8)
	aa := int64(kk * 100)
	// discounts = append(discounts, map[int]string{
	// 100: "保持原价",
	// 90:  "9折优惠",
	// 80:  "8折优惠",
	// })

	fmt.Printf("%+v", aa)
	// fmt.Println((aa * int64(bb)) / 100)
}

func main() {
	Testmaps()
}

package main

import "fmt"

func main() {
	valueing2()
}

func valueing2() {
	var w interface{}
	w = (*int)(nil)
	fmt.Println(w == nil, w)
}

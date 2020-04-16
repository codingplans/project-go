package main

func main() {
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println(err)
	//		println(4487774)
	//
	//	} else {
	//		println(444888)
	//
	//		fmt.Println("fatal")
	//	}
	//}()

	defer func() {
		panic("defer panic")
	}()
	println(33333)

	defer func() {
		panic("333qweqw")

	}()

}

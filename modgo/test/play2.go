package main

import "github.com/samber/lo"

type bazzz struct {
	bar int
	foo string
}

type fooStruct struct {
	a int
	b int
	c string
}

func main3() {

	var arr []*bazzz
	arr = append(arr, &bazzz{bar: 1, foo: "1"})
	arr = append(arr, &bazzz{bar: 2, foo: "2"})
	arr = append(arr, &bazzz{bar: 3, foo: "3"})
	arr = append(arr, &bazzz{bar: 4, foo: "4"})
	arr = append(arr, &bazzz{bar: 5, foo: "5"})

	// 有没有一种 封装方式 可以做以下for循环中的事情
	f := &fooStruct{c: "foo"}
	for i := 0; i < 5; i++ {
		tmp, _ := f.fooFunc(arr[i], i)
		arr = append(arr, tmp)
	}
	_ = arr

	// 	如果有这种封装方式  我们姑且叫他 FormatFunc ,他把上面for循环中的东西都做掉了。 最后的结果应该是这样的
	// arr = FormatFunc(arr, fooFunc)
	// _ = arr

	// 实现：
	ff := &fooStruct{c: "foo"}
	arr = lo.FilterMap[*bazzz, *bazzz](arr, ff.fooFunc)
	_ = arr

}

func (v *fooStruct) fooFunc(a *bazzz, index int) (*bazzz, bool) {
	a.bar += index
	a.foo += v.c
	return a, true
}

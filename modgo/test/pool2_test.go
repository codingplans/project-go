package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	var sp = sync.Pool{
		// Tip: 声明对象池的New函数，这里以一个简单的int为例
		New: func() interface{} {
			return 100
		},
	}
	fmt.Println(sp.Get().(int))

	sp.Put(200)
	sp.Put(300)
	sp.Put(400)
	sp.Put(500)
	sp.Put(600)

	for i := 0; i < 20; i++ {
		// fmt.Println(sp.Get().(int))

		go func() {
			fmt.Println(sp.Get().(int))
		}()
	}
}

func TestSyncPool1(t *testing.T) {
	i := 100
	var sp = sync.Pool{
		// Tip: 声明对象池的New函数，这里以一个简单的int为例
		New: func() interface{} {
			return i + 1
		},
	}
	// Tip: 从对象池里获取一个对象
	data := sp.Get().(int)
	fmt.Println(data)
	data = sp.Get().(int)
	fmt.Println(data)
	// Tip: 往对象池里放回一个对象
	sp.Put(103)

	go func() {
		data = sp.Get().(int)
		fmt.Println(data)
	}()

}
func TestSyncPool2(t *testing.T) {
	var sp = sync.Pool{
		// Tip: 声明对象池的New函数，这里以一个简单的int为例
		New: func() interface{} {
			return 100
		},
	}
	// Tip: 从对象池里获取一个对象
	data := sp.Get().(int)
	fmt.Println(data)

	sp.Put(data)
	sp.Put(data)
	sp.Put(data)
	sp.Put(data)
	data = sp.Get().(int)
	fmt.Println(data)
	// Tip: 往对象池里放回一个对象
	sp.Put(data)
}
func TestSyncPool3(t *testing.T) {
	var sp = sync.Pool{
		// Tip: 声明对象池的New函数，这里以一个简单的int为例
		New: func() interface{} {
			return 100
		},
	}
	// Tip: 往对象池里放回一个对象

	// Tip: 从对象池里获取一个对象
	data := sp.Get().(int)
	fmt.Println(data)
	sp.Put(data)
}
func TestSyncPool4(t *testing.T) {

	var sp = sync.Pool{
		// Tip: 声明对象池的New函数，这里以一个简单的int为例
		New: func() interface{} {
			return 100
		},
	}
	// Tip: 从对象池里获取一个对象
	data := sp.Get().(int)
	fmt.Println(data)
	// Tip: 往对象池里放回一个对象
	sp.Put(data)
}

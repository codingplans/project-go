package main

import (
	"iceberg/frame/icelog"
	"sync"
)

//下面的迭代会有什么问题？

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	// ch := make(chan interface{}) // 解除注释看看！
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()

		for _, elem := range set.s {
			ch <- elem
			println("Iter:", elem)
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func main() {

	icelog.Infof("%s%v  ", "ch", 1212)

	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	v := <-th.Iter()
	icelog.Infof("%s%v  ", "ch", v)
}

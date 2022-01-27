package main

import "time"

// 实现无限缓存 chan

type unlimited struct {
	// 写chan
	in chan msg
	// 读chan
	out    chan msg
	buffer []msg
}

type msg struct {
	id      int
	content string
}

var maxBuffer = 1

func UnlimitChan() {
	obj := &unlimited{
		in:     make(chan msg, maxBuffer),
		out:    make(chan msg, maxBuffer),
		buffer: make([]msg, 0, maxBuffer),
	}
	// 永动机
	go obj.Handle()
	// 生产机器
	go obj.In()
	// 消费机器
	obj.Out()

	println(len(obj.in), len(obj.out), len(obj.buffer))
	println("永动结束 完毕")

}

// 不停的涌动生产数据
func (u *unlimited) In() {
	defer println("return !!!in！！")

	i := 0
	for {
		m := msg{id: i}
		// println("in 开始！")

		u.in <- m
		// println("in 成功！")

		i++
		if i == 50 {
			println("################关闭！in")
			close(u.in)
			return
		}
		// time.Sleep(time.Second)
		time.Sleep(time.Millisecond * 1)
	}
}

// 较慢的消费数据 以达到无限大chan
func (u *unlimited) Out() {
	defer println("return !!!消费out！！")
	i := 0
	for {
		val, ok := <-u.out
		if !ok {
			println("关闭！out")
			return
		}
		i++
		println(val.id, len(u.in), len(u.out), len(u.buffer), i)
		time.Sleep(time.Millisecond * 5)
	}

}

// 永动机处理机制
func (u *unlimited) Handle() {
	defer println("return !!!handle！！")
	defer close(u.out)
loop:
	for {
		// println("**** 开始读in")

		val, ok := <-u.in
		if !ok {
			println("已关闭")
			break loop
		}
		// println("读取到in ", val.id)

		select {
		case u.out <- val:
			// println("写入成功，下一轮")
			continue
		default:
			// println("当out 阻塞了，就走下面塞给切片")

			// 	当out 阻塞了，就走下面塞给切片
		}

		u.buffer = append(u.buffer, val)

		for u.BufferLen() != 0 {
			select {
			// case val, ok := <-u.in:
			// 	if !ok {
			// 		break loop
			// 	}
			// 	u.buffer = append(u.buffer, val)
			case u.out <- u.buffer[0]:
				u.buffer = u.buffer[1:]
				if u.BufferLen() == 0 {
					u.buffer = []msg{}
				}
				println("下一轮")
				continue
			default:
				// println("当out 阻塞了，")

				// 	当out 阻塞了，就走下面
				goto loop
			}
		}
	}

	println("return 兜底消费")
	// 兜底消费
	for u.BufferLen() != 0 {
		select {
		case u.out <- u.buffer[0]:
			u.buffer = u.buffer[1:]
		}
	}

}

func (u *unlimited) Len() int {
	return len(u.out)
}

func (u *unlimited) BufferLen() int {
	return len(u.buffer)
}

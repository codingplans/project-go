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

var maxBuffer = 100

func main() {
	obj := &unlimited{
		in:     make(chan msg, maxBuffer/10),
		out:    make(chan msg, maxBuffer/10),
		buffer: make([]msg, 0, maxBuffer),
	}
	obj.In()
	obj.Out()

}

// 不停的涌动生产数据
func (u *unlimited) In() {
	for {
		i := 0
		m := msg{id: i}
		u.buffer = append(u.buffer, m)

		if len(u.in) < maxBuffer/10 {

		}
		i++
		// time.Sleep(time.Second)
		time.Sleep(time.Millisecond * 500)
	}
}

// 较慢的消费数据 以达到无限大chan
func (u *unlimited) Out() {
	defer close(u.out)
loop:
	for {
		val, ok := <-u.in
		if !ok {
			break loop
		}
		select {
		case u.out <- val:
			continue
		default:

		}
		// 当out满了 就塞到slice里
		u.buffer = append(u.buffer, val)
		for u.BufferLen() != 0 {
			select {
			case val, ok := <-u.in:
				if !ok {
					break loop
				}
				u.buffer = append(u.buffer, val)
			case u.out <- u.buffer[0]:
				u.buffer = u.buffer[1:]
				if u.BufferLen() == 0 {
					u.buffer = []msg{}
				}
			}
		}

	}

	for u.BufferLen() != 0 {
		u.out <- u.buffer[0]
		u.buffer = u.buffer[1:]
	}
	u.buffer = []msg{}
}
func (u *unlimited) Len() int {
	return len(u.buffer) + len(u.out)
}

func (u *unlimited) BufferLen() int {
	return len(u.buffer) + len(u.out)
}

func (u *unlimited) Outs() {

}

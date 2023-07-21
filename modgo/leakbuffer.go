package main

import "bytes"

func main() {
	go clientByte()
	go clientByte()
	go clientByte()
	ServerByte()
}

var freeList = make(chan *bytes.Buffer, 100)
var serverChan = make(chan *bytes.Buffer)

func clientByte() {
	for {
		var b *bytes.Buffer

		select {
		case b = <-freeList:

		default:
			b = new(bytes.Buffer)
		}
		load(b)
		serverChan <- b
	}
}

func ServerByte() {
	for {
		b := <-serverChan
		process(b)

		select {
		case freeList <- b:

		default:

		}
	}
}

// 载入新数据
func load(buffer *bytes.Buffer) {
	_ = buffer
}

// 消费数据
func process(buffer *bytes.Buffer) {
	_ = buffer
}

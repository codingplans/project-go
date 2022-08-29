package main

import (
	"bytes"
	"runtime"
)

type Buiding interface {
	Builds()
}

type House struct {
}

func (House) Builds() {
	println("House")
}

type Shop struct {
}

func (Shop) Builds() {
	println("Shop")
}

type Toilet struct {
}

func (Toilet) Builds() {
	println("Toilet")
}

// 大于2m 结构体
type BigBar struct {
	foo runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
	_   runtime.MemStats
}

var NL = []byte{'\n'}

// 用于检测无效字符 脱* 处理
func trimComments(data []byte) (data1 []byte) {
	confLines := bytes.Split(data, NL)
	for k, line := range confLines {
		confLines[k] = trimCommentsLine(line)
	}
	return bytes.Join(confLines, NL)
}

func trimCommentsLine(line []byte) []byte {
	var newLine []byte
	var i, quoteCount int
	lastIdx := len(line) - 1
	for i = 0; i <= lastIdx; i++ {
		if line[i] == '\\' {
			if i != lastIdx && (line[i+1] == '\\' || line[i+1] == '"') {
				newLine = append(newLine, line[i], line[i+1])
				i++
				continue
			}
		}
		if line[i] == '"' {
			quoteCount++
		}
		if line[i] == '#' {
			if quoteCount%2 == 0 {
				break
			}
		}
		newLine = append(newLine, line[i])
	}
	return newLine
}


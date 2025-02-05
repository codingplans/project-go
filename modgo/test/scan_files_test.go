package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func ReadAllFiles(filePath, keyword string) []string {
	resFiles := make([]string, 0, 20)
	// 查找并打印所有 .log 文件
	err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 判断文件是否以 .log 结尾
		if !info.IsDir() && strings.HasSuffix(info.Name(), keyword) {
			resFiles = append(resFiles, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", filePath, err)
		return resFiles
	}
	return resFiles
}

func TestScan(t *testing.T) {
	// 定义命令行参数
	filePath := flag.String("test.filePath", "", "文件路径")
	keywords := flag.String("test.keyword", "log", "文件后缀")

	// 解析命令行参数
	flag.Parse()

	// 检查必须的参数
	if *filePath == "" {
		fmt.Println("请提供文件路径参数 -path")
		return
	}
	// t.Log(ReadAllFiles("/Users/darren/fund/hj-golang", keywords))
	t.Log(ReadAllFiles(*filePath, *keywords))

}

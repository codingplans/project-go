package main

import (
	"fmt"
	"testing"
)

func TestSshSimple(t *testing.T) {
	username := "zhangzhy"
	// keyPath := "/Users/darren/.ssh/id_rsa.zzygmail"
	// ip := "192.168.0.205"
	keyPath := "/Users/darren/.ssh/id_rsa.zzy-m1-hj"
	ip := "earth"
	port := "22"
	client := NewSSHClient(username, keyPath, "", ip, port)
	// 1.运行远程命令
	// cmd := `find /home/zhangzhy/darren/IndicatorSentinel  -type f \( -name "*.md" -o -name "*.yaml" \)`
	cmd := `ls`
	backinfo, err := client.Run(cmd)
	if err != nil {
		fmt.Printf("failed to run shell,err=[%v]\n", err)
		return
	}
	fmt.Printf("%v back info: \n[%v]\n", cmd, backinfo)
	// // 2. 上传一文件
	filename := "Foo.txt"
	// WriteFile(filename, []byte("hello ssh\r\n"))
	// // 上传
	// n, err := client.UploadFile(filename, "/tmp/"+filename)
	// if err != nil {
	// 	fmt.Printf("upload failed: %v\n", err)
	// 	return
	// }
	// 3. 显示该文件
	// cmd = "cat " + "/tmp/" + filename
	// backinfo, err = client.Run(cmd)
	// if err != nil {
	// 	fmt.Printf("run cmd faild: %v\n", err)
	// 	return
	// }
	// fmt.Printf("%v back info: \n[%v]\n", cmd, backinfo)
	// 4. 下载该文件到本地
	n, err := client.DownloadFile("/tmp/"+filename, "fo2o.txt")
	if err != nil {
		fmt.Printf("download failed: %v\n", err)
		return
	}
	fmt.Printf("download file[%v] ok, size=[%d]\n", filename, n)
}

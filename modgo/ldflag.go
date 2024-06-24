package main

import "fmt"

var (
	version   string
	buildTime string
	osArch    string
)

// func main(){
// 	Version()
// }

func Version() {
	fmt.Printf("Version: %s\nBuilt: %s\nOS/Arch: %s\n", version, buildTime, osArch)
}

// 注入参数内容
// go run  -ldflags "-X 'main.version=0.1' -X 'main.buildTime=2022-03-25' -X 'main.osArch=darwin/amd64'"  main.go

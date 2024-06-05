package person_go

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := `ps aux | grep "filter-dispatcher" | grep zhangzhy | grep -v grep  |awk '{print $2}' |head -n 1`

	// 使用exec.Command执行shell命令
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		log.Fatalf("命令执行失败: %v", err)
	}

	// 打印输出结果
	fmt.Println(1212)
	fmt.Println(strings.TrimSpace(string(out)))

	fmt.Println(1212)
}
func main22() {
	// 创建一个cmd对象
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "-v", "grep")
	cmd3 := exec.Command("grep", "filter-dispatcher")

	// 将cmd2的标准输入连接到cmd1的标准输出
	cmd2.Stdin, _ = cmd1.StdoutPipe()
	// 将cmd3的标准输入连接到cmd2的标准输出
	cmd3.Stdin, _ = cmd2.StdoutPipe()

	// 创建一个缓冲器，用于接收输出
	outputBuffer := &strings.Builder{}

	// 将cmd3的标准输出连接到缓冲器
	cmd3.Stdout = outputBuffer

	// 依次运行这三个命令
	cmd3.Start()
	cmd2.Start()
	cmd1.Run()

	// 等待cmd3完成
	cmd3.Wait()

	// 打印输出
	fmt.Println(outputBuffer.String())
}

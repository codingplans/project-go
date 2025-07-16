package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ResourceConsumer struct {
	memoryData [][]byte
	stopCPU    chan bool
	wg         sync.WaitGroup
}

func NewResourceConsumer() *ResourceConsumer {
	return &ResourceConsumer{
		memoryData: make([][]byte, 0),
		stopCPU:    make(chan bool),
	}
}

// 消耗指定百分比的CPU，确保每个核心都达到目标负载
func (rc *ResourceConsumer) consumeCPU(targetPercent float64, duration time.Duration) {
	numCPU := runtime.NumCPU()

	// 使用更精确的时间控制，基于毫秒级别
	cycleTime := 100 * time.Millisecond // 每个周期100ms
	workTime := time.Duration(float64(cycleTime) * targetPercent / 100)

	fmt.Printf("开始消耗CPU %.1f%% 持续 %v，使用 %d 个goroutine\n", targetPercent, duration, numCPU)
	fmt.Printf("每个核心: 工作时间=%v, 周期=%v\n", workTime, cycleTime)

	for i := 0; i < numCPU; i++ {
		rc.wg.Add(1)
		go func(coreId int) {
			defer rc.wg.Done()

			// 绑定到特定CPU核心
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()

			endTime := time.Now().Add(duration)

			for time.Now().Before(endTime) {
				select {
				case <-rc.stopCPU:
					return
				default:
					cycleStart := time.Now()
					workEnd := cycleStart.Add(workTime)

					// CPU密集型工作阶段 - 使用更复杂的计算
					counter := 0
					for time.Now().Before(workEnd) {
						// 执行一些CPU密集型计算
						for j := 0; j < 10000; j++ {
							counter += j*j + j*3 + 1
						}
						// 避免编译器优化
						if counter < 0 {
							break
						}
					}

					// 精确控制周期时间
					elapsed := time.Since(cycleStart)
					if elapsed < cycleTime {
						time.Sleep(cycleTime - elapsed)
					}
				}
			}
		}(i)
	}

	// 等待指定时间后停止
	go func() {
		time.Sleep(duration)
		close(rc.stopCPU)
	}()

	rc.wg.Wait()
	// 重新创建channel以便下次使用
	rc.stopCPU = make(chan bool)
	fmt.Println("CPU消耗结束")
}

// 获取系统总内存 (macOS兼容)
func getSystemMemory() uint64 {
	// 使用sysctl命令获取系统内存
	cmd := exec.Command("sysctl", "-n", "hw.memsize")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("获取系统内存失败，使用默认值16GB: %v\n", err)
		return 16 * 1024 * 1024 * 1024 // 16GB 默认值
	}

	memStr := strings.TrimSpace(string(output))
	memBytes, err := strconv.ParseUint(memStr, 10, 64)
	if err != nil {
		fmt.Printf("解析内存大小失败，使用默认值16GB: %v\n", err)
		return 16 * 1024 * 1024 * 1024 // 16GB 默认值
	}

	return memBytes
}

// 消耗指定百分比的内存
func (rc *ResourceConsumer) consumeMemory(targetPercent float64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// 获取系统总内存
	totalMemory := getSystemMemory()
	targetMemory := uint64(float64(totalMemory) * targetPercent / 100)

	fmt.Printf("系统总内存: %.1f GB\n", float64(totalMemory)/(1024*1024*1024))
	fmt.Printf("开始消耗内存 %.1f%% (约 %d MB)\n", targetPercent, targetMemory/(1024*1024))

	// 分块分配内存，每块1MB
	blockSize := 1024 * 1024 // 1MB
	blocksNeeded := int(targetMemory / uint64(blockSize))

	rc.memoryData = make([][]byte, blocksNeeded)
	for i := 0; i < blocksNeeded; i++ {
		rc.memoryData[i] = make([]byte, blockSize)
		// 写入一些数据防止被优化，并确保内存真正被使用
		for j := 0; j < blockSize; j += 4096 {
			rc.memoryData[i][j] = byte(i % 256)
		}

		// 每分配100MB显示一次进度
		if i%100 == 0 && i > 0 {
			fmt.Printf("已分配 %d MB\n", i)
		}
	}

	runtime.ReadMemStats(&m)
	fmt.Printf("内存分配完成，当前堆内存使用: %d MB\n", m.HeapAlloc/(1024*1024))
}

// 释放内存
func (rc *ResourceConsumer) releaseMemory() {
	rc.memoryData = nil
	runtime.GC()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("内存释放完成，当前堆内存使用: %d MB\n", m.HeapAlloc/(1024*1024))
}

func main() {
	consumer := NewResourceConsumer()

	// 设置GOMAXPROCS确保使用所有CPU核心
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("当前系统CPU核心数: %d\n", runtime.NumCPU())

	// 先消耗20%的内存
	consumer.consumeMemory(20)

	fmt.Println("\n开始CPU消耗循环...")
	fmt.Println("你可以使用 'top' 或 'htop' 命令监控CPU使用情况")
	fmt.Println("每个CPU核心应该显示约30%的使用率")

	// 创建一个ticker，每5秒触发一次CPU消耗
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// 运行10次循环作为示例，可以根据需要调整或改为无限循环
	for i := 0; i < 10111; i++ {
		select {
		case <-ticker.C:
			fmt.Printf("\n=== 第 %d 次CPU消耗 ===\n", i+1)

			// 显示当前时间
			fmt.Printf("开始时间: %s\n", time.Now().Format("15:04:05"))

			consumer.consumeCPU(50, 2*time.Second)

			fmt.Printf("结束时间: %s\n", time.Now().Format("15:04:05"))

			if i < 9 {
				fmt.Printf("等待3秒后开始下一次循环...\n")
			}
		}
	}

	fmt.Println("\n程序结束，释放内存...")
	consumer.releaseMemory()

	fmt.Println("建议: 你可以在运行时使用以下命令监控:")
	fmt.Println("  macOS: top -pid $(pgrep -f 'go run') -s 1")
	fmt.Println("  或者: htop")
}

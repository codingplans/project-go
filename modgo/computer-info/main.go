package main

import (
	"fmt"
	"github.com/StackExchange/wmi"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"log"
)

// GPU 信息结构体
type Win32_VideoController struct {
	Name                 string
	DriverVersion        string
	AdapterRAM           uint64
	VideoProcessor       string
	VideoModeDescription string
}

func main() {
	// 获取 CPU 信息
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Fatalf("获取 CPU 信息失败: %v", err)
	}
	fmt.Println("===== CPU 信息 =====")
	for _, ci := range cpuInfo {
		fmt.Printf("型号: %s\n核心数: %d\n频率: %.2f MHz\n", ci.ModelName, ci.Cores, ci.Mhz)
	}

	// 获取内存信息
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalf("获取内存信息失败: %v", err)
	}
	fmt.Println("\n===== 内存信息 =====")
	fmt.Printf("总内存: %.2f GB\n可用内存: %.2f GB\n使用率: %.2f%%\n",
		float64(vmStat.Total)/1e9,
		float64(vmStat.Available)/1e9,
		vmStat.UsedPercent)

	// 获取显卡信息
	var gpuControllers []Win32_VideoController
	err = wmi.Query("SELECT Name, DriverVersion, AdapterRAM, VideoProcessor, VideoModeDescription FROM Win32_VideoController", &gpuControllers)
	if err != nil {
		log.Fatalf("获取 GPU 信息失败: %v", err)
	}
	fmt.Println("\n===== 显卡信息 =====")
	for _, gpu := range gpuControllers {
		fmt.Printf("名称: %s\n驱动版本: %s\n显存: %.2f GB\n处理器: %s\n分辨率: %s\n\n",
			gpu.Name, gpu.DriverVersion, float64(gpu.AdapterRAM)/1e9, gpu.VideoProcessor, gpu.VideoModeDescription)
	}
}

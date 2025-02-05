package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

type FileInfo struct {
	Path    string
	ModTime time.Time
	Content string
	Size    int64
	IsDir   bool
}

type FileWatcher struct {
	watcher      *fsnotify.Watcher
	done         chan bool
	directory    string
	delay        time.Duration
	eventCache   map[string]time.Time
	initialFiles []FileInfo
}

// ScanDirectory 扫描整个目录及其子目录的所有文件
func ScanDirectory(directory string) ([]FileInfo, error) {
	var allFiles []FileInfo

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("访问路径 %s 错误: %v", path, err)
			return nil // 继续扫描其他文件
		}

		fileInfo := FileInfo{
			Path:    path,
			ModTime: info.ModTime(),
			Size:    info.Size(),
			IsDir:   info.IsDir(),
		}

		// 如果是文件（不是目录），则读取内容
		if !info.IsDir() {
			content, err := readFileContent(path)
			if err != nil {
				log.Printf("警告: 读取文件 %s 失败: %v", path, err)
			} else {
				fileInfo.Content = content
			}
		}

		if IsOldFileWithTime(info.ModTime()) {
			log.Println("文件超过,直接跳出", info.Name())
			return nil
		}

		allFiles = append(allFiles, fileInfo)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("扫描目录失败: %v", err)
	}

	return allFiles, nil
}

func IsOldFileWithTime(lastTime time.Time) bool {
	// if viper.GetInt("last_days") > 0 {
	// lastTime, err := utils.GetLastModifiedTime(filePath)
	if time.Now().AddDate(0, 0, -1).After(lastTime) {
		return true
	}
	// }
	return false
}

// readFileContent 读取文件内容
func readFileContent(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// NewFileWatcher 创建新的文件监控器
func NewFileWatcher(directory string) (*FileWatcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("创建监控器失败: %v", err)
	}

	return &FileWatcher{
		watcher:    watcher,
		done:       make(chan bool),
		directory:  directory,
		delay:      200 * time.Millisecond,
		eventCache: make(map[string]time.Time),
	}, nil
}

// Start 开始监控目录
func (fw *FileWatcher) Start() error {
	// 首先检查目录是否存在
	if _, err := os.Stat(fw.directory); os.IsNotExist(err) {
		return fmt.Errorf("目录不存在: %s", fw.directory)
	}

	// 在开始监控之前，先扫描所有文件
	log.Printf("开始扫描目录: %s", fw.directory)
	allFiles, err := ScanDirectory(fw.directory)
	if err != nil {
		log.Printf("扫描目录失败: %v", err)
	} else {
		fw.initialFiles = allFiles
		// 打印扫描结果
		for _, file := range allFiles {
			if file.IsDir {
				log.Printf("找到目录: %s", file.Path)
			} else {
				log.Printf("找到文件: %s, 大小: %d bytes, 修改时间: %v",
					file.Path, file.Size, file.ModTime)
				// 可以根据需要打印文件内容
				// log.Printf("文件内容: %s", file.Content)
			}
		}
		log.Printf("扫描完成，共发现 %d 个文件和目录", len(allFiles))
	}

	// 递归添加所有子目录到监控
	err = filepath.Walk(fw.directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return fw.watcher.Add(path)
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("添加目录到监控失败: %v", err)
	}

	go fw.watchEvents()
	return nil
}

// GetInitialFiles 获取初始扫描的文件列表
func (fw *FileWatcher) GetInitialFiles() []FileInfo {
	return fw.initialFiles
}

// watchEvents 处理文件变化事件
func (fw *FileWatcher) watchEvents() {
	for {
		select {
		case event, ok := <-fw.watcher.Events:
			if !ok {
				return
			}

			if !fw.shouldProcessEvent(event.Name) {
				continue
			}

			fw.handleEvent(event)

		case err, ok := <-fw.watcher.Errors:
			if !ok {
				return
			}
			log.Printf("监控错误: %v", err)

		case <-fw.done:
			return
		}
	}
}

// shouldProcessEvent 检查是否应该处理该事件（防抖）
func (fw *FileWatcher) shouldProcessEvent(filename string) bool {
	now := time.Now()
	if lastTime, exists := fw.eventCache[filename]; exists {
		if now.Sub(lastTime) < fw.delay {
			return false
		}
	}
	fw.eventCache[filename] = now
	return true
}

// handleEvent 处理具体的文件事件
func (fw *FileWatcher) handleEvent(event fsnotify.Event) {
	switch {
	case event.Op&fsnotify.Create == fsnotify.Create:
		info, err := os.Stat(event.Name)
		if err == nil && info.IsDir() {
			fw.watcher.Add(event.Name)
			log.Printf("新建目录: %s", event.Name)
		} else {
			log.Printf("新建文件: %s", event.Name)
			if content, err := readFileContent(event.Name); err == nil {
				log.Printf("新文件内容: %s", content)
			}
		}

	case event.Op&fsnotify.Write == fsnotify.Write:
		log.Printf("文件内容修改: %s", event.Name)
		if content, err := readFileContent(event.Name); err == nil {
			log.Printf("更新后的内容: %s", content)
		}

	case event.Op&fsnotify.Remove == fsnotify.Remove:
		log.Printf("删除文件/目录: %s", event.Name)

	case event.Op&fsnotify.Rename == fsnotify.Rename:
		log.Printf("重命名文件/目录: %s", event.Name)
	}
}

// Stop 停止监控
func (fw *FileWatcher) Stop() {
	fw.done <- true
	fw.watcher.Close()
}

func main() {
	// watchDir := "/home/PublicData/Oxygen/V1/log" // 替换为实际的NAS挂载目录路径
	watchDir := "/Users/darren/fund/hj-golang/build-parse/logs" // 替换为实际的NAS挂载目录路径

	watcher, err := NewFileWatcher(watchDir)
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Stop()

	if err := watcher.Start(); err != nil {
		log.Fatal(err)
	}

	// 获取初始扫描的文件列表
	files := watcher.GetInitialFiles()
	log.Printf("初始扫描发现 %d 个文件和目录", len(files))

	// 保持程序运行
	select {}
}

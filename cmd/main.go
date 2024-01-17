package main

import (
	"filefire/internal/sync"
	"fmt"
	"os"
)

func main() {
	// 指定源目录和目标目录
	sourceDir := "/path/to/source/directory"
	targetDir := "/path/to/target/directory"
	// 执行文件同步操作
	err := sync.SyncFiles(sourceDir, targetDir)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("File sync completed.")
}

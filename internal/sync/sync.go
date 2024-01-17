package sync

import (
	"io"
	"os"
	"path/filepath"
)

// 主要同步函数
func SyncFiles(sourceDir, targetDir string) error {
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(sourceDir, path)
		targetPath := filepath.Join(targetDir, relPath)

		if info.IsDir() {
			// 如果是目录，创建相应的目录
			err := createDirectory(targetPath)
			if err != nil {
				return err
			}
		} else {
			// 如果是文件，检查是否需要更新
			err := syncFile(path, targetPath)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// 创建目录
func createDirectory(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// 同步文件
func syncFile(sourcePath, targetPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	targetFile, err := os.Open(targetPath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if err == nil {
		sourceFileInfo, _ := sourceFile.Stat()
		targetFileInfo, _ := targetFile.Stat()
		if sourceFileInfo.ModTime().After(targetFileInfo.ModTime()) {
			targetFile.Close()
			targetFile, err = os.Create(targetPath)
			if err != nil {
				return err
			}
			defer targetFile.Close()

			_, err = io.Copy(targetFile, sourceFile)
			if err != nil {
				return err
			}
		}
	} else {
		targetFile, err = os.Create(targetPath)
		if err != nil {
			return err
		}
		defer targetFile.Close()

		_, err = io.Copy(targetFile, sourceFile)
		if err != nil {
			return err
		}
	}
	return nil
}

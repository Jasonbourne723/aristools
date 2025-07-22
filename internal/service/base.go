package service

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	DataDir              = "aris_data"
	todoFileName         = "todo.txt"
	wordFileName         = "word.txt"
	errorWordFileName    = "wordError.txt"
	wordAnalysisFileName = "wordAnalysis.txt"
)

func getDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(exePath)
	if err := ensureDir(dir); err != nil {
		return "", err
	}
	return dir, nil
}

func getFilePath(fileName string) (string, error) {

	dir, err := getDir()
	if err != nil {
		return "", err
	}

	p := filepath.Join(dir, DataDir, fileName)
	return p, nil
}

func ensureDir(path string) error {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("创建目录失败: %w", err)
		}
		fmt.Println("目录已创建:", path)
		return nil
	}
	if err != nil {
		return fmt.Errorf("检查目录失败: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("路径已存在但不是目录: %s", path)
	}
	return nil
}

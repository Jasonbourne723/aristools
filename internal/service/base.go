package service

import (
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

func getFilePath(fileName string) (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(exePath)
	p := filepath.Join(dir, DataDir, fileName)
	return p, nil
}

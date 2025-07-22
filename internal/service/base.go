package service

import (
	"os"
	"path/filepath"
)

const (
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
	p := filepath.Join(dir, "aris_data", fileName)
	return p, nil
}

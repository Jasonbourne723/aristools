package cmd

import (
	"aristools/internal/service"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "使数据于远程仓库保持一致",
	Run: func(cmd *cobra.Command, args []string) {
		if err := sync(); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	},
}

func sync() error {

	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	dir := filepath.Dir(exePath)
	dir = filepath.Join(dir, service.DataDir)

	executeCmd(dir, "pull")
	executeCmd(dir, "add", ".")
	executeCmd(dir, "commit", "-m", "\"sync\"")
	executeCmd(dir, "push")

	return nil
}

func executeCmd(dir string, subCmd ...string) {
	cmd := exec.Command("git", subCmd...)
	cmd.Dir = dir

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("output: %v\n", string(output))
}

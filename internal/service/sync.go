package service

import (
	"fmt"
	"os/exec"
)

type syncService struct{}

var SyncSrv = &syncService{}

func (s *syncService) Sync() error {

	dir, err := getDir()
	if err != nil {
		return err
	}

	s.executeCmd(dir, "pull")
	s.executeCmd(dir, "add", ".")
	s.executeCmd(dir, "commit", "-m", "\"sync\"")
	s.executeCmd(dir, "push")

	return nil
}

func (s *syncService) executeCmd(dir string, subCmd ...string) {
	cmd := exec.Command("git", subCmd...)
	cmd.Dir = dir

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("output: %v\n", string(output))
}

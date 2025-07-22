package cmd

import (
	"aristools/internal/service"
	"fmt"

	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "使数据于远程仓库保持一致",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.SyncSrv.Sync(); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	},
}

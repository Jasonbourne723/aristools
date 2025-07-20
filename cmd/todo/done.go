package todo

import (
	"aristools/internal/service"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	id int64
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "完成任务",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.TodoSrv.Done(id); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	},
}

func init() {
	doneCmd.Flags().Int64VarP(&id, "id", "i", 0, "任务Id")
}

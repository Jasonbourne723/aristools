package todo

import (
	"aristools/internal/service"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "完成任务",
	Run: func(cmd *cobra.Command, args []string) {

		ids := make([]int64, 0, len(args))
		for _, arg := range args {
			if id, err := strconv.Atoi(arg); err != nil {
				ids = append(ids, int64(id))
			}
		}
		if err := service.TodoSrv.Done(ids); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	},
}

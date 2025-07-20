package todo

import (
	"aristools/internal/dto"
	"aristools/internal/service"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	SetAddFlag()
}

var (
	name  string
	desc  string
	today bool
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "新增一个todo项",
	Run: func(cmd *cobra.Command, args []string) {

		var addtodo = dto.AddTodoDto{
			Name: name,
		}
		if today {
			addtodo.DoAt = time.Now().Format("2006-01-02")
		}

		service.TodoSrv.Add(addtodo)

	},
}

func SetAddFlag() {
	addCmd.Flags().StringVarP(&name, "name", "n", "unknown", "任务名称")
	addCmd.Flags().StringVarP(&desc, "desc", "d", "", "任务描述")
	addCmd.Flags().BoolVarP(&today, "today", "t", false, "是否是今天的任务")
	addCmd.MarkFlagRequired("name")
}

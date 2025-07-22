package todo

import (
	"aristools/internal/service"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	all bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列举任务",
	Run: func(cmd *cobra.Command, args []string) {
		showList(today, all)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "显示所有任务，包括已完成的任务")
	listCmd.Flags().BoolVarP(&today, "today", "t", false, "显示今天的任务")
}

func showList(today bool, all bool) {
	todos, err := service.TodoSrv.List(today, all)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	// 打印表头，增加竖线和横线
	fmt.Printf(" %-3s %-26s \n", "ID", "Name")

	// 打印数据行，增加竖线
	for _, d := range todos {
		var format = " %-3d   %-26s\n"
		if len(d.DoneAt) != 0 {
			format = strikethrough(format)
		}
		fmt.Printf(format, d.Id, d.Name)
	}
}

func strikethrough(text string) string {
	return "\033[9m" + text + "\033[0m"
}

package todo

import (
	"aristools/internal/service"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	all bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列举任务",
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := service.TodoSrv.List(today, all)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		// 打印表头，增加竖线和横线
		fmt.Printf(" %-3s   %-26s %-3s\n", "ID", "Name", "today")

		// 打印数据行，增加竖线
		for _, d := range todos {
			if len(d.DoneAt) == 0 {
				fmt.Printf(" %-3d    %-26s   %-3t\n", d.Id, d.Name, d.DoAt == time.Now().Format("2006-01-02"))
			} else {
				fmt.Printf(strikethrough(" %-3d    %-26s   %-3t\n"), d.Id, d.Name, d.DoAt == time.Now().Format("2006-01-02"))
			}

		}
	},
}

func init() {
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "显示所有任务，包括已完成的任务")
	listCmd.Flags().BoolVarP(&today, "today", "t", false, "显示今天的任务")
}

func strikethrough(text string) string {
	return "\033[9m" + text + "\033[0m"
}

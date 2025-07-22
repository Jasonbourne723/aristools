package word

import (
	"aristools/internal/service"
	"fmt"

	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "统计单词数量",
	Run: func(cmd *cobra.Command, args []string) {
		count, m, err := service.WordSrv.Count()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Printf("总单词量: %v\n", count)
		}
		fmt.Println("-------------------------------")
		fmt.Printf("%-10s %-10s\n", "正确次数", "数量")
		for key, value := range m {
			fmt.Printf("%-10d %-10d\n", key, value)
		}

	},
}

func init() {
	WordCmd.AddCommand(countCmd)
}

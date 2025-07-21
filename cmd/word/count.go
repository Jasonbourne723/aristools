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
		count, err := service.WordSrv.Count()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Printf("总单词量: %v\n", count)
		}
		errorCount, err := service.ErrorWordSrv.Count()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Printf("错词本单词数量: %v\n", errorCount)
		}
	},
}

func init() {
	WordCmd.AddCommand(countCmd)
}

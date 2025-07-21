package word

import (
	"aristools/internal/service"
	"fmt"

	"github.com/spf13/cobra"
)

var analysisCmd = &cobra.Command{
	Use:   "analysis",
	Short: "分析统计背单词情况",
	Run: func(cmd *cobra.Command, args []string) {

		dtos, err := service.WordAnalysisSrv.GetAll()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		if len(dtos) == 0 {
			fmt.Println("还没有记录，赶快开始吧")
			return
		}
		fmt.Printf(" %-10s  %-5s %-5s\n", "date", "count", "ErrorCount")
		for _, item := range dtos {
			fmt.Printf(" %-10s  %-5d %-5d\n", item.Date, item.Count, item.ErrCount)
		}
	},
}

func init() {
	WordCmd.AddCommand(analysisCmd)
}

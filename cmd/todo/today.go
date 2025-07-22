package todo

import (
	"aristools/internal/service"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "manage today`s todo",
	Long: `manage todo item of today
		eg:
		   display todo item of today : aris todo today 
		   add todo item to today: aris todo today 1 3 `,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			showList(true, false)
		} else {
			var ids []int64
			for _, item := range args {
				if id, err := strconv.ParseInt(item, 10, 64); err != nil {
					fmt.Printf("\"err\": %v\n", "参数错误")
				} else {
					ids = append(ids, id)
				}
			}
			service.TodoSrv.Today(ids)
		}
	},
}

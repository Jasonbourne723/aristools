package todo

import (
	"aristools/internal/service"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "删除todo项,eg: aris todo del 1 2 3",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("请输入Id参数,eg: aris todo today 1")
		}
		var ids []int64
		for _, item := range args {
			if id, err := strconv.ParseInt(item, 10, 64); err != nil {
				fmt.Printf("\"err\": %v\n", "参数错误")
			} else {
				ids = append(ids, id)
			}
		}
		service.TodoSrv.Del(ids)
	},
}

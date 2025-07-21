package word

import (
	"aristools/internal/dto"
	"aristools/internal/service"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "新增单词",
	Long:  "新增单词",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("输入单词eg: home 家,回家\n Ctrl+C 退出\n")
		var en string
		var cn string
		for {
			fmt.Scanf("%s %s\n", &en, &cn)
			if err := service.WordSrv.Add(dto.AddWordDto{
				En: en,
				Cn: strings.Split(cn, ","),
			}); err != nil {
				fmt.Printf("err: %v\n", err)
			} else {
				fmt.Println("Add Successed")
			}
		}

	},
}

func init() {
	WordCmd.AddCommand(addCmd)
}

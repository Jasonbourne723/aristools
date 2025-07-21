package word

import (
	"aristools/internal/service"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	filepath string
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "通过csv文件导入单词",
	Run: func(cmd *cobra.Command, args []string) {
		if count, err := service.WordSrv.Import(filepath); err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Printf("import Successed,count:%d\n", count)
		}
	},
}

func init() {
	importCmd.Flags().StringVarP(&filepath, "filepath", "f", "", "csv文件路径")
	importCmd.MarkFlagRequired("filepath")
	WordCmd.AddCommand(importCmd)
}

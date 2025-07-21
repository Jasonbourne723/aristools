package word

import (
	"fmt"

	"github.com/spf13/cobra"
)

var WordCmd = &cobra.Command{
	Use:   "word",
	Short: "背单词",
	Long:  "背单词工具",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("背单词")
	},
}

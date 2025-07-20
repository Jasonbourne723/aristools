package todo

import (
	"fmt"

	"github.com/spf13/cobra"
)

var TodoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo is a task manager",
	Long:  "Todo is a task manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run aris...")
	},
}

func init() {
	TodoCmd.AddCommand(addCmd)
	TodoCmd.AddCommand(listCmd)
	TodoCmd.AddCommand(doneCmd)
}

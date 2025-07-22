package cmd

import (
	"aristools/cmd/todo"
	"aristools/cmd/word"
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aris",
	Short: "一个大而全的命令行工具",
	Long:  `一个大而全的命令行工具`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run aris...")
	},
}

func Execute() {
	rootCmd.AddCommand(todo.TodoCmd)
	rootCmd.AddCommand(word.WordCmd)
	rootCmd.AddCommand(syncCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

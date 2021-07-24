package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task to your tasks list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
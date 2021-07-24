package cmd

import (
	"fmt"
	"os"
	"task/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			println("Something went wrong:", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			println("You have no tasks complete!")
		}else {
			fmt.Println("You have the following tasks: ")
			for _, task := range tasks {
				fmt.Printf("%d. %s\n",task.Key, task.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

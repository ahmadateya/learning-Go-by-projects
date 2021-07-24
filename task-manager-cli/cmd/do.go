package cmd

import (
	"strconv"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks task as completed.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				println("Failed to package the argument:", arg)
			}else {
				ids = append(ids, id)
			}
		}
		println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}

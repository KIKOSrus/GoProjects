package cmd

import (
	"TaskManager/db"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task",
	Run: func(cmd *cobra.Command, args []string) {

		task := strings.Join(args, " ")

		err := db.CreateTask(task)

		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {

	rootCmd.AddCommand(addCmd)
}

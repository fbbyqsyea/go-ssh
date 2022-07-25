package cmd

import (
	"github.com/fbbyqsyea/go-ssh/service"
	"github.com/spf13/cobra"
)

var name string

var password bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List ssh connect",
	Run: func(cmd *cobra.Command, args []string) {
		if name != "" {
			service.ListName(name, password)
		} else {
			service.List(password)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the ssh connect")
	listCmd.Flags().BoolVarP(&password, "password", "p", false, "Is show password")
}

package cmd

import (
	"os"

	"github.com/fbbyqsyea/go-ssh/service"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-ssh",
	Short: "A kit for manage shell connect",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// pre check
	service.PreCheck()
	// init sqlite3 connect and database
	service.InitSqlite3ConnectAndDatabase()
	defer service.CloseDB()
	// run command
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

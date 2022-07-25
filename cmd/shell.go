/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fbbyqsyea/go-ssh/service"
	"github.com/fbbyqsyea/go-ssh/utils"
	"github.com/spf13/cobra"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Login a ssh connect",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			utils.Error(fmt.Errorf("ssh name cannot be empty"))
			os.Exit(1)
		}
		name := args[0]
		// check name is exists in table
		if !service.NameExists(name) {
			utils.Error(fmt.Errorf("name %s is not exists", name))
			os.Exit(1)
		}
		// get ssh connect
		sc := service.GetConnectByName(name)
		service.Shell(sc.Host, sc.User, sc.Password, sc.Port)
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
}

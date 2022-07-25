package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fbbyqsyea/go-ssh/service"
	"github.com/fbbyqsyea/go-ssh/utils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a ssh connect",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		// ssh connect name
		utils.Info("Enter ssh connect name (inside 32 characters):")
		cmdString, err := reader.ReadString('\n')
		utils.CheckIfError(err)
		name := strings.TrimRight(cmdString, "\n")
		// name cannot be empty
		if name == "" {
			utils.Error(fmt.Errorf("ssh connect name cannot be empty"))
			os.Exit(1)
		}
		// name length must lt 32 characters
		if len(name) > 32 {
			utils.Error(fmt.Errorf("ssh name must lt 32 characters"))
		}

		// ssh connect host
		utils.Info("Enter ssh connect host (inside 32 characters):")
		cmdString, err = reader.ReadString('\n')
		utils.CheckIfError(err)
		host := strings.TrimRight(cmdString, "\n")
		// host cannot be empty
		if host == "" {
			utils.Error(fmt.Errorf("ssh connect host cannot be empty"))
			os.Exit(1)
		}
		// host length must lt 32 characters
		if len(host) > 32 {
			utils.Error(fmt.Errorf("ssh host must lt 32 characters"))
		}

		// ssh connect port
		utils.Info("Enter ssh connect port (in (0, 65535) default:22):")
		cmdString, err = reader.ReadString('\n')
		utils.CheckIfError(err)
		portStr := strings.TrimRight(cmdString, "\n")
		port := 0
		if portStr != "" {
			port, err = strconv.Atoi(portStr)
			utils.CheckIfError(err)
		}
		// default port is 22
		if port == 0 {
			port = 22
		}
		// port must lt 65535
		if port < 0 || port > 65535 {
			utils.Error(fmt.Errorf("ssh port must in (0, 65535)"))
		}

		// ssh connect user
		utils.Info("Enter ssh connect user (inside 32 characters):")
		cmdString, err = reader.ReadString('\n')
		utils.CheckIfError(err)
		user := strings.TrimRight(cmdString, "\n")
		// user cannot be empty
		if user == "" {
			utils.Error(fmt.Errorf("ssh connect user cannot be empty"))
			os.Exit(1)
		}
		// user length must lt 32 characters
		if len(user) > 32 {
			utils.Error(fmt.Errorf("ssh user must lt 32 characters"))
		}

		// ssh connect password
		utils.Info("Enter ssh connect password (inside 255 characters):")
		cmdString, err = reader.ReadString('\n')
		utils.CheckIfError(err)
		password := strings.TrimRight(cmdString, "\n")
		// passwd length must lt 255 characters
		if len(password) > 255 {
			utils.Error(fmt.Errorf("ssh host must lt 255 characters"))
		}

		service.Create(name, host, user, password, port)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

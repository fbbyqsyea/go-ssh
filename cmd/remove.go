package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fbbyqsyea/go-ssh/service"
	"github.com/fbbyqsyea/go-ssh/utils"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a ssh connect",
	Run: func(cmd *cobra.Command, args []string) {

		reader := bufio.NewReader(os.Stdin)
		if name == "" {
			// ssh connect name
			utils.Info("Enter ssh connect name:")
			cmdString, err := reader.ReadString('\n')
			utils.CheckIfError(err)
			name = strings.TrimRight(cmdString, "\n")
		}
		if !service.NameExists(name) {
			utils.Error(fmt.Errorf("name %s is not exists", name))
			os.Exit(1)
		}

		// list ssh connect info
		utils.Info("%s ssh connect info:", name)
		service.ListName(name, false)

		// confirm remove
		utils.Info("Are you confirm remove this ssh connect? please enter (yes) to confirm:")
		cmdString, err := reader.ReadString('\n')
		utils.CheckIfError(err)
		confirm := strings.TrimRight(cmdString, "\n")
		if confirm == "yes" {
			service.Remove(name)
		} else {
			utils.Warning("cancle remove command")
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the ssh connect")
}

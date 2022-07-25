package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ExecCommand(cmdStr string) {
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func ExecCommandWithResult(cmdStr string) string {
	out, err := exec.Command("bash", "-c", cmdStr).CombinedOutput()
	if err != nil && !strings.Contains(err.Error(), "exit status") {
		fmt.Fprintf(os.Stderr, "err: [%v].\n", err)
		return ""
	}
	return strings.Trim(string(out), "\n")
}

package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// CheckArgs should be used to ensure the right command line arguments are
// passed before executing an example.
func CheckArgs(arg ...string) {
	if len(os.Args) < len(arg)+1 {
		Warning("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Fprintf(os.Stderr, "\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// debug
func Debug(args ...interface{}) {
	spew.Dump(args...)
	os.Exit(2)
}

// error
func Error(err error) {
	fmt.Fprintf(os.Stderr, "\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
}

// info
func Info(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, "\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// warning
func Warning(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// error
func ErrorithResult(err error) string {
	return fmt.Sprintf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
}

// info
func InfoWithResult(format string, args ...interface{}) string {
	return fmt.Sprintf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// warning
func WarningithResult(format string, args ...interface{}) string {
	return fmt.Sprintf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

package service

import (
	"fmt"
	"os"
	"time"

	"github.com/fbbyqsyea/go-ssh/utils"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func Shell(host, user, password string, port int) {
	// ssh
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	})
	utils.CheckIfError(err)

	// session
	session, err := client.NewSession()
	utils.CheckIfError(err)
	defer session.Close()

	// store fd state
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	utils.CheckIfError(err)
	defer term.Restore(fd, oldState)

	// bind stdout stderr stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	termWidth, termHeight, err := term.GetSize(fd)
	utils.CheckIfError(err)

	// session modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	// request a pty
	err = session.RequestPty("xterm-256color", termHeight, termWidth, modes)
	utils.CheckIfError(err)

	// start connect ssh server
	err = session.Shell()
	utils.CheckIfError(err)

	// wait ssh session exit
	session.Wait()

}

// +build windows

package helper

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
)

func RunCommand(command string) bool {
	c := exec.Command("cmd") // dummy executable that actually needs to exist but we'll overwrite using .SysProcAttr

	// Based on the official docs, syscall.SysProcAttr.CmdLine doesn't exist.
	// But it does and is vital:
	// https://github.com/golang/go/issues/15566#issuecomment-333274825
	// https://medium.com/@felixge/killing-a-child-process-and-all-of-its-children-in-go-54079af94773
	c.SysProcAttr = &syscall.SysProcAttr{CmdLine: command}

	var stderr bytes.Buffer
	c.Stderr = &stderr

	err := c.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return false
	}
	return true
}

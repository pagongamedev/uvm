package helper

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func GetFolderVersion(sdkPath string, sVersion string, sTag string) (string, string) {
	sFolderVersion := GetVersionWithV(sVersion)

	if sTag != "" {
		sFolderVersion += "-" + sTag
	}

	sSDKPathVersion := filepath.Join(sdkPath, sFolderVersion)

	return sFolderVersion, sSDKPathVersion
}

func GetVersionTagFromPath(baseFile string) (string, string) {
	sVersion := ""
	sTag := ""
	strList := strings.SplitN(baseFile, "/", 2)

	if len(strList) > 0 {
		sVersion = strList[0]
	}
	if len(strList) > 1 {
		sTag = strList[1]
	}

	return sVersion, sTag
}

func GetVersionWithV(sVersion string) string {
	if []rune(sVersion)[0] != 'v' {
		sVersion = "v" + sVersion
	}
	return sVersion
}

func GetVersionWithOutV(sVersion string) string {
	if []rune(sVersion)[0] == 'v' {
		sVersion = strings.Replace(sVersion, "v", "", 1)
	}
	return sVersion
}

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

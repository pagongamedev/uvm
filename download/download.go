package download

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pagongamedev/uvm/sdk"
)

var client = &http.Client{}

//Test func
func Loading(sd sdk.SDK, rootPath string, sdkPath string, sUrl string, sTempFile string, sFileName string, sVersion string, sTag, sFolderVersion string, sSDKPathVersion string) error {
	err := checkExistUrl(sUrl)
	if err != nil {
		return errors.New("check url version not exist")
	}
	fmt.Println("installing :", sd.GetName(), sVersion, sTag)
	fmt.Println()

	Downloading(sUrl, sTempFile)

	return nil
}

func Downloading(sUrl string, sTempFile string) bool {
	fmt.Println("Downloading...")
	output, err := os.Create(sTempFile)
	if err != nil {
		fmt.Println("Error while creating", sTempFile, "-", err)
		return false
	}
	defer output.Close()

	response, err := client.Get(sUrl)
	if err != nil {
		fmt.Println("Error while downloading", sUrl, "-", err)
		return false
	}
	defer response.Body.Close()

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Download interrupted.Rolling back...")
		output.Close()
		response.Body.Close()
		err = os.Remove(sTempFile)
		if err != nil {
			fmt.Println("Error while rolling back", err)
		}
		os.Exit(1)
	}()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", sUrl, "-", err)
	}
	if response.Status[0:3] != "200" {
		fmt.Println("Download failed. Rolling Back.")
		err := os.Remove(sTempFile)
		if err != nil {
			fmt.Println("Rollback failed.", err)
		}
		return false
	}

	return true
}

func checkExistUrl(url string) error {
	response, err := client.Head(url)
	if err != nil {
		return err
	}
	if response.Status[0:3] != "200" {
		return errors.New("404 Not Found")
	}

	return nil
}

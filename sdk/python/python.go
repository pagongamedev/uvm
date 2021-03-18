package python

import (
	"github.com/pagongamedev/uvm/sdk"
)

func NewSDK(sPlatform string) (sdk.SDK, error) {
	var provider sdk.Provider
	switch sPlatform {
	case "windows":
		provider, _ = NewProviderPythonOrg(sPlatform)
	case "darwin":
		provider, _ = NewProviderPythonOrg(sPlatform)
	case "linux":
		provider, _ = NewProviderPythonOrg(sPlatform)
	default:
		provider, _ = NewProviderPythonOrg(sPlatform)
	}

	// ==================================

	return sdk.SDK{
		Name:       "Python",
		LinkName:   "Python",
		Command:    "-p",
		Env:        "",
		EnvBin:     "",
		EnvChannel: "",
		Provider:   provider,
	}, nil
}

func NewProviderPythonOrg(sPlatform string) (sdk.Provider, error) {
	fileType := ""
	archiveType := ""

	switch sPlatform {
	case "windows":
		fileType = "zip"
		archiveType = "zip"
	case "darwin":
		fileType = "zip"
		archiveType = "zip"
	case "linux":
		fileType = "zip"
		archiveType = "zip"
	default:
		fileType = "zip"
		archiveType = "zip"
	}

	// ==================================
	mapOSList := map[string]string{}
	// windows
	// darwin
	// linux
	mapArchList := map[string]string{}
	// 386
	// amd64
	// arm
	// arm64

	mapTagList := map[string]string{}

	mapTagFolderList := map[string]string{}
	// ==================================

	return sdk.Provider{
		IsManualInstall:  true,
		IsCreateFolder:   false,
		IsRenameFolder:   true,
		Header:           "Provider Python.org",
		LinkPage:         "https://www.python.org/downloads/",
		LinkDist:         "https://www.python.org/ftp/python/",
		Path:             "{{version}}/{{fileName}}.{{type}}",
		FileName:         "Python-{{version}}",
		ZipFolderName:    "Python-{{version}}.{{type}}",
		FileType:         fileType,
		ArchiveType:      archiveType,
		MapOSList:        mapOSList,
		MapArchList:      mapArchList,
		MapTagList:       mapTagList,
		MapTagFolderList: mapTagFolderList,
	}, nil
}

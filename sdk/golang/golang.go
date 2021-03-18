package golang

import (
	"github.com/pagongamedev/uvm/sdk"
)

func NewSDK(sPlatform string) (sdk.SDK, error) {
	var provider sdk.Provider
	switch sPlatform {
	case "windows":
		provider, _ = NewProviderGolangOrg(sPlatform)
	case "darwin":
		provider, _ = NewProviderGolangOrg(sPlatform)
	case "linux":
		provider, _ = NewProviderGolangOrg(sPlatform)
	default:
		provider, _ = NewProviderGolangOrg(sPlatform)
	}

	// ==================================

	return sdk.SDK{
		Name:       "Golang",
		LinkName:   "Golang",
		Command:    "-g",
		Env:        "",
		EnvBin:     "\\bin",
		EnvChannel: "",
		Provider:   provider,
	}, nil
}

func NewProviderGolangOrg(sPlatform string) (sdk.Provider, error) {
	fileType := ""
	archiveType := ""

	switch sPlatform {
	case "windows":
		fileType = "zip"
		archiveType = "zip"
	case "darwin":
		fileType = "tar.gz"
		archiveType = "tar"
	case "linux":
		fileType = "tar.gz"
		archiveType = "tar"
	default:
		fileType = "tar.gz"
		archiveType = "tar"
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
	mapTagList[""] = "stable"

	mapTagFolderList := map[string]string{}
	// ==================================

	return sdk.Provider{
		IsManualInstall:  false,
		IsCreateFolder:   false,
		IsRenameFolder:   true,
		Header:           "Provider Golang.org",
		LinkPage:         "https://golang.org/dl/",
		LinkDist:         "https://golang.org/dl/",
		Path:             "{{fileName}}.{{type}}",
		FileName:         "go{{version}}.{{os}}-{{arch}}",
		ZipFolderName:    "go",
		FileType:         fileType,
		ArchiveType:      archiveType,
		MapOSList:        mapOSList,
		MapArchList:      mapArchList,
		MapTagList:       mapTagList,
		MapTagFolderList: mapTagFolderList,
	}, nil
}

package java

import (
	"github.com/pagongamedev/uvm/sdk"
)

func NewSDK(sPlatform string) (*sdk.SDK, error) {
	var provider sdk.Provider
	switch sPlatform {
	case "windows":
		provider, _ = NewProviderOracle(sPlatform)
	case "darwin":
		provider, _ = NewProviderOracle(sPlatform)
	case "linux":
		provider, _ = NewProviderOracle(sPlatform)
	default:
		provider, _ = NewProviderOracle(sPlatform)
	}

	// ==================================

	return &sdk.SDK{
		Name:       "Java",
		LinkName:   "Java",
		Command:    "-oj",
		Env:        "JAVA_HOME",
		EnvBin:     "\\bin",
		EnvChannel: "UVM_JAVA_CHANNEL",
		Provider:   provider,
	}, nil
}

func NewProviderOracle(sPlatform string) (sdk.Provider, error) {
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

	mapTagFolderList := map[string]string{}
	// ==================================

	return sdk.Provider{
		IsManualInstall:  true,
		IsCreateFolder:   false,
		IsRenameFolder:   true,
		Header:           "Provider Oracle",
		LinkPage:         "https://www.oracle.com/java/technologies/javase-jdk16-downloads.html",
		LinkDist:         "https://download.oracle.com/otn-pub/java/jdk/",
		Path:             "{{key}}/{{fileName}}.{{type}}",
		FileName:         "jdk-{{version}}_{{os}}-{{arch}}_bin",
		ZipFolderName:    "jdk-{{version}}",
		FileType:         fileType,
		ArchiveType:      archiveType,
		MapOSList:        mapOSList,
		MapArchList:      mapArchList,
		MapTagList:       mapTagList,
		MapTagFolderList: mapTagFolderList,
	}, nil
}

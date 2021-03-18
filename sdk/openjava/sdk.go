package openjava

import (
	"github.com/pagongamedev/uvm/sdk"
)

func NewSDK(sPlatform string) (*sdk.SDK, error) {
	var provider sdk.Provider
	switch sPlatform {
	case "windows":
		provider, _ = NewProviderJavaNet(sPlatform)
	case "darwin":
		provider, _ = NewProviderJavaNet(sPlatform)
	case "linux":
		provider, _ = NewProviderJavaNet(sPlatform)
	default:
		provider, _ = NewProviderJavaNet(sPlatform)
	}

	// ==================================

	return &sdk.SDK{
		Name:       "OpenJava",
		LinkName:   "Java",
		Command:    "-oj",
		Env:        "JAVA_HOME",
		EnvBin:     "\\bin",
		EnvChannel: "UVM_JAVA_CHANNEL",
		Provider:   provider,
	}, nil
}

func NewProviderJavaNet(sPlatform string) (sdk.Provider, error) {
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
	mapOSList["darwin"] = "osx"
	// linux
	mapArchList := map[string]string{}
	// 386
	mapArchList["amd64"] = "x64"
	// arm
	// arm64

	mapTagList := map[string]string{}

	mapTagFolderList := map[string]string{}

	// ==================================

	return sdk.Provider{
		IsManualInstall:  false,
		IsCreateFolder:   false,
		IsRenameFolder:   true,
		Header:           "Provider Java.net",
		LinkPage:         "https://jdk.java.net/archive/",
		LinkDist:         "https://download.java.net/java/GA/jdk",
		Path:             "{{version}}/{{key}}/GPL/{{fileName}}.{{type}}",
		FileName:         "openjdk-{{version}}_{{os}}-{{arch}}_bin",
		ZipFolderName:    "jdk-{{version}}",
		FileType:         fileType,
		ArchiveType:      archiveType,
		MapOSList:        mapOSList,
		MapArchList:      mapArchList,
		MapTagList:       mapTagList,
		MapTagFolderList: mapTagFolderList,
	}, nil
}

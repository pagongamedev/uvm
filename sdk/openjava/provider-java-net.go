package openjava

import "github.com/pagongamedev/uvm/sdk"

func NewProviderJavaNet(sPlatform string) (sdk.Provider, error) {
	fileType := ""
	archiveType := ""

	switch sPlatform {
	case "windows":
		fileType = "zip"
		archiveType = "zip"
	case "darwin":
		fileType = "tar.gz"
		archiveType = "targz"
	case "linux":
		fileType = "tar.gz"
		archiveType = "targz"
	default:
		fileType = "tar.gz"
		archiveType = "targz"
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
		IsUseKey:         true,
		DetailKey:        "https://download.java.net/java/GA/jdk0.0.0/{{key}}/GPL/openjdk-0.0.0_linux-x64_bin.tar.gz",
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

package nodejs

import "github.com/pagongamedev/uvm/sdk"

func NewProviderNodejsOrg(sPlatform string) (sdk.Provider, error) {
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
	mapOSList["windows"] = "win"
	// linux
	// darwin

	mapArchList := map[string]string{}
	mapArchList["386"] = "x32"
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
		Header:           "Provider Nodejs.org",
		LinkPage:         "https://nodejs.org/dist/",
		LinkDist:         "https://nodejs.org/dist/",
		Path:             "v{{version}}/{{fileName}}.{{type}}",
		FileName:         "node-v{{version}}-{{os}}-{{arch}}",
		ZipFolderName:    "node-v{{version}}-{{os}}-{{arch}}",
		FileType:         fileType,
		ArchiveType:      archiveType,
		MapOSList:        mapOSList,
		MapArchList:      mapArchList,
		MapTagList:       mapTagList,
		MapTagFolderList: mapTagFolderList,
	}, nil
}

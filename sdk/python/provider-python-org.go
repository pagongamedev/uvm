package python

import "github.com/pagongamedev/uvm/sdk"

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

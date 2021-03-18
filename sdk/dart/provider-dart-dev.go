package dart

import "github.com/pagongamedev/uvm/sdk"

func NewProviderDartDev(sPlatform string) (sdk.Provider, error) {
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
	mapOSList["darwin"] = "macos"
	// linux
	mapArchList := map[string]string{}
	// 386
	mapArchList["amd64"] = "x64"
	// arm
	// arm64

	mapTagList := map[string]string{}
	mapTagList["beta"] = ".beta"
	mapTagList["dev"] = ".dev"

	mapTagFolderList := map[string]string{}
	mapTagFolderList[""] = "stable"

	return sdk.Provider{
		IsManualInstall:  false,
		IsCreateFolder:   false,
		IsRenameFolder:   true,
		Header:           "Provider Dart.dev",
		LinkPage:         "https://dart.dev/tools/sdk/archive",
		LinkDist:         "https://storage.googleapis.com/dart-archive/channels/",
		Path:             "{{tagFolder}}/release/{{version}}{{tag}}/sdk/{{fileName}}.{{type}}",
		FileName:         "dartsdk-{{os}}-{{arch}}-release",
		ZipFolderName:    "dart-sdk",
		FileType:         fileType,
		ArchiveType:      archiveType,
		MapOSList:        mapOSList,
		MapArchList:      mapArchList,
		MapTagList:       mapTagList,
		MapTagFolderList: mapTagFolderList,
	}, nil
}

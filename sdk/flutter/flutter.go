package flutter

import (
	"github.com/pagongamedev/uvm/sdk"
)

func NewSDK(sPlatform string) (sdk.SDK, error) {
	var provider sdk.Provider
	switch sPlatform {
	case "windows":
		provider, _ = NewProviderFlutterDev(sPlatform)
	case "darwin":
		provider, _ = NewProviderFlutterDev(sPlatform)
	case "linux":
		provider, _ = NewProviderFlutterDev(sPlatform)
	default:
		provider, _ = NewProviderFlutterDev(sPlatform)
	}

	// ==================================

	return sdk.SDK{
		Name:       "Flutter",
		LinkName:   "Flutter",
		Command:    "-f",
		Env:        "",
		EnvBin:     "\\bin",
		EnvChannel: "",
		Provider:   provider,
	}, nil
}

func NewProviderFlutterDev(sPlatform string) (sdk.Provider, error) {
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
		fileType = "tar.xz"
		archiveType = "tar"
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
	// amd64
	// arm
	// arm64

	mapTagList := map[string]string{}
	mapTagList[""] = "stable"

	mapTagFolderList := map[string]string{}
	mapTagFolderList[""] = "stable"
	// ==================================

	return sdk.Provider{
		IsManualInstall:  false,
		IsCreateFolder:   false,
		IsRenameFolder:   true,
		Header:           "Provider Flutter.dev",
		LinkPage:         "https://flutter.dev/docs/development/tools/sdk/releases",
		LinkDist:         "https://storage.googleapis.com/flutter_infra/releases/",
		Path:             "{{tagFolder}}/{{os}}/{{fileName}}.{{type}}",
		FileName:         "flutter_{{os}}_{{version}}-{{tag}}",
		ZipFolderName:    "flutter",
		FileType:         fileType,
		ArchiveType:      archiveType,
		MapOSList:        mapOSList,
		MapArchList:      mapArchList,
		MapTagList:       mapTagList,
		MapTagFolderList: mapTagFolderList,
	}, nil
}

package ruby

import (
	"github.com/pagongamedev/uvm/sdk"
)

func NewSDK(sPlatform string) (sdk.SDK, error) {
	var provider sdk.Provider
	switch sPlatform {
	case "windows":
		provider, _ = NewProviderRubyinstallerOrg(sPlatform)
	// case "darwin":
	// case "linux":
	default:
		provider, _ = NewProviderRubyinstallerOrg(sPlatform)
	}

	// ==================================

	return sdk.SDK{
		Name:       "Ruby",
		LinkName:   "Ruby",
		Command:    "-r",
		Env:        "",
		EnvBin:     "\\bin",
		EnvChannel: "",
		Provider:   provider,
	}, nil
}

func NewProviderRubyinstallerOrg(sPlatform string) (sdk.Provider, error) {
	fileType := ""
	archiveType := ""

	switch sPlatform {
	case "windows":
		fileType = "7z"
		archiveType = "7z"
	case "darwin":
		fileType = "7z"
		archiveType = "7z"
	case "linux":
		fileType = "7z"
		archiveType = "7z"
	default:
		fileType = "7z"
		archiveType = "7z"
	}

	// ==================================
	mapOSList := map[string]string{}
	// windows
	// darwin
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
		IsManualInstall:  true,
		IsCreateFolder:   false,
		IsRenameFolder:   true,
		Header:           "Provider Rubyinstaller.org\nhttps://www.ruby-lang.org/en/downloads",
		LinkPage:         "https://rubyinstaller.org/",
		LinkDist:         "https://github.com/oneclick/rubyinstaller2/releases/download/",
		Path:             "RubyInstaller-{{version}}/{{fileName}}.{{type}}",
		FileName:         "rubyinstaller-{{version}}-{{arch}}",
		ZipFolderName:    "rubyinstaller-{{version}}-{{arch}}",
		FileType:         fileType,
		ArchiveType:      archiveType,
		MapOSList:        mapOSList,
		MapArchList:      mapArchList,
		MapTagList:       mapTagList,
		MapTagFolderList: mapTagFolderList,
	}, nil
}

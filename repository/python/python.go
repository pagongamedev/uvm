package python

import (
	"github.com/pagongamedev/uvm/repository"
)

func NewRepository(sPlatform string) (repository.Repository, error) {
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

	// ==================================

	r := repo{
		isManualInstall:  true,
		name:             "Python",
		linkName:         "Python",
		command:          "-p",
		env:              "",
		envBin:           "",
		envChannel:       "",
		linkPage:         "https://www.python.org/downloads/",
		dist:             "https://www.python.org/ftp/python/",
		path:             "{{version}}/{{fileName}}.{{type}}",
		fileName:         "Python-{{version}}",
		zipFolderName:    "Python-{{version}}.{{type}}",
		fileType:         fileType,
		archiveType:      archiveType,
		mapOSList:        mapOSList,
		mapArchList:      mapArchList,
		mapTagList:       mapTagList,
		mapTagFolderList: mapTagFolderList,
		isCreateFolder:   false,
		isRenameFolder:   true,
	}

	return &r, nil
}

type repo struct {
	isManualInstall  bool
	name             string
	linkName         string
	command          string
	env              string
	envChannel       string
	envBin           string
	linkPage         string
	dist             string
	path             string
	fileName         string
	zipFolderName    string
	fileType         string
	archiveType      string
	mapOSList        map[string]string
	mapArchList      map[string]string
	mapTagList       map[string]string
	mapTagFolderList map[string]string
	isCreateFolder   bool
	isRenameFolder   bool
}

func (r *repo) GetDist() string {
	return r.dist
}

func (r *repo) GetName() string {
	return r.name
}

func (r *repo) GetLinkName() string {
	return r.linkName
}

func (r *repo) GetCommand() string {
	return r.command
}

func (r *repo) GetEnv() string {
	return r.env
}

func (r *repo) GetEnvChannel() string {
	return r.envChannel
}

func (r *repo) GetEnvBin() string {
	return r.envBin
}

func (r *repo) GetLinkPage() string {
	return r.linkPage
}

func (r *repo) GetFileName() string {
	return r.fileName
}

func (r *repo) GetZipFolderName() string {
	return r.zipFolderName
}

func (r *repo) GetFileType() string {
	return r.fileType
}

func (r *repo) GetArchiveType() string {
	return r.archiveType
}

func (r *repo) GetPath() string {
	return r.path
}

func (r *repo) GetMapOSList(key string) string {
	val := r.mapOSList[key]
	if val == "" {
		return key
	}
	return r.mapOSList[key]
}

func (r *repo) GetMapArchList(key string) string {
	val := r.mapArchList[key]
	if val == "" {
		return key
	}
	return r.mapArchList[key]
}

func (r *repo) GetMapTagList(key string) string {
	val := r.mapTagList[key]
	if val == "" {
		return key
	}
	return r.mapTagList[key]
}

func (r *repo) GetMapTagFolderList(key string) string {
	val := r.mapTagFolderList[key]
	if val == "" {
		return key
	}
	return r.mapTagFolderList[key]
}

func (r *repo) GetIsCreateFolder() bool {
	return r.isCreateFolder
}

func (r *repo) GetIsRenameFolder() bool {
	return r.isRenameFolder
}

func (r *repo) GetIsManualInstall() bool {
	return r.isManualInstall
}

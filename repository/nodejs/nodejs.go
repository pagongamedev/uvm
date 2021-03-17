package nodejs

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
		fileType = "tar.xz"
		archiveType = "tar"
	case "linux":
		fileType = "tar.xz"
		archiveType = "tar"
	default:
		fileType = "tar.xz"
		archiveType = "tar"
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
	// ==================================

	r := repo{
		name:          "NodeJS",
		command:       "-n",
		env:           "UVM_NODEJS_HOME",
		envBin:        "",
		dist:          "https://nodejs.org/dist/",
		path:          "v{{version}}/{{fileName}}.{{type}}",
		fileName:      "node-v{{version}}-{{os}}-{{arch}}",
		zipFolderName: "node-v{{version}}-{{os}}-{{arch}}",
		fileType:      fileType,
		archiveType:   archiveType,
		mapOSList:     mapOSList,
		mapArchList:   mapArchList,
		mapTagList:    mapTagList, isCreateFolder: false,
		isRenameFolder: true,
	}

	return &r, nil
}

type repo struct {
	name           string
	command        string
	env            string
	envBin         string
	dist           string
	path           string
	fileName       string
	zipFolderName  string
	fileType       string
	archiveType    string
	mapOSList      map[string]string
	mapArchList    map[string]string
	mapTagList     map[string]string
	isCreateFolder bool
	isRenameFolder bool
}

func (r *repo) GetDist() string {
	return r.dist
}

func (r *repo) GetName() string {
	return r.name
}

func (r *repo) GetCommand() string {
	return r.command
}

func (r *repo) GetEnv() string {
	return r.env
}

func (r *repo) GetEnvBin() string {
	return r.envBin
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

func (r *repo) GetIsCreateFolder() bool {
	return r.isCreateFolder
}

func (r *repo) GetIsRenameFolder() bool {
	return r.isRenameFolder
}

// https://nodejs.org/dist/v14.16.0/node-v14.16.0-win-x64.zip
// https://nodejs.org/dist/v14.16.0/node-v14.16.0-win-x64.7z
// node-v14.16.0-darwin-x64.tar.gz                    23-Feb-2021 00:29            31567754
// node-v14.16.0-darwin-x64.tar.xz
// https://nodejs.org/dist/latest-v14.x/

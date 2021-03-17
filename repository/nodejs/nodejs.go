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
	default:
		fileType = "tar.gz"
		archiveType = "tar"
	}

	// ==================================
	mapList := map[string]string{}
	mapList["windows"] = "win"
	// linux
	// darwin

	mapList["386"] = "x32"
	mapList["amd64"] = "x64"
	// arm
	// arm64
	// ==================================

	r := repo{
		name:           "NodeJS",
		command:        "-n",
		env:            "UVM_NODEJS_HOME",
		envBin:         "\\bin",
		dist:           "https://nodejs.org/dist/",
		fileName:       "node-{{version}}-{{os}}-{{arch}}",
		fileType:       fileType,
		path:           "{{version}}/{{fileName}}.{{type}}",
		archiveType:    archiveType,
		mapList:        mapList,
		isCreateFolder: false,
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
	fileName       string
	fileType       string
	path           string
	archiveType    string
	mapList        map[string]string
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

func (r *repo) GetFileType() string {
	return r.fileType
}

func (r *repo) GetArchiveType() string {
	return r.archiveType
}

func (r *repo) GetPath() string {
	return r.path
}

func (r *repo) GetMapList(key string) string {
	return r.mapList[key]
}

func (r *repo) GetIsCreateFolder() bool {
	return r.isCreateFolder
}

func (r *repo) GetIsRenameFolder() bool {
	return r.isRenameFolder
}

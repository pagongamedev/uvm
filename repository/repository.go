package repository

// Repository interface
type Repository interface {
	GetDist() string
	GetName() string
	GetLinkName() string
	GetCommand() string
	GetEnv() string
	GetEnvChannel() string
	GetEnvBin() string
	GetLinkPage() string
	GetFileName() string
	GetZipFolderName() string
	GetFileType() string
	GetArchiveType() string
	GetPath() string
	GetMapOSList(key string) string
	GetMapArchList(key string) string
	GetMapTagList(key string) string
	GetMapTagFolderList(key string) string
	GetIsCreateFolder() bool
	GetIsRenameFolder() bool
	GetIsManualInstall() bool
}

package repository

// Repository interface
type Repository interface {
	GetDist() string
	GetName() string
	GetCommand() string
	GetEnv() string
	GetEnvBin() string
	GetFileName() string
	GetZipFolderName() string
	GetFileType() string
	GetArchiveType() string
	GetPath() string
	GetMapOSList(key string) string
	GetMapArchList(key string) string
	GetMapTagList(key string) string
	GetIsCreateFolder() bool
	GetIsRenameFolder() bool
}

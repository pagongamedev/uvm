package repository

// Repository interface
type Repository interface {
	GetDist() string
	GetName() string
	GetCommand() string
	GetEnv() string
	GetEnvBin() string
	GetFileName() string
	GetFileType() string
	GetPath() string
	GetMapList(key string) string
	GetIsCreateFolder() bool
	GetIsRenameFolder() bool
}

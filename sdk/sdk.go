package sdk

type SDK struct {
	Name       string
	LinkName   string
	Command    string
	Env        string
	EnvBin     string
	EnvChannel string
	Provider   Provider
}

func (sdk *SDK) GetHeader() string {
	return sdk.Provider.Header
}

func (sdk *SDK) GetDist() string {
	return sdk.Provider.LinkDist
}

func (sdk *SDK) GetName() string {
	return sdk.Name
}

func (sdk *SDK) GetLinkName() string {
	return sdk.LinkName
}

func (sdk *SDK) GetCommand() string {
	return sdk.Command
}

func (sdk *SDK) GetEnv() string {
	return sdk.Env
}

func (sdk *SDK) GetEnvChannel() string {
	return sdk.EnvChannel
}

func (sdk *SDK) GetEnvBin() string {
	return sdk.EnvBin
}

func (sdk *SDK) GetLinkPage() string {
	return sdk.Provider.LinkPage
}

func (sdk *SDK) GetFileName() string {
	return sdk.Provider.FileName
}

func (sdk *SDK) GetZipFolderName() string {
	return sdk.Provider.ZipFolderName
}

func (sdk *SDK) GetFileType() string {
	return sdk.Provider.FileType
}

func (sdk *SDK) GetArchiveType() string {
	return sdk.Provider.ArchiveType
}

func (sdk *SDK) GetPath() string {
	return sdk.Provider.Path
}

func (sdk *SDK) GetMapOSList(key string) string {
	val := sdk.Provider.MapOSList[key]
	if val == "" {
		return key
	}
	return sdk.Provider.MapOSList[key]
}

func (sdk *SDK) GetMapArchList(key string) string {
	val := sdk.Provider.MapArchList[key]
	if val == "" {
		return key
	}
	return sdk.Provider.MapArchList[key]
}

func (sdk *SDK) GetMapTagList(key string) string {
	val := sdk.Provider.MapTagList[key]
	if val == "" {
		return key
	}
	return sdk.Provider.MapTagList[key]
}

func (sdk *SDK) GetMapTagFolderList(key string) string {
	val := sdk.Provider.MapTagFolderList[key]
	if val == "" {
		return key
	}
	return sdk.Provider.MapTagFolderList[key]
}

func (sdk *SDK) GetIsCreateFolder() bool {
	return sdk.Provider.IsCreateFolder
}

func (sdk *SDK) GetIsRenameFolder() bool {
	return sdk.Provider.IsRenameFolder
}

func (sdk *SDK) GetIsManualInstall() bool {
	return sdk.Provider.IsManualInstall
}

func (sdk *SDK) GetIsUseKey() bool {
	return sdk.Provider.IsUseKey
}

func (sdk *SDK) GetDetailKey() string {
	return sdk.Provider.DetailKey
}

type Provider struct {
	IsManualInstall  bool
	IsCreateFolder   bool
	IsRenameFolder   bool
	IsUseKey         bool
	DetailKey        string
	Header           string
	LinkPage         string
	LinkDist         string
	Path             string
	FileName         string
	ZipFolderName    string
	FileType         string
	ArchiveType      string
	MapOSList        map[string]string
	MapArchList      map[string]string
	MapTagList       map[string]string
	MapTagFolderList map[string]string
}

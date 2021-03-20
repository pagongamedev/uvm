package helper

import (
	"path/filepath"
	"strings"
)

func GetFolderVersion(sdkPath string, sVersion string, sTag string) (string, string) {
	sFolderVersion := GetVersionWithV(sVersion)

	if sTag != "" {
		sFolderVersion += "__" + sTag
	}

	sSDKPathVersion := filepath.Join(sdkPath, sFolderVersion)

	return sFolderVersion, sSDKPathVersion
}

func GetVersionTagFromPath(baseFile string) (string, string) {
	sVersion := ""
	sTag := ""
	strList := strings.SplitN(baseFile, "__", 2)

	if len(strList) > 0 {
		sVersion = strList[0]
	}
	if len(strList) > 1 {
		sTag = strList[1]
	}

	return sVersion, sTag
}

func GetVersionWithV(sVersion string) string {
	if sVersion != "" {
		if []rune(sVersion)[0] != 'v' {
			sVersion = "v" + sVersion
		}
	}
	return sVersion
}

func GetVersionWithOutV(sVersion string) string {
	if sVersion != "" {
		if []rune(sVersion)[0] == 'v' {
			sVersion = strings.Replace(sVersion, "v", "", 1)
		}
	}
	return sVersion
}

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/pagongamedev/uvm/download"
	"github.com/pagongamedev/uvm/file"
	"github.com/pagongamedev/uvm/helper"
	"github.com/pagongamedev/uvm/sdk"
	"github.com/pagongamedev/uvm/sdk/dart"
	"github.com/pagongamedev/uvm/sdk/flutter"
	"github.com/pagongamedev/uvm/sdk/golang"
	"github.com/pagongamedev/uvm/sdk/java"
	"github.com/pagongamedev/uvm/sdk/nodejs"
	"github.com/pagongamedev/uvm/sdk/openjava"
	"github.com/pagongamedev/uvm/sdk/python"
	"github.com/pagongamedev/uvm/sdk/ruby"
)

const (
	UVMVersion    = "0.0.1"
	UVMTagVersion = "uvm@" + UVMVersion
	ENVUVMLink    = "UVM_LINK"
)

const (
	InstallPathWindows = "C:\\Program Files"
	InstallPathLinux   = "/usr/local/"
	InstallPathDarwin  = "/usr/local/"
)

const (
	FileJsonChannel = "channel.json"
)

func main() {
	data1 := ""
	data2 := ""
	data3 := ""

	argList := os.Args
	sPlatform := runtime.GOOS
	sArch := runtime.GOARCH

	if len(argList) > 1 {
		if argList[1] == "version" || argList[1] == "v" {
			fmt.Println(UVMTagVersion)
			os.Exit(1)
		}
	}

	if len(argList) < 2 {
		printHelp()
		return
	}

	if len(argList) > 3 {
		data1 = argList[3]
	}
	if len(argList) > 4 {
		data2 = argList[4]
	}
	if len(argList) > 5 {
		data3 = argList[5]
	}

	rootsPath, err := os.Executable()
	MustError(err)
	rootsPath = filepath.Dir(rootsPath)

	sd, err := GetSDK(argList[1], sPlatform)
	MustError(err)

	// check version by use
	if argList[2] == "use" && data1 == "" && data2 == "" {
		basePath, sCurrentVersion, sCurrentTag := getSDKCurrentVersion(*sd, sPlatform)
		if basePath != "" {
			fmt.Println(sCurrentVersion, sCurrentTag)
		} else {
			fmt.Println("(none)")
		}
		return
	}

	// ================
	fmt.Printf("seleted sdk :%v os: %v arch: %v\n", sd.GetName(), sPlatform, sArch)

	RunCommand(argList[2], *sd, data1, data2, data3, rootsPath, sPlatform, sArch)

}

func GetSDK(sSDK string, sPlatform string) (*sdk.SDK, error) {
	var sd *sdk.SDK
	var err error
	switch strings.ToLower(sSDK) {
	case "-d", "dart": // Dart
		sd, _ = dart.NewSDK(sPlatform)
	case "-f", "flutter": // Flutter
		sd, _ = flutter.NewSDK(sPlatform)
	case "-g", "golang": // Golang
		sd, _ = golang.NewSDK(sPlatform)
	case "-j", "java": // Java
		sd, _ = java.NewSDK(sPlatform)
	case "-n", "nodejs": // NodeJS
		sd, _ = nodejs.NewSDK(sPlatform)
	case "-oj", "openjava": // OpenJava
		sd, _ = openjava.NewSDK(sPlatform)
	case "-p", "python": // Python
		sd, _ = python.NewSDK(sPlatform)
	case "-r", "ruby": // Ruby
		sd, _ = ruby.NewSDK(sPlatform)
	case "list", "ls":
		printAllList(sPlatform)

	}

	if sd == nil {
		err = errors.New("not have platform \"" + sSDK + "\"")
		return nil, err
	}

	return sd, nil
}

func printAllList(sPlatform string) {
	fmt.Println(printSDKWithPaddingSpace("UVM", "v"+UVMVersion, ""))
	printVersionQuick(dart.NewSDK, sPlatform)
	printVersionQuick(flutter.NewSDK, sPlatform)
	printVersionQuick(golang.NewSDK, sPlatform)
	printVersionQuick(java.NewSDK, sPlatform)
	printVersionQuick(nodejs.NewSDK, sPlatform)
	printVersionQuick(openjava.NewSDK, sPlatform)
	printVersionQuick(python.NewSDK, sPlatform)
	printVersionQuick(ruby.NewSDK, sPlatform)
	os.Exit(1)
}

func RunCommand(sCommand string, sd sdk.SDK, data1 string, data2 string, data3 string, rootPath string, sPlatform string, sArch string) {
	switch sCommand {
	case "install":
		install(sd, data1, data2, data3, rootPath, sPlatform, sArch)
	case "uninstall":
		uninstall(sd, data1, data2, rootPath, sPlatform)
	case "use":
		use(sd, data1, data2, rootPath, sPlatform)
	case "ls", "list":
		list(sd, rootPath, sPlatform)
	case "unuse":
		unuse(sd, rootPath, sPlatform)
	case "root":
		fmt.Println("current root: " + rootPath)
	case "v", "version":
		fmt.Println(UVMTagVersion)
	default:
		printHelp()
	}
}

func printVersionQuick(sdFunc func(string) (*sdk.SDK, error), sPlatform string) {
	sd, _ := sdFunc(sPlatform)
	_, sCurrentVersion, sCurrentTag := getSDKCurrentVersion(*sd, sPlatform)
	fmt.Println(printSDKWithPaddingSpace(sd.GetName(), sCurrentVersion, sCurrentTag))
}

func paddingSpace(str string, lenght int) string {
	for {
		str += " "
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

func setUrl(sd sdk.SDK, sVersion string, sTag string, sKey string, sPlatform string, sArch string) (string, string, string) {

	sVersion = helper.GetVersionWithOutV(sVersion)

	if sd.GetMapOSList(sPlatform) != "" {
		sPlatform = sd.GetMapOSList(sPlatform)
	}

	if sd.GetMapArchList(sArch) != "" {
		sArch = sd.GetMapArchList(sArch)
	}

	sMapTag := sd.GetMapTagList(sTag)
	sMapTagFolder := sd.GetMapTagFolderList(sTag)

	sFileName := sd.GetFileName()
	sFileName = strings.ReplaceAll(sFileName, "{{version}}", sVersion)
	sFileName = strings.ReplaceAll(sFileName, "{{tag}}", sMapTag)
	sFileName = strings.ReplaceAll(sFileName, "{{tagFolder}}", sMapTagFolder)
	sFileName = strings.ReplaceAll(sFileName, "{{os}}", sPlatform)
	sFileName = strings.ReplaceAll(sFileName, "{{arch}}", sArch)
	sFileName = strings.ReplaceAll(sFileName, "{{key}}", sKey)

	sZipFolderName := sd.GetZipFolderName()
	sZipFolderName = strings.ReplaceAll(sZipFolderName, "{{version}}", sVersion)
	sZipFolderName = strings.ReplaceAll(sZipFolderName, "{{tag}}", sMapTag)
	sZipFolderName = strings.ReplaceAll(sZipFolderName, "{{tagFolder}}", sMapTagFolder)
	sZipFolderName = strings.ReplaceAll(sZipFolderName, "{{os}}", sPlatform)
	sZipFolderName = strings.ReplaceAll(sZipFolderName, "{{arch}}", sArch)
	sZipFolderName = strings.ReplaceAll(sZipFolderName, "{{key}}", sKey)

	sPath := sd.GetPath()
	sPath = strings.ReplaceAll(sPath, "{{fileName}}", sFileName)
	sPath = strings.ReplaceAll(sPath, "{{version}}", sVersion)
	sPath = strings.ReplaceAll(sPath, "{{tag}}", sMapTag)
	sPath = strings.ReplaceAll(sPath, "{{tagFolder}}", sMapTagFolder)
	sPath = strings.ReplaceAll(sPath, "{{os}}", sPlatform)
	sPath = strings.ReplaceAll(sPath, "{{key}}", sKey)
	sPath = strings.ReplaceAll(sPath, "{{type}}", sd.GetFileType())

	sUrl := sd.GetDist() + sPath
	fmt.Println("url : ", sUrl)
	return sUrl, sFileName, sZipFolderName
}

func getSDKCurrentVersion(sd sdk.SDK, sPlatform string) (string, string, string) {
	linkPath := ""
	var err error

	if sd.GetEnvChannel() != "" {
		envChannel, _ := os.LookupEnv(sd.GetEnvChannel())
		if envChannel != sd.GetName() {
			return "", "", ""
		}

	}
	symPath := getSymPath(sd, sPlatform)

	// Search Version
	linkPath, err = os.Readlink(symPath)
	if linkPath == "" || err != nil {
		return "", "", ""
	}

	baseFile := filepath.Base(linkPath)
	sVersion, sTag := helper.GetVersionTagFromPath(baseFile)

	return baseFile, sVersion, sTag
}

func install(sd sdk.SDK, sVersion string, sTag string, sKey string, rootPath string, sPlatform string, sArch string) {
	sVersion = helper.GetVersionWithV(sVersion)
	sdkPath := filepath.Join(rootPath, strings.ToLower(sd.GetName()))
	if !file.IsExist(sdkPath) {
		os.Mkdir(sdkPath, os.ModeDir)
	}

	fmt.Println(sd.GetHeader())
	fmt.Println(sd.GetLinkPage())

	if sd.GetIsManualInstall() {

		fmt.Printf("\nFailed Install : %v not supported cli download\n", sd.GetName())
		fmt.Printf("\nplease download archive at :" + sd.GetLinkPage() + "\n")

		fmt.Printf("and install at : " + filepath.Join(sdkPath, "{{v0.0.0}}") + "\n\n")
		return
	}

	if sTag == "-" {
		sTag = ""
	}
	sFolderVersion, sSDKPathVersion := helper.GetFolderVersion(sdkPath, sVersion, sTag)

	if !file.IsExist(sSDKPathVersion) {
		sUrl, sFileName, sZipFolderName := setUrl(sd, sVersion, sTag, sKey, sPlatform, sArch)

		sTempName := "temp." + sd.GetFileType()
		sTempFile := filepath.Join(sdkPath, sTempName)

		err := download.Loading(sd, rootPath, sdkPath, sUrl, sTempFile, sFileName, sVersion, sTag, sFolderVersion, sSDKPathVersion)
		if err != nil {
			printError(err, "\nPlease Check Archive at :", sd.GetLinkPage())
			return
		}

		// unzip
		err = file.UnArchive(sd.GetArchiveType(), sTempFile, sdkPath, sd.GetIsRenameFolder(), sd.GetIsCreateFolder(), sZipFolderName, sFolderVersion)
		if err != nil {
			fmt.Println("Unzip Error ", err)
			return
		}

		// remove Temp
		err = os.Remove(sTempFile)
		if err != nil {
			fmt.Println("Error Delete Temp File", err)
			return
		}
		fmt.Println("installed.")
		fmt.Println()
		fmt.Println("please run Command:", "uvm", sd.GetCommand(), "use", sVersion, sTag)
	} else {
		fmt.Println("already installed :", sd.GetName(), sVersion, sTag)
	}

}

func uninstall(sd sdk.SDK, sVersion string, sTag string, rootPath string, sPlatform string) {
	sVersion = helper.GetVersionWithV(sVersion)

	// Make sure a version is specified
	if sVersion == "" && sTag == "" {
		printError("Provide the version you want to uninstall.")
		printHelp()
		return
	}

	// check current use and remove symlink when use
	_, sCurrentVersion, sCurrentTag := getSDKCurrentVersion(sd, sPlatform)

	if sCurrentVersion == sVersion && sCurrentTag == sTag {
		unuse(sd, rootPath, sPlatform)
	}

	// fmt.Println(sCurrentVersion, " == ", sVersion, " && ", sCurrentTag, " == ", sTag)
	sdkPath := filepath.Join(rootPath, strings.ToLower(sd.GetName()))
	_, sSDKPathVersion := helper.GetFolderVersion(sdkPath, sVersion, sTag)

	if file.IsExist(sSDKPathVersion) {
		fmt.Println("uninstalling", sd.GetName(), sVersion, sTag)

		err := os.RemoveAll(sSDKPathVersion)
		if err != nil {
			fmt.Println("error removing ", sd.GetName(), sVersion, sTag)
			fmt.Println("manually remove " + sSDKPathVersion)
		} else {
			fmt.Printf("\nuninstalled.")
		}
	} else {
		fmt.Println("not have installed :", sd.GetName(), sVersion, sTag)
	}
}

func list(sd sdk.SDK, rootPath string, sPlatform string) {

	baseFile, _, _ := getSDKCurrentVersion(sd, sPlatform)

	sPre := ""
	sPost := ""

	sdkPath := filepath.Join(rootPath, strings.ToLower(sd.GetName()))

	// versionList := []string{}
	reg, _ := regexp.Compile("v")

	files, _ := ioutil.ReadDir(sdkPath)
	for i := len(files) - 1; i >= 0; i-- {
		if files[i].IsDir() {
			if reg.MatchString(files[i].Name()) {

				if baseFile == files[i].Name() {
					sPre = "*"
					sPost = "(Currently)"
				} else {
					sPre = " "
					sPost = ""
				}
				sVersion, sTag := helper.GetVersionTagFromPath(files[i].Name())

				fmt.Printf(" %v %v %v %v\n", sPre, sVersion, sTag, sPost)
				// versionList = append(versionList, files[i].Name())
			}
		}
	}

}

func use(sd sdk.SDK, sVersion string, sTag string, rootPath string, sPlatform string) {
	// remove symlink if it already exists
	removeSymLink(sd, rootPath, sPlatform)

	// create symlink
	sVersion = helper.GetVersionWithV(sVersion)

	sdkPath := filepath.Join(rootPath, strings.ToLower(sd.GetName()))
	_, sSDKPathVersion := helper.GetFolderVersion(sdkPath, sVersion, sTag)

	if !file.IsExist(sSDKPathVersion) {
		fmt.Println("error : not have installed :", sd.GetName(), sVersion, sTag)
		return
	}
	symPath := getSymPath(sd, sPlatform)
	switch sPlatform {
	case "windows":
		// create symlink
		if !helper.RunCommand(fmt.Sprintf(`"%s" cmd /C mklink /D "%s" "%s"`,
			filepath.Join(rootPath, "bin", "elevate.cmd"),
			filepath.Clean(symPath),
			sSDKPathVersion)) {
			return
		}
	case "darwin", "linux":
		err := os.Symlink(sSDKPathVersion, symPath)
		if err != nil {
			fmt.Println("Symlink Error ", err)
			return
		}
	}

	fmt.Printf("create symlink\n\n")
	fmt.Printf("use %v %v %v\n\n", sd.GetName(), sVersion, sTag)

	// check env
	isUpdateEnv := false

	// check UVMLink Env
	uvmlinkData, ok := os.LookupEnv(ENVUVMLink)
	if !ok || !strings.Contains(uvmlinkData, symPath) {
		CheckOrCreateEnv(sd, true, sPlatform, ENVUVMLink, uvmlinkData, rootPath, symPath)
		isUpdateEnv = true
	}

	// Set Env
	if sd.GetEnv() != "" {
		uvmSDKData, ok := os.LookupEnv(sd.GetEnv())
		if !ok || !strings.Contains(uvmSDKData, symPath) {
			isUpdateEnv = true
		}

		CheckOrCreateEnv(sd, false, sPlatform, sd.GetEnv(), uvmSDKData, rootPath, symPath)
	}

	// Set Env Channel
	jsonChannelFilePath := filepath.Clean(filepath.Join(rootPath, FileJsonChannel))

	if sd.GetEnvChannel() != "" {
		channelDataList := file.ReadJSONFile(jsonChannelFilePath)
		sdkChannelData := channelDataList[sd.GetEnvChannel()]

		isSameSDK := false
		if sdkChannelData != nil {
			if sdkChannelData.(string) == sd.GetName() {
				isSameSDK = true
			}
		}

		if !isSameSDK {
			channelDataList[sd.GetEnvChannel()] = sd.GetName()
			file.WriteJSONFile(jsonChannelFilePath, channelDataList)
		}
	}

	if isUpdateEnv {
		switch sPlatform {
		case "windows":
			fmt.Println("env updated please restart shell")
		case "darwin", "linux":
		}
	} else {
		// Check uvmlink in path
		path := os.Getenv("path")
		if !strings.Contains(path, symPath) {
			switch sPlatform {
			case "windows":
				fmt.Println("please add env : PATH = %" + ENVUVMLink + "% and restart shell")
			case "darwin", "linux":
				fmt.Println("Please Append \"" + ENVUVMLink + "\"")
				fmt.Println("In \"/etc/profile.d/uvm.sh\" at export PATH")
			}

		}
	}
}
func CheckOrCreateEnv(sd sdk.SDK, isAppend bool, sPlatform string, envName string, envData string, rootPath string, symPath string) {

	symPathBin := filepath.Clean(filepath.Join(symPath, sd.GetEnvBin()))

	switch sPlatform {
	case "windows":
		sPre := ""
		if envData != "" && isAppend {
			sPre = envData + ";"
		}
		if !helper.RunCommand(fmt.Sprintf(`"%s" cmd /C SETX /M "%s" "%s"`,
			filepath.Join(rootPath, "bin", "elevate.cmd"),
			envName,
			sPre+symPathBin)) {
			return
		}
	case "darwin", "linux":
		strAdd := "Add"
		if isAppend {
			strAdd = "Append"
		}

		fmt.Println("Env:" + envName + " Not Found.")
		fmt.Println("Please " + strAdd + " \"" + symPathBin + "\"")
		fmt.Println("In \"/etc/profile.d/uvm.sh\" at export " + envName)
	}
}

func unuse(sd sdk.SDK, rootPath string, sPlatform string) {
	removeSymLink(sd, rootPath, sPlatform)
	fmt.Printf("remove symlink\n\n")
}

func removeSymLink(sd sdk.SDK, rootPath string, sPlatform string) {
	symPath := getSymPath(sd, sPlatform)
	switch sPlatform {
	case "windows":
		// remove symlink if it already exists
		sym, _ := os.Stat(symPath)
		if sym != nil {
			if !helper.RunCommand(fmt.Sprintf(`"%s" cmd /C rmdir "%s"`,
				filepath.Join(rootPath, "bin", "elevate.cmd"),
				filepath.Clean(symPath))) {
				return
			}
		}
	case "darwin", "linux":
		// remove symlink if it already exists
		sym, _ := os.Stat(symPath)
		if sym != nil {
			err := os.Remove(symPath)
			if err != nil {
				fmt.Println("Symlink Remove Error ", err)
				return
			}
		}
	}
}

func ConvertLinkName(sName string) string {
	return "uvm_" + strings.ToLower(sName)
}

func printSDKWithPaddingSpace(sName string, sVersion string, sTag string) string {
	return paddingSpace(sName, 10) + ": " + sVersion + " " + sTag
}

func getSymPath(sd sdk.SDK, sPlatform string) string {
	symPath := ""
	switch sPlatform {
	case "windows":
		symPath = filepath.Join(InstallPathWindows, ConvertLinkName(sd.GetLinkName()))
	case "darwin":
		symPath = filepath.Join(InstallPathDarwin, ConvertLinkName(sd.GetLinkName()))
	case "linux":
		symPath = filepath.Join(InstallPathLinux, ConvertLinkName(sd.GetLinkName()))
	}

	return symPath
}

// =====================================================================
//                              Add On
// =====================================================================

func printError(iText ...interface{}) {
	fmt.Println("=======================================")
	fmt.Println()
	fmt.Println(iText...)
	fmt.Println()
	fmt.Println("=======================================")
}

// MustError Func
func MustError(err error, strList ...string) {
	if err != nil {
		if strList != nil {
			printError(strList)
			printHelp()
			os.Exit(1)
		} else {
			printError("error :", err)
			printHelp()
			os.Exit(1)
		}
	}
}

func printHelp() {
	fmt.Println("\nRunning version " + UVMTagVersion + ".")
	fmt.Println("\nOS : " + runtime.GOOS + " Arch : " + runtime.GOARCH + ".")

	fmt.Println("\nSupport:")
	fmt.Println("                                        |     Window     |     Darwin     |     Linux")
	fmt.Println("  uvm -d  , uvm dart        : Dart           Suported")
	fmt.Println("  uvm -f  , uvm flutter     : Flutter        Suported")
	fmt.Println("  uvm -go , uvm golang      : Golang         Suported")
	fmt.Println("  uvm -j  , uvm java        : Java         [Manual Ins.]")
	fmt.Println("  uvm -n  , uvm nodejs      : NodeJS          Suported")
	fmt.Println("  uvm -oj , uvm openjava    : OpenJava       [Use Key]")
	fmt.Println("  uvm -p  , uvm python      : Python       [Manual Ins.]")
	fmt.Println("  uvm -r  , uvm ruby        : Ruby         [Manual Ins.]")
	fmt.Println("\nUsage:")
	fmt.Println(" ")
	fmt.Println("  uvm [-SDK] install <version> <tag>    : Install SDK Version.")
	fmt.Println("  uvm [-SDK] uninstall <version>        : The version must be a specific version.")
	fmt.Println("  uvm [-SDK] list                       : List Version Installed and Show Current Use")
	fmt.Println("  uvm [-SDK] use <version> <tag> <key>  : Switch to use the specified version.")
	fmt.Println("                                          <tag> for channel look like dev , beta ")
	fmt.Println("                                          <key> for download link with random string")
	fmt.Println("  uvm [-SDK] unuse                      : Disable uvm.")
	fmt.Println("  uvm [-SDK] root                       : Show Root Path")
	fmt.Println("  uvm [-SDK] version                    : Displays the current running version of uvm")
}

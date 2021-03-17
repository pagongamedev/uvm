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
	"github.com/pagongamedev/uvm/repository"
	"github.com/pagongamedev/uvm/repository/dart"
	"github.com/pagongamedev/uvm/repository/flutter"
	"github.com/pagongamedev/uvm/repository/golang"
	"github.com/pagongamedev/uvm/repository/java"
	"github.com/pagongamedev/uvm/repository/nodejs"
	"github.com/pagongamedev/uvm/repository/openjava"
	"github.com/pagongamedev/uvm/repository/python"
)

const (
	UVMVersion = "uvm@v0.0.1"
	ENVUVMLink = "UVM_LINK"
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
			fmt.Println(UVMVersion)
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

	repo, err := GetRepository(argList[1], sPlatform)
	MustError(err)

	// check version by use
	if argList[2] == "use" && data1 == "" && data2 == "" {
		basePath, sCurrentVersion, sCurrentTag := getSDKCurrentVersion(repo, sPlatform)
		if basePath != "" {
			fmt.Println(sCurrentVersion, sCurrentTag)
		} else {
			fmt.Println("(none)")
		}
		return
	}

	// ================
	fmt.Printf("seleted sdk :%v os: %v arch: %v\n", repo.GetName(), sPlatform, sArch)

	RunCommand(argList[2], repo, data1, data2, data3, rootsPath, sPlatform, sArch)

}

func GetRepository(sSDK string, sPlatform string) (repository.Repository, error) {
	var repo repository.Repository
	var err error
	switch strings.ToLower(sSDK) {
	case "-d": // Dart
		repo, _ = dart.NewRepository(sPlatform)
	case "-f": // Flutter
		repo, _ = flutter.NewRepository(sPlatform)
	case "-g": // Golang
		repo, _ = golang.NewRepository(sPlatform)
	case "-j": // Java
		repo, _ = java.NewRepository(sPlatform)
	case "-n": // NodeJS
		repo, _ = nodejs.NewRepository(sPlatform)
	case "-oj": // OpenJava
		repo, _ = openjava.NewRepository(sPlatform)
	case "-p": // Python
		repo, _ = python.NewRepository(sPlatform)
	//// case "-r": // Ruby
	// 	// repo, _ = ruby.NewRepository(sPlatform)
	case "list":
		printVersionQuick(dart.NewRepository, sPlatform)
		printVersionQuick(flutter.NewRepository, sPlatform)
		printVersionQuick(golang.NewRepository, sPlatform)
		printVersionQuick(java.NewRepository, sPlatform)
		printVersionQuick(nodejs.NewRepository, sPlatform)
		printVersionQuick(openjava.NewRepository, sPlatform)
		printVersionQuick(python.NewRepository, sPlatform)
		// printVersionQuick(ruby.NewRepository, sPlatform)

		os.Exit(1)
	}

	if repo == nil {
		err = errors.New("not have platform \"" + sSDK + "\"")
	}

	return repo, err
}

func RunCommand(sCommand string, repo repository.Repository, data1 string, data2 string, data3 string, rootPath string, sPlatform string, sArch string) {
	switch sCommand {
	case "install":
		install(repo, data1, data2, data3, rootPath, sPlatform, sArch)
	case "uninstall":
		uninstall(repo, data1, data2, rootPath, sPlatform)
	case "use":
		use(repo, data1, data2, rootPath, sPlatform)
	case "list":
		list(repo, rootPath, sPlatform)
	case "ls":
		list(repo, rootPath, sPlatform)
	case "unuse":
		unuse(repo, rootPath, sPlatform)
	case "root":
		fmt.Println("current root: " + rootPath)
	case "version":
		fmt.Println(UVMVersion)
	case "v":
		fmt.Println(UVMVersion)
	default:
		printHelp()
	}
}

func printVersionQuick(repoFunc func(string) (repository.Repository, error), sPlatform string) {
	repo, _ := repoFunc(sPlatform)
	sdkName := paddingSpace(repo.GetName(), 8)
	_, sCurrentVersion, sCurrentTag := getSDKCurrentVersion(repo, sPlatform)
	fmt.Println(sdkName+": "+sCurrentVersion, sCurrentTag)
}

func paddingSpace(str string, lenght int) string {
	for {
		str += " "
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

func setUrl(repo repository.Repository, sVersion string, sTag string, sKey string, sPlatform string, sArch string) (string, string, string) {

	sVersion = helper.GetVersionWithOutV(sVersion)

	if repo.GetMapOSList(sPlatform) != "" {
		sPlatform = repo.GetMapOSList(sPlatform)
	}

	if repo.GetMapArchList(sArch) != "" {
		sArch = repo.GetMapArchList(sArch)
	}

	sMapTag := repo.GetMapTagList(sTag)
	sMapTagFolder := repo.GetMapTagFolderList(sTag)

	sFileName := repo.GetFileName()
	sFileName = strings.ReplaceAll(sFileName, "{{version}}", sVersion)
	sFileName = strings.ReplaceAll(sFileName, "{{tag}}", sMapTag)
	sFileName = strings.ReplaceAll(sFileName, "{{tagFolder}}", sMapTagFolder)
	sFileName = strings.ReplaceAll(sFileName, "{{os}}", sPlatform)
	sFileName = strings.ReplaceAll(sFileName, "{{arch}}", sArch)
	sFileName = strings.ReplaceAll(sFileName, "{{key}}", sKey)

	sZipFolderName := repo.GetZipFolderName()
	sZipFolderName = strings.ReplaceAll(sZipFolderName, "{{version}}", sVersion)
	sZipFolderName = strings.ReplaceAll(sZipFolderName, "{{tag}}", sMapTag)
	sZipFolderName = strings.ReplaceAll(sZipFolderName, "{{tagFolder}}", sMapTagFolder)
	sZipFolderName = strings.ReplaceAll(sZipFolderName, "{{os}}", sPlatform)
	sZipFolderName = strings.ReplaceAll(sZipFolderName, "{{arch}}", sArch)
	sZipFolderName = strings.ReplaceAll(sZipFolderName, "{{key}}", sKey)

	sPath := repo.GetPath()
	sPath = strings.ReplaceAll(sPath, "{{fileName}}", sFileName)
	sPath = strings.ReplaceAll(sPath, "{{version}}", sVersion)
	sPath = strings.ReplaceAll(sPath, "{{tag}}", sMapTag)
	sPath = strings.ReplaceAll(sPath, "{{tagFolder}}", sMapTagFolder)
	sPath = strings.ReplaceAll(sPath, "{{os}}", sPlatform)
	sPath = strings.ReplaceAll(sPath, "{{key}}", sKey)
	sPath = strings.ReplaceAll(sPath, "{{type}}", repo.GetFileType())

	sUrl := repo.GetDist() + sPath
	fmt.Println("url : ", sUrl)
	return sUrl, sFileName, sZipFolderName
}

func getSDKCurrentVersion(repo repository.Repository, sPlatform string) (string, string, string) {
	linkPath := ""
	var err error

	if repo.GetEnvChannel() != "" {
		envChannel, _ := os.LookupEnv(repo.GetEnvChannel())
		if envChannel != repo.GetName() {
			return "", "", ""
		}

	}
	// Search Version
	switch sPlatform {
	case "windows":
		symPath := filepath.Join("C:\\Program Files", "UVM_"+repo.GetLinkName())
		linkPath, err = os.Readlink(symPath)

	}

	if err != nil {
		return "", "", ""
	}

	baseFile := filepath.Base(linkPath)
	sVersion, sTag := helper.GetVersionTagFromPath(baseFile)

	return baseFile, sVersion, sTag
}

func install(repo repository.Repository, sVersion string, sTag string, sKey string, rootPath string, sPlatform string, sArch string) {
	sVersion = helper.GetVersionWithV(sVersion)
	sdkPath := filepath.Join(rootPath, repo.GetName())
	if !file.IsExist(sdkPath) {
		os.Mkdir(sdkPath, os.ModeDir)
	}

	if repo.GetIsManualInstall() {
		fmt.Printf("\nFailed Install : %v not supported cli download\n", repo.GetName())
		fmt.Printf("\nplease download archive at :" + repo.GetLinkPage() + "\n")
		fmt.Printf("and install at :" + sdkPath + "\\{{v0.0.0}}\\\n\n")
		return
	}

	if sTag == "-" {
		sTag = ""
	}
	sFolderVersion, sSDKPathVersion := helper.GetFolderVersion(sdkPath, sVersion, sTag)

	if !file.IsExist(sSDKPathVersion) {
		sUrl, sFileName, sZipFolderName := setUrl(repo, sVersion, sTag, sKey, sPlatform, sArch)

		sTempFile, err := download.Loading(repo, rootPath, sdkPath, sUrl, sFileName, sVersion, sTag, sFolderVersion, sSDKPathVersion)
		if err != nil {
			printError(err, "\nPlease Check Archive at :", repo.GetLinkPage())
			return
		}

		// unzip
		err = file.UnArchive(repo.GetArchiveType(), sTempFile, sdkPath, repo.GetIsRenameFolder(), sZipFolderName, sFolderVersion)
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
		fmt.Println("please run command:", "uvm", repo.GetCommand(), "use", sVersion, sTag)
	} else {
		fmt.Println("already installed :", repo.GetName(), sVersion, sTag)
	}

}

func uninstall(repo repository.Repository, sVersion string, sTag string, rootPath string, sPlatform string) {
	sVersion = helper.GetVersionWithV(sVersion)

	// Make sure a version is specified
	if sVersion == "" && sTag == "" {
		printError("Provide the version you want to uninstall.")
		printHelp()
		return
	}

	// check current use and remove symlink when use
	_, sCurrentVersion, sCurrentTag := getSDKCurrentVersion(repo, sPlatform)

	if sCurrentVersion == sVersion && sCurrentTag == sTag {
		unuse(repo, rootPath, sPlatform)
	}

	// fmt.Println(sCurrentVersion, " == ", sVersion, " && ", sCurrentTag, " == ", sTag)
	sdkPath := filepath.Join(rootPath, repo.GetName())
	_, sSDKPathVersion := helper.GetFolderVersion(sdkPath, sVersion, sTag)

	if file.IsExist(sSDKPathVersion) {
		fmt.Println("uninstalling", repo.GetName(), sVersion, sTag)

		err := os.RemoveAll(sSDKPathVersion)
		if err != nil {
			fmt.Println("error removing ", repo.GetName(), sVersion, sTag)
			fmt.Println("manually remove " + sSDKPathVersion)
		} else {
			fmt.Printf("\nuninstalled.")
		}
	} else {
		fmt.Println("not have installed :", repo.GetName(), sVersion, sTag)
	}
}

func list(repo repository.Repository, rootPath string, sPlatform string) {

	baseFile, _, _ := getSDKCurrentVersion(repo, sPlatform)

	sPre := ""
	sPost := ""

	sdkPath := filepath.Join(rootPath, repo.GetName())

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

func use(repo repository.Repository, sVersion string, sTag string, rootPath string, sPlatform string) {
	// remove symlink if it already exists
	removeSymLink(repo, rootPath, sPlatform)

	// create symlink
	sVersion = helper.GetVersionWithV(sVersion)

	sdkPath := filepath.Join(rootPath, repo.GetName())
	_, sSDKPathVersion := helper.GetFolderVersion(sdkPath, sVersion, sTag)

	if !file.IsExist(sSDKPathVersion) {
		fmt.Println("error : not have installed :", repo.GetName(), sVersion, sTag)
		return
	}

	switch sPlatform {
	case "windows":

		symPath := filepath.Join("C:\\Program Files", "UVM_"+repo.GetLinkName())

		// create symlink
		if !helper.RunCommand(fmt.Sprintf(`"%s" cmd /C mklink /D "%s" "%s"`,
			filepath.Join(rootPath, "bin", "elevate.cmd"),
			filepath.Clean(symPath),
			sSDKPathVersion)) {
			return
		}

		// err := os.Symlink(symPath, sSDKPathVersion)
		// if err != nil {
		// 	fmt.Println("Symlink Error ", err)
		// 	return
		// }

		fmt.Printf("create symlink\n\n")
	}
	fmt.Printf("use %v %v %v\n\n", repo.GetName(), sVersion, sTag)

	// check env
	switch sPlatform {
	case "windows":
		// sheck repo env
		isUpdateEnv := false
		symPath := filepath.Join("C:\\Program Files", "UVM_"+repo.GetLinkName())

		uvmlink, ok := os.LookupEnv(ENVUVMLink)
		if !ok {
			createEnvUVMLink(repo, uvmlink, rootPath, symPath)
			isUpdateEnv = true
		} else {
			if !strings.Contains(uvmlink, symPath) {
				createEnvUVMLink(repo, uvmlink, rootPath, symPath)
				isUpdateEnv = true
			}
		}

		// Set Env
		if repo.GetEnv() != "" {
			_, ok := os.LookupEnv(repo.GetEnv())
			if !ok {
				isUpdateEnv = true
			}

			if !helper.RunCommand(fmt.Sprintf(`"%s" cmd /C SETX /M "%s" "%s"`,
				filepath.Join(rootPath, "bin", "elevate.cmd"),
				repo.GetEnv(),
				filepath.Clean(symPath+repo.GetEnvBin()))) {
				return
			}
		}

		// Set Env Channel
		if repo.GetEnvChannel() != "" {
			_, ok := os.LookupEnv(repo.GetEnvChannel())
			if !ok {
				isUpdateEnv = true
			}

			if !helper.RunCommand(fmt.Sprintf(`"%s" cmd /C SETX /M "%s" "%s"`,
				filepath.Join(rootPath, "bin", "elevate.cmd"),
				repo.GetEnvChannel(),
				repo.GetName())) {
				return
			}
		}

		if isUpdateEnv {
			fmt.Println("env updated please restart shell")
		} else {
			// Check uvmlink in path
			path := os.Getenv("path")
			if !strings.Contains(path, symPath) {
				fmt.Println("please add env : PATH = %" + ENVUVMLink + "% and restart shell")
			}
		}
	}
}
func createEnvUVMLink(repo repository.Repository, uvmlink string, rootPath string, symPath string) {

	sPre := ""
	if uvmlink != "" {
		sPre = ";"
	}

	if !helper.RunCommand(fmt.Sprintf(`"%s" cmd /C SETX /M "%s" "%s"`,
		filepath.Join(rootPath, "bin", "elevate.cmd"),
		ENVUVMLink,
		uvmlink+sPre+filepath.Clean(symPath+repo.GetEnvBin()))) {
		return
	}
}

func unuse(repo repository.Repository, rootPath string, sPlatform string) {
	removeSymLink(repo, rootPath, sPlatform)
	fmt.Printf("remove symlink\n\n")
}

func removeSymLink(repo repository.Repository, rootPath string, sPlatform string) {
	switch sPlatform {
	case "windows":
		symPath := filepath.Join("C:\\Program Files", "UVM_"+repo.GetLinkName())

		// remove symlink if it already exists
		sym, _ := os.Stat(symPath)
		if sym != nil {
			if !helper.RunCommand(fmt.Sprintf(`"%s" cmd /C rmdir "%s"`,
				filepath.Join(rootPath, "bin", "elevate.cmd"),
				filepath.Clean(symPath))) {
				return
			}
		}

	}
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
	fmt.Println("\nRunning version " + UVMVersion + ".")
	fmt.Println("\nOS : " + runtime.GOOS + " Arch : " + runtime.GOARCH + ".")

	fmt.Println("\nSupport:")
	fmt.Println(" ")
	fmt.Println("  uvm -d            : Dart")
	fmt.Println("  uvm -f            : Flutter")
	fmt.Println("  uvm -g            : Golang")
	fmt.Println("  uvm -j            : Java         [Manual Install]")
	fmt.Println("  uvm -n            : NodeJS")
	fmt.Println("  uvm -oj           : OpenJava     [Use Key]")
	fmt.Println("  uvm -p            : Python       [Manual Install]")
	// fmt.Println("  uvm -r            : Ruby")
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

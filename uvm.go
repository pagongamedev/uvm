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
	"github.com/pagongamedev/uvm/repository/nodejs"
)

const (
	UVMVersion = "uvm@v0.0.1"
	ENVUVMLink = "UVM_LINK"
)

func main() {
	data1 := ""
	data2 := ""

	argList := os.Args
	sPlatfrom := runtime.GOOS
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

	rootsPath, err := os.Executable()
	MustError(err)
	rootsPath = filepath.Dir(rootsPath)

	repo, err := GetRepository(argList[1], sPlatfrom)
	MustError(err)

	// check version by use
	if argList[2] == "use" && data1 == "" && data2 == "" {
		basePath, sCurrentVersion, sCurrentTag := getSDKCurrentVersion(repo, sPlatfrom)
		if basePath != "" {
			fmt.Println(sCurrentVersion, sCurrentTag)
		} else {
			fmt.Println("(none)")
		}
		return
	}

	// ================
	fmt.Printf("\nos: %v arch: %v\n", sPlatfrom, sArch)
	fmt.Printf("seleted sdk : " + repo.GetName() + "\n\n")

	RunCommand(argList[2], repo, data1, data2, rootsPath, sPlatfrom, sArch)

}

func printHelp() {
	fmt.Println("\nRunning version " + UVMVersion + ".")
	fmt.Println("\nUsage:")
	fmt.Println(" ")
	fmt.Println("  uvm arch                     : Show if is running in 32 or 64 bit mode.")
}

func GetRepository(sSDK string, sPlatform string) (repository.Repository, error) {
	var repo repository.Repository
	var err error
	switch strings.ToLower(sSDK) {
	case "-d": // Dart
		// repo, _ = dart.NewRepository(sPlatform)
	case "-f": // Flutter
		// repo, _ = flutter.NewRepository(sPlatform)
	case "-g": // Golang
		// repo, _ = golang.NewRepository(sPlatform)
	case "-j": // Java
		// repo, _ = java.NewRepository(sPlatform)
	case "-n": // NodeJS
		repo, err = nodejs.NewRepository(sPlatform)
	case "-p": // Python
		// repo, _ = python.NewRepository(sPlatform)
	case "-r": // Ruby
		// repo, _ = ruby.NewRepository(sPlatform)
	case "-t": // Terraform
		// repo, _ = terraform.NewRepository(sPlatform)
	case "list":

	}

	if repo == nil {
		err = errors.New("not have platform \"" + sSDK + "\"")
	}

	return repo, err
}

func RunCommand(sCommand string, repo repository.Repository, data1 string, data2 string, rootPath string, sPlatfrom string, sArch string) {
	switch sCommand {
	case "install":
		install(repo, data1, data2, rootPath, sPlatfrom, sArch)
	case "uninstall":
		uninstall(repo, data1, data2, rootPath, sPlatfrom)
	case "use":
		use(repo, data1, data2, rootPath, sPlatfrom)
	case "list":
		list(repo, rootPath, sPlatfrom)
	case "ls":
		list(repo, rootPath, sPlatfrom)
	case "unuse":
		unuse(repo, rootPath, sPlatfrom)
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

func setUrl(repo repository.Repository, sVersion string, sTag string, sPlatfrom string, sArch string) (string, string) {

	if repo.GetMapList(sPlatfrom) != "" {
		sPlatfrom = repo.GetMapList(sPlatfrom)
	}

	if repo.GetMapList(sArch) != "" {
		sArch = repo.GetMapList(sArch)
	}

	sFileName := repo.GetFileName()
	sFileName = strings.ReplaceAll(sFileName, "{{version}}", sVersion)
	sFileName = strings.ReplaceAll(sFileName, "{{tag}}", sTag)
	sFileName = strings.ReplaceAll(sFileName, "{{os}}", sPlatfrom)
	sFileName = strings.ReplaceAll(sFileName, "{{arch}}", sArch)

	sPath := repo.GetPath()
	sPath = strings.ReplaceAll(sPath, "{{fileName}}", sFileName)
	sPath = strings.ReplaceAll(sPath, "{{version}}", sVersion)
	sPath = strings.ReplaceAll(sPath, "{{tag}}", sTag)
	sPath = strings.ReplaceAll(sPath, "{{type}}", repo.GetFileType())

	sUrl := repo.GetDist() + sPath

	return sUrl, sFileName
}

func getSDKCurrentVersion(repo repository.Repository, sPlatfrom string) (string, string, string) {
	linkPath := ""
	var err error

	// Search Version
	switch sPlatfrom {
	case "windows":
		symPath := filepath.Join("C:\\Program Files", "UVM_"+repo.GetName())
		linkPath, err = os.Readlink(symPath)

	}

	if err != nil {
		return "", "", ""
	}

	// if !ok {
	// 	fmt.Println()
	// 	fmt.Println("env: \"" + repo.GetEnv() + "\" not Found.")
	// 	fmt.Println("please  set  env :", repo.GetEnv(), "=", symPath)
	// 	fmt.Println("and add path env : %" + repo.GetEnv() + "%" + repo.GetEnvBin())
	// 	fmt.Println()
	// }
	baseFile := filepath.Base(linkPath)
	sVersion, sTag := helper.GetVersionTagFromPath(baseFile)

	return baseFile, sVersion, sTag
}

func install(repo repository.Repository, sVersion string, sTag string, rootPath string, sPlatfrom string, sArch string) {

	sVersion = helper.GetVersionWithV(sVersion)

	sdkPath := filepath.Join(rootPath, repo.GetName())
	if !file.IsExist(sdkPath) {
		os.Mkdir(sdkPath, os.ModeDir)
	}

	sFolderVersion, sSDKPathVersion := helper.GetFolderVersion(sdkPath, sVersion, sTag)

	if !file.IsExist(sSDKPathVersion) {
		sUrl, sFileName := setUrl(repo, sVersion, sTag, sPlatfrom, sArch)

		sTempFile, err := download.Loading(repo, rootPath, sdkPath, sUrl, sFileName, sVersion, sTag, sFolderVersion, sSDKPathVersion)
		if err != nil {
			printError(err)
			return
		}

		// unzip
		err = file.UnArchive(repo.GetArchiveType(), sTempFile, sdkPath, repo.GetIsRenameFolder(), sFileName, sFolderVersion)
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
		fmt.Println("please run command:", "nvm", repo.GetCommand(), "use", sVersion, sTag)
	} else {
		fmt.Println("already installed :", repo.GetName(), sVersion, sTag)
	}

}

func uninstall(repo repository.Repository, sVersion string, sTag string, rootPath string, sPlatfrom string) {
	sVersion = helper.GetVersionWithV(sVersion)

	// Make sure a version is specified
	if sVersion == "" && sTag == "" {
		printError("Provide the version you want to uninstall.")
		printHelp()
		return
	}

	// check current use and remove symlink when use
	_, sCurrentVersion, sCurrentTag := getSDKCurrentVersion(repo, sPlatfrom)

	if sCurrentVersion == sVersion && sCurrentTag == sTag {
		unuse(repo, rootPath, sPlatfrom)
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

func list(repo repository.Repository, rootPath string, sPlatfrom string) {

	baseFile, _, _ := getSDKCurrentVersion(repo, sPlatfrom)

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

func use(repo repository.Repository, sVersion string, sTag string, rootPath string, sPlatfrom string) {
	// remove symlink if it already exists
	removeSymLink(repo, rootPath, sPlatfrom)

	// create symlink
	sVersion = helper.GetVersionWithV(sVersion)

	sdkPath := filepath.Join(rootPath, repo.GetName())
	_, sSDKPathVersion := helper.GetFolderVersion(sdkPath, sVersion, sTag)

	if !file.IsExist(sSDKPathVersion) {
		fmt.Println("error : not have installed :", repo.GetName(), sVersion, sTag)
		return
	}

	switch sPlatfrom {
	case "windows":

		symPath := filepath.Join("C:\\Program Files", "UVM_"+repo.GetName())

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
	switch sPlatfrom {
	case "windows":
		// sheck repo env
		isUpdateEnv := false
		symPath := filepath.Join("C:\\Program Files", "UVM_"+repo.GetName())
		uvmlink, ok := os.LookupEnv(ENVUVMLink)
		if !ok {
			if !helper.RunCommand(fmt.Sprintf(`"%s" cmd /C SETX /M "%s" "%s"`,
				filepath.Join(rootPath, "bin", "elevate.cmd"),
				ENVUVMLink,
				filepath.Clean(symPath))) {
				return
			}

			isUpdateEnv = true
		}

		if !strings.Contains(uvmlink, symPath) {

			if !helper.RunCommand(fmt.Sprintf(`"%s" cmd /C SETX /M "%s" "%s"`,
				filepath.Join(rootPath, "bin", "elevate.cmd"),
				ENVUVMLink,
				";"+filepath.Clean(symPath))) {
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

	// Arch Adapter
	// Zip Selector
	//Dowmload
	// https://nodejs.org/dist/v14.16.0/node-v14.16.0-win-x64.zip
	// https://nodejs.org/dist/v14.16.0/node-v14.16.0-win-x64.7z
	// node-v14.16.0-darwin-x64.tar.gz                    23-Feb-2021 00:29            31567754
	// node-v14.16.0-darwin-x64.tar.xz
	// https://nodejs.org/dist/latest-v14.x/

	// root: D:\SDK\Node\nvm
	// path: C:\Program Files\nodejs
	// https://nodejs.org/dist/
	// node-v14.16.0-darwin-x64.tar.gz
	//

	// https://storage.googleapis.com/flutter_infra/releases/stable/windows/flutter_windows_2.0.2-stable.zip

}

func unuse(repo repository.Repository, rootPath string, sPlatfrom string) {
	removeSymLink(repo, rootPath, sPlatfrom)
	fmt.Printf("remove symlink\n\n")
}

func removeSymLink(repo repository.Repository, rootPath string, sPlatfrom string) {
	switch sPlatfrom {
	case "windows":
		symPath := filepath.Join("C:\\Program Files", "UVM_"+repo.GetName())

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

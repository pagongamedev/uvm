package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pagongamedev/uvm/download"
	"github.com/pagongamedev/uvm/file"
	"github.com/pagongamedev/uvm/repository"
	"github.com/pagongamedev/uvm/repository/nodejs"
)

const (
	UVMVersion = "0.0.1"
)

func main() {
	data1 := ""
	data2 := ""

	argList := os.Args
	sPlatfrom := runtime.GOOS
	sArch := runtime.GOARCH

	log.Printf("args %v\n", argList)
	fmt.Printf("\nos: %v arch: %v\n", sPlatfrom, sArch)

	if len(argList) < 2 {
		helper()
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

	if argList[1] == "version" || argList[1] == "v" {
		fmt.Println(UVMVersion)
		os.Exit(1)
	}

	repo, err := GetRepository(argList[1], sPlatfrom)
	MustError(err)
	fmt.Println("seleted sdk : " + repo.GetName())

	RunCommand(argList[2], repo, data1, data2, rootsPath, sPlatfrom, sArch)

}

func helper() {
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
		// repo, _ = flutter.NewRepository(sPlatform)
	case "-n": // NodeJS
		repo, err = nodejs.NewRepository(sPlatform)
	case "-p": // Python
		// repo, _ = python.NewRepository(sPlatform)
	case "-r": // Ruby
		// repo, _ = python.NewRepository(sPlatform)
	case "-t": // Terraform
		// repo, _ = terraform.NewRepository(sPlatform)
	}

	if repo == nil {
		err = errors.New("not have platform \"" + sPlatform + "\"")
	}

	return repo, err
}

func RunCommand(sCommand string, repo repository.Repository, data1 string, data2 string, rootPath string, sPlatfrom string, sArch string) {
	switch sCommand {
	case "install":
		install(repo, data1, data2, rootPath, sPlatfrom, sArch)
	case "uninstall": // uninstall(detail)
	case "use": // use(detail,procarch)
	case "list": // list(detail)
	case "ls": // list(detail)
	case "on": // enable()
	case "off": // disable()
	case "root":
	//   if len(args) == 3 {
	// 	updateRootDir(args[2])
	//   } else {
	// 	fmt.Println("\nCurrent Root: "+env.root)
	//   }
	case "version":
		// fmt.Println(UVMVersion)
	case "v":
		// fmt.Println(UVMVersion)
	default:
		helper()
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
	sFileName = strings.ReplaceAll(sFileName, "{{version}}", "v"+sVersion)
	sFileName = strings.ReplaceAll(sFileName, "{{tag}}", sTag)
	sFileName = strings.ReplaceAll(sFileName, "{{os}}", sPlatfrom)
	sFileName = strings.ReplaceAll(sFileName, "{{arch}}", sArch)

	sPath := repo.GetPath()
	sPath = strings.ReplaceAll(sPath, "{{fileName}}", sFileName)
	sPath = strings.ReplaceAll(sPath, "{{version}}", "v"+sVersion)
	sPath = strings.ReplaceAll(sPath, "{{tag}}", "v"+sTag)
	sPath = strings.ReplaceAll(sPath, "{{type}}", repo.GetFileType())

	sUrl := repo.GetDist() + sPath

	return sUrl, sFileName
}

func install(repo repository.Repository, data1 string, data2 string, rootPath string, sPlatfrom string, sArch string) {
	sUrl, sFileName := setUrl(repo, data1, data2, sPlatfrom, sArch)

	sdkPath := filepath.Join(rootPath, repo.GetName())
	if !file.IsExist(sdkPath) {
		os.Mkdir(sdkPath, os.ModeDir)
	}

	sTempFile, sFolderVersion, sSDKPathVersion, err := download.Loading(repo, rootPath, sdkPath, sUrl, sFileName, data1, data2)
	if err != nil {
		printError(err)
		return
	}

	err = file.UnArchive("zip", sTempFile, sdkPath, repo.GetIsRenameFolder(), sFileName, sFolderVersion)
	if err != nil {
		fmt.Println("Unzip Error ", err)
		return
	}

	fmt.Println("sSDKPathVersion ; ", sSDKPathVersion)
	// create symlink

	switch sPlatfrom {
	case "windows":

		// symPath := filepath.Join("C:", "Program Files", repo.GetName())
		// err = os.Symlink(symPath, sSDKPathVersion)
		// if err != nil {
		// 	fmt.Println("Symlink Error ", err)
		// 	return
		// }

		fmt.Println("create symlink")
	}

	// remove Temp
	err = os.Remove(sTempFile)
	if err != nil {
		fmt.Println("Error Delete Temp File", err)
		return
	}
	fmt.Println("installed.")
	fmt.Println()
	fmt.Println("please run command:", "nvm", repo.GetCommand(), "use", data1, data2)

	// check env
	switch sPlatfrom {
	case "windows":
		symPath := filepath.Join("C:", "Program Files", repo.GetName())
		_, ok := os.LookupEnv(repo.GetEnv())
		if !ok {
			fmt.Println()
			fmt.Println("env: \"" + repo.GetEnv() + "\" not Found.")
			fmt.Println("please  set  env :", repo.GetEnv(), "=", symPath)
			fmt.Println("and add path env : %" + repo.GetEnv() + "%" + repo.GetEnvBin())
			fmt.Println()
		}
	}

	// //   // Remove symlink if it already exists
	// //   sym, _ := os.Stat(env.symlink)
	// //   if sym != nil {
	// // 	if !runElevated(fmt.Sprintf(`"%s" cmd /C rmdir "%s"`,
	// // 	  filepath.Join(env.root, "elevate.cmd"),
	// // 	  filepath.Clean(env.symlink))) {
	// // 	  return
	// // 	}
	// //   }

	// // // Create new symlink
	// // if !cmd.Command(fmt.Sprintf(`"%s" cmd /C mklink /D "%s" "%s"`,
	// // 	filepath.Join(root, "bin", "elevate.cmd"),
	// // 	filepath.Clean(symPath),
	// // 	path)) {
	// // 	return
	// // }

	// Delete Download

	// pathOS := "C:\\Program Files\\"
	// a, ok := os.LookupEnv("UVM_HOME_NODE")
	// fmt.Println("a : ", a, ok)
	// if !ok {
	// 	os.Setenv("UVM_HOME_NODE", filepath.Join(pathOS, "Node", "bin"))
	// }
	// a, _ = os.LookupEnv("UVM_HOME_NODE")
	// fmt.Println("b : ", a)

	// fmt.Println("ENV sPath : " + os.Getenv("Path"))

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
			helper()
			os.Exit(1)
		} else {
			printError("error :", err)
			helper()
			os.Exit(1)
		}
	}
}

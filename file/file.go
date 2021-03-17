package file

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return !os.IsNotExist(err)
	}
	return false
}

func UnArchive(ziptype string, src string, dest string, isRename bool, nameOld string, nameNew string) error {
	switch ziptype {
	case "zip":
		return UnZip(src, dest, isRename, nameOld, nameNew)
	case "tar":
		return UnTar(src, dest, isRename, nameOld, nameNew)
	case "7z":
		return Un7z(src, dest, isRename, nameOld, nameNew)
	}
	return errors.New("not match zip type")
}
func UnTar(src string, dest string, isRename bool, nameOld string, nameNew string) error {
	return nil
}

// Unzip func
// https://stackoverflow.com/questions/20357223/easy-way-to-unzip-file-with-golang
func UnZip(src string, dest string, isRename bool, nameOld string, nameNew string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		if isRename {
			f.Name = strings.Replace(f.Name, nameOld, nameNew, 1)
		}
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func Un7z(src string, dest string, isRename bool, nameOld string, nameNew string) error {
	// fmt.Println("src ", src)

	// a, err := lzmadec.NewArchive(src)
	// if err != nil {
	// 	fmt.Printf("lzmadec.NewArchive('%s') failed with '%s'\n", src, err)
	// 	os.Exit(1)
	// }

	// os.MkdirAll(dest, 0755)

	// fmt.Printf("opened archive '%s'\n", src)
	// fmt.Printf("Extracting %d entries\n", len(a.Entries))
	// for _, e := range a.Entries {
	// 	os.MkdirAll(dest+"\\"+e.Path, os.ModeDir)
	// 	err = a.ExtractToFile(dest+"\\"+e.Path, e.Path)
	// 	if err != nil {
	// 		fmt.Printf("a.ExtractToFile('%s') failed with '%s'\n", e.Path, err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Printf("Extracted '%s'\n", e.Path)
	// }

	return nil
}

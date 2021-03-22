package file

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/xi2/xz"
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
	case "targz":
		return UnTarGz(src, dest, isRename, nameOld, nameNew)
	case "tarxz":
		return UnTarXz(src, dest, isRename, nameOld, nameNew)
	case "7z":
		return Un7z(src, dest, isRename, nameOld, nameNew)
	}
	return errors.New("not match zip type")
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
			return fmt.Errorf("illegal file Path: %s", path)
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

func UnTarGz(src string, dest string, isRename bool, nameOld string, nameNew string) error {
	fi, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fi.Close()

	os.MkdirAll(dest, 0755)

	zipReader, err := gzip.NewReader(fi)
	if err != nil {
		log.Fatal("ExtractTarGz: NewReader failed")
	}
	defer zipReader.Close()

	err = extractTar(zipReader, dest, isRename, nameOld, nameNew)
	if err != nil {
		return err
	}

	return nil
}

func UnTarXz(src string, dest string, isRename bool, nameOld string, nameNew string) error {
	fi, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fi.Close()

	os.MkdirAll(dest, 0755)

	xzReader, err := xz.NewReader(fi, 0)
	if err != nil {
		return err
	}

	err = extractTar(xzReader, dest, isRename, nameOld, nameNew)
	if err != nil {
		return err
	}

	return nil
}

type symlinkPath struct {
	path   string
	target string
}

func extractTar(r io.Reader, dest string, isRename bool, nameOld string, nameNew string) error {
	tarReader := tar.NewReader(r)

	symlinkPathList := []symlinkPath{}
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return errors.New("ExtractTarGz: Next() failed : " + err.Error())
		}
		if isRename {
			header.Name = strings.Replace(header.Name, nameOld, nameNew, 1)
		}
		path := filepath.Join(dest, header.Name)
		// fmt.Println(header.Typeflag, " ::: ", path)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, 0755); err != nil {
				return errors.New("ExtractTarGz: Mkdir() failed : " + err.Error())
			}
		case tar.TypeReg:
			os.MkdirAll(filepath.Dir(path), 0755)

			outFile, err := os.Create(path)
			if err != nil {
				return errors.New("ExtractTarGz: Create() failed : " + err.Error())
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return errors.New("ExtractTarGz: Copy() failed : " + err.Error())
			}
			outFile.Close()
		case tar.TypeSymlink:
			name := header.Linkname
			target := filepath.Clean(filepath.Join(filepath.Dir(path), name))

			symlinkPathList = append(symlinkPathList, symlinkPath{path: path, target: target})
		default:
			return fmt.Errorf(
				"ExtractTarGz: uknown type: %b in %s",
				header.Typeflag,
				path)
		}
	}

	for _, symlink := range symlinkPathList {
		os.MkdirAll(filepath.Dir(symlink.path), 0755)

		err := os.Symlink(symlink.target, symlink.path)
		if err != nil {
			return fmt.Errorf("symlink error %v", err.Error())
		}
	}
	//
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

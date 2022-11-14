package util

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Zipper func(string) error

func ZipDirs(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	// Make a zip for all versions
	for _, f := range files {
		if f.IsDir() {
			versionRoot := filepath.Join(path, f.Name())
			file, err := os.Create(fmt.Sprintf("coinbase-wallet-chrome-%s.zip", strings.Split(f.Name(), "_")[0]))
			if err != nil {
				panic(err)
			}
			defer file.Close()

			w := zip.NewWriter(file)
			defer w.Close()

			walker := func(path string, info os.FileInfo, err error) error {
				fmt.Printf("Adding: %#v\n", path)
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				relPath, _ := filepath.Rel(versionRoot, path)
				if err != nil {
					return err
				}

				f, err := w.Create(relPath)
				if err != nil {
					return err
				}

				_, err = io.Copy(f, file)
				if err != nil {
					return err
				}

				return nil
			}

			err = filepath.Walk(versionRoot, walker)
			if err != nil {
				panic(err)
			}
		}
	}

	return nil
}

package utils

import (
	"archive/zip"
	"path/filepath"
	"os"
	"fmt"
	"io"
	"strings"
)

func ZipName(name string) string {
	return name + ".zip"
}

func Package(name, dir string) (string, error) {
	zipPath := filepath.Join(dir, ZipName(name))
	fzip, err := os.Create(zipPath)
	if err != nil {
		return "", err
	}
	defer fzip.Close()

	w := zip.NewWriter(fzip)
	defer w.Close()

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.Name() == ZipName(name) {
			return nil
		}
		if strings.HasPrefix(filepath.Base(info.Name()),".") {
			return nil
		}
		if !info.IsDir() {
			fDest, err := w.Create(path[len(dir)+1:])
			if err != nil {
				return err
			}
			fSrc, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fSrc.Close()

			fmt.Printf("Compressing %s, size %d\n", path, info.Size())
			if _, err = io.Copy(fDest, fSrc); err != nil {
				return err
			}
		}
		return nil
	})
	return zipPath, err
}

func Unpack(f string, dir string) error {
	r, err := zip.OpenReader(f)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name != "README.md" && f.Name != "cmd.sh" {
			fmt.Printf("Unpacking %s \n ", f.Name)
			rc, err := f.Open()
			if err != nil {
				return err
			}

			dstPath := filepath.Join(dir, f.Name)
			w, err := os.Create(dstPath)
			if err != nil {
				rc.Close()
				return err
			}

			if _, err := io.Copy(w, rc); err != nil {
				rc.Close()
				w.Close()
				return err
			}

			rc.Close()
			w.Close()

		}
	}
	return nil
}

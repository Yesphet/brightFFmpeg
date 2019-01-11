package utils

import (
	"syscall"
	"os"
	"io"
	"io/ioutil"
)

const commonPerm = 0777

func Makedir(path string) error {
	mask := syscall.Umask(0)
	defer syscall.Umask(mask)
	err := os.MkdirAll(path, commonPerm)
	return err
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil || os.IsExist(err) {
		return true
	}
	return false
}

func DeleteFile(path string) error {
	return os.Remove(path)
}

func WriteToFile(path string, reader io.Reader) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	return err
}

func Read(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ret, err := ioutil.ReadAll(f)
	return ret, err
}

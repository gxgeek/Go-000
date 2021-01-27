package file

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path"
)

//GetSize：获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)
	return len(content), err
}

//GetExt：获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}




func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

func IsNotExistMkDir(src string) error {
	if CheckNotExist(src) {
		if err := MkDir(src);err != nil {
			return err
		}
	}
	return nil
}
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return errors.Wrap(err,"os MkDir error")
	}
	return nil
}


func OpenLogFilePath(filePath string,flag int, perm os.FileMode) (*os.File, error) {
	open, err := os.OpenFile(filePath, flag, perm)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	return open,nil
}



// Open a file according to a specific mode
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}


// MustOpen maximize trying to open the file
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrapf(err,"os.Getwd err: %v")
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, errors.Wrapf(err,"file.IsNotExistMkDir src: %s", src)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, errors.Wrapf(err,"Fail to OpenFile")
	}

	return f, nil
}

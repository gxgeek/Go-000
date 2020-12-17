package file

import (
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


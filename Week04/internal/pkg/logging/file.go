package logging

import (
	"Go-000/Week04/internal/pkg/file"
	"Go-000/Week04/internal/pkg/setting"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"time"
)

var (
	//LogSavePath = "runtime/logs"
	//LogSaveName = "log"
	//LogFileExt  = "log"
	//TimeFormat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s", setting.AppSetting.LogSaveName,time.Now().Format(setting.AppSetting.TimeFormat), setting.AppSetting.LogFileExt)
}


func OpenLogFilePath(filePath, fileName string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err,"os.Getwd err")
	}
	src := dir + "/" + filePath
	if file.CheckPermission(src) {
		return nil, errors.Wrapf(err,"file.CheckPermission Permission denied src %s",src)
	}
	if err = file.IsNotExistMkDir(src); err!=nil {
		return nil, errors.Wrapf(err,"file.IsNotExistMkDir src: %s,",src)
	}
	open, err := file.OpenLogFilePath(src + fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, errors.Wrapf(err,"Fail to OpenFile : %s,",src)
	}
	return open,nil
}

func mkDir() error {
	dir, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err,"os Getwd error")
	}
	path := dir + "/" + getLogFilePath()
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "create file: %s fail", path)
	}
	return nil
}

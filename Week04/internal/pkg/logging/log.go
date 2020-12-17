package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int
var (
	F *os.File
	DefaultPrefix = ""
	DefaultCallerDepth = 2
	logger *log.Logger
	logPrefix = ""
	levelFlags = []string{"DEBUG","INFO","WARN","ERROR","FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Setup()  {
	F, err := OpenLogFilePath(getLogFilePath(),getLogFileName())
	if err != nil {
		log.Fatalf("log init error %+v",err)
	}
	//log.New：创建一个新的日志记录器。out定义要写入日志数据的IO句柄。
	//prefix定义每个生成的日志行的开头。flag定义了日志记录属性
	//log.LstdFlags：日志记录的格式属性之一
	writer := io.MultiWriter(os.Stderr, F)
	logger = log.New(writer,DefaultPrefix,log.LstdFlags)
}


func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}
func InfoF(format string, v ...interface{}) {
	setPrefix(INFO)
	logger.Printf(format, v... )
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}

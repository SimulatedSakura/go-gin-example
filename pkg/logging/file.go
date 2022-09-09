package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20010502"
)

// 获取文件路径
func getLogFilePath() string {
	return LogSavePath
}

// 获取完全路径
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// 打开日志路径
func openLogFile(filePath string) *os.File {
	//返回文件信息结构描述文件。如果出现错误，会返回*PathError
	_, err := os.Stat(filePath)
	switch {
	//检查文件是否存在
	case os.IsNotExist(err):
		mkDir()
	//检查文件权限是否满足
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}
	//调用文件，首个参数为文件路径，接下来是打开方式(常量选择，如下方的含义是写入数据时追加到文件尾部、不存在则新建文件、以只写方式打开)和文件组权限配置
	//https://mileslin.github.io/2020/09/Golang/os-OpenFile-的第三個參數-perm-FileMode/
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

func mkDir() {
	//返回与当前目录对应的目录名
	dir, _ := os.Getwd()
	// 创建对应的目录和所需的子目录，若成功返回nil否则返回err
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

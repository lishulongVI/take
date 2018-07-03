package utils

import (
	"os"

	"github.com/take/log"
)

var logger = log.NewLogger(os.Stdout)

type myfile struct {
}

var File = myfile{}

// 文件大小
func (*myfile) GetFileSize(path string) int64 {

	logger.Info("query path:\t" + path)
	fi, err := os.Stat(path)

	if nil != err {
		logger.Error(err)
		return -1
	}
	return fi.Size()
}

//文件是否存在
func (*myfile) IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

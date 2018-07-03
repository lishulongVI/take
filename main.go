package main

import (
	"fmt"

	"github.com/take/utils"
)

func main() {
	// a := &utils.File
	PATH := "/Users/lishulong/go/src/github.com/take/main.go"
	res := utils.File.GetFileSize(PATH)
	b := utils.File.IsExist(PATH)
	fmt.Println(b)
	fmt.Println(res)

	/**
	I 2018/07/03 14:52:03 files.go:19: query path:  /Users/lishulong/go/src/github.com/take/main.go
	true
	257
	***/
}

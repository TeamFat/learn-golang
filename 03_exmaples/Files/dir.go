/**
 * datetime: 2017-06-09
 * Go 目录操作
 */
package main

import (
	"fmt"
	"os"
)

func main() {

	//创建目录
	err := os.Mkdir("tmp_dir", 0755)
	if err != nil {
		return
	}
	//进入目录
	os.Chdir("tmp_dir")
	//进入目录创建文件
	os.Create("tmp.txt")

	//打开目录
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()
	//删除目录
	os.Remove("tmp_dir")
	fileInfos, err := dir.Readdir(-1)

	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Printf("FileName:%s\n", fi.Name())
	}
}

//$ go run readdir.go

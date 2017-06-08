/**
 * datetime: 2017-06-07
 * 文件操作
 */
package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {

	//创建文件
	f, _ := os.Create("log.log")
	defer f.Close()
	//文件句柄类型
	fmt.Println(reflect.ValueOf(f).Type()) //*os.File
	//打开文件
	f2, _ := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	// f2, _:=os.Open("log.log")
	defer f2.Close()
	fmt.Println(f2.Stat())
	//文件信息
	f2Stat, _ := f2.Stat()
	//获取信息
	fmt.Println("文件大小：", f2Stat.Size())
	fmt.Println("文件名称：", f2Stat.Name())
	//文件最后修改时间
	fmt.Println("文件最后修改时间：", f2Stat.ModTime())
	//判断文件是否存在
	var notExistsFilename = "not_exists.file"
	if _, err := os.Stat(notExistsFilename); os.IsNotExist(err) {
		fmt.Println(notExistsFilename + "文件不存在")
	}

}

//$ go run file.go
// *os.File
// &{log.log 0 420 {63632452529 0 0x110c4e0} {16777221 33188 1 670686 501 80 0 [0 0 0 0] {1496855162 0} {1496855729 0} {1496855729 0} {1496855162 0} 0 0 4096 0 0 0 [0 0]}} <nil>
// 文件大小： 0
// 文件名称： log.log
// 文件最后修改时间： 2017-06-07 11:15:29 +0800 CST
// not_exists.file文件不存在

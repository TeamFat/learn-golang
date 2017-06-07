/*
 * datetime: 2017-06-08
 * 文件操作：三种方式获取文件md5值
 */
package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	//创建一个文件
	filename := "testfile.md5"
	var file *os.File
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		//文件不存在-创建
		file, _ = os.Create(filename)
	} else {
		//文件存在-打开
		//openFile 可以设置读写权限，open函数只能是只读权限
		file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	}
	defer file.Close()
	//文件写入内容
	//也可以使用file.WriteString("这是写入的字符串")
	// _, err := io.WriteString(file, "这是写入的字符串")
	var content string
	content = "这是写入文件的内容-jsser"
	err := ioutil.WriteFile(filename, []byte(content), 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取文件md5 方式一
	c, err := ioutil.ReadAll(file)
	if err == nil {
		fmt.Printf("方式一：md5 = %x\n", md5.Sum(c))
	}

	//获取文件md5 方式2
	md5Hash := md5.New()
	_, err1 := io.Copy(md5Hash, file)
	if err1 == nil {
		fmt.Printf("方式二：md5 = %x\n", md5Hash.Sum(nil))
	}

	//获取文件md5 方式3
	bf := bufio.NewReader(file) //增加bufio 缓存读取
	h := md5.New()
	_, err3 := io.Copy(h, bf)
	if err3 == nil {
		fmt.Printf("方式三：md5 = %x\n", h.Sum(nil))
	}
}

//$ go run file_md5.go
// 方式一：md5 = d41d8cd98f00b204e9800998ecf8427e
// 方式二：md5 = d41d8cd98f00b204e9800998ecf8427e
// 方式三：md5 = d41d8cd98f00b204e9800998ecf8427e

//总结：优先选择bufio方式计算Md5,尽量避免使用readall计算文件md5值（大文件会有内存分配问题）

/**
 * datetime: 2017-05-28
 * Go错误处理2
 * 1. Go中提供 panic 函数来报告程序在运行过程中遇到的致命错误。Panic 可以传入任意类型参数，一般传递string或者error对象，用来方便查看异常信息
 * 2. Go中提供 recover 函数来捕捉运行时产生的panic报告的错误信息, 防止程序崩溃。recover一般会和defer一起使用，放在函数的开头, 可以有效的该函数及其下层调用中代码引起的恐慌。
 * 3. 编译器运行时系统报出的错误，recover不能捕捉到。(静态语言)
 */
package main

import (
	"errors"
	"fmt"
)

func main() {

	defer func() {
		if p := recover(); p != nil {
			//执行到此，程序没有崩溃, 捕捉到了该错误，这里可以做错误类型判断，来处理对应的错误逻辑
			fmt.Printf("Recover a panic : %s \n", p)
			fmt.Println("程序会执行到这里") //如果捕捉到panic 那么这里会执行下去
		}
	}()

	panic(errors.New("this is fatal error"))
	fmt.Println("这里还会执行吗？") //这里不会执行
	arr1 := [3]int{1, 2, 3}
	_ = arr1
	//fmt.Printf("arr[3] = %d\n", arr1[3]) //数组越界访问错误，是由Go 运行时系统报告的。
}

// $ go run panic_recovier.go

// Recover a panic : this is fatal error
// 程序会执行到这里

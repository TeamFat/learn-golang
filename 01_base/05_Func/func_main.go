/**
 *  datetime: 2017-05-23
 *  Go语言 - 函数
 *  0. 函数的零值是nil
 *  1. Go中函数声明包含：函数名称，参数列表, 参数名称，返回值名称，返回值类型。
 *  2. Go函数可以有多个返回结果。在函数有多个返回结果情况下，要么全部省略返回值名称，要么全部加上返回值名称
 *  3. 如果函数结果有名称，那么函数在被调用时，他们的变量就会被隐显声明。
 *  4. 在Go中习惯性的将错误类型作为函数结果列表中的最后一员。errors.New("异常")
 *  5. 函数在Go中是一等类型，函数可以当做其他函数的参数，也可以当做结果。
 *  6. Go中的函数闭包：内层函数引用了外层函数环境的变量，其返回值也是一个函数。
 */
package main

import (
	"errors"
	"fmt"
)

type myint int

/**
 *  变量函数
 *  调用方法：sumFunc(1,2)
 */
var sumFunc = func(i, j int) int {
	k := i + j
	return k
}

func main() {

	fmt.Println("")
	var i myint
	i = 10
	fmt.Println(f1(i))

	/**
	 *  bop作为operat的闭包函数
	 */
	bop := func(op1, op2 int) (result int, err error) {
		if op2 == 0 {
			err = errors.New("divison by zero")
		}
		result = op1 / op2
		return
	}

	result, err := operate(20, 20, bop)
	if err == nil {
		fmt.Println("result = ", result)
	}

	/**
	 * 调用变量函数
	 */
	fmt.Println("sumFunc(10, 20) = ", sumFunc(10, 20))

	//可变参数函数
	total1 := sumNumbers(1, 2, 3, 4, 5)
	fmt.Printf("1 + 2 + 3 + 4 + 5  = %d\n", total1)

}

/**
 *   声明一个函数
 *   函数名：f1,  参数名称：i, 参数类型：myint,  返回结果类型int
 */
func f1(i myint) int {
	i++
	r := +int(i + 1)
	return r
}

/**
 *  函数结果又名称t1, 结果参数t1会被隐式声明。所以在调用f2(1,2,3)的时候，函数返回结果为：0
 */
func f2(i, j, k int) (t1 int) {
	return
}

/**
 *   1. 函数中有异常信息一般作为函数结果列表中最后一个参数抛出。
 */
func divide(i, j int) (result int, err error) {

	if j == 0 {
		err = errors.New("divison by zero")
		return
	}
	result = i / j
	return
}

/**
 *  实现binaryOperation 闭包
 *  在main函数实现了调用
 */
type binaryOperation func(op1 int, op2 int) (result int, err error)

func operate(op1 int, op2 int, bop binaryOperation) (result int, err error) {
	if bop == nil {
		err = errors.New("invalid binary operation function")
		return
	}
	//闭包
	return bop(op1, op2)
}

//可变参数，最多只能有一个可变参数，可变参数只能放到函数参数的末尾
// sumNumbers(1,2,3,4,5)
func sumNumbers(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

//$ go run func_main.go

// 2
// result =  1
// sumFunc(10, 20) =  30

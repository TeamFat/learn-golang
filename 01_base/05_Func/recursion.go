/**
 *  datetime: 2017-05-23
 *  Go语言 - 递归
 *  Go中函数递归深度没有做限制
 */
package main

import (
	"fmt"
)

//1 * 2 * 3 * 4 * 5
func fact(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println(n)
	return n * fact(n-1)
}

func main() {

	fmt.Println(fact(5))
	// fmt.Println(fact(100))
}

//$go run recursion.go

// 5
// 4
// 3
// 2
// 1
// 120

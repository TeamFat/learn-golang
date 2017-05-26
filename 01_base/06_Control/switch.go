/**
 *  datetime: 2017-05-26
 *  Go语言 - switch
 *  分支表达式可以支持任意类型，并且可以省去break, 需要需要继续下一分支，可以使用fallthrough （但不在判断条件）
 */

package main

import (
	"fmt"
	"time"
)

func main() {

	i := 10
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("three")
	case 10:
		fmt.Println("ten")
		fallthrough
	default:
		fmt.Println("nothing")
	}

	// ten, nothing

	//case 分支  支持多个值

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// case 分支 可以是表达式。可代替if else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's Before noon")
	default:
		fmt.Println("It's After noon")
	}
}

//$ go run switch.go
// ten
// nothing
// It's the weekend
// It's Before noon

/**
 * datetime: 2017-05-28
 * Go错误处理
 *
 * 1. errors.New，fmt.Errorf 创建一个可描述性的错误信息，实现了error接口的错误对象
 * 2. 通过判断错误对象实例来确定具体错误类型
 */
package main

import (
	"errors"
	"fmt"
)

func main() {
	//当被除数为0时，会抛出err错误
	d1, err := divide(1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d1)
	}

	//fmt 标准库中 Errorf  创建一个可描述性的错误信息
	const name, id = "jsser", 100
	err2 := fmt.Errorf("user %q (id %d) not found", name, id)
	if err2 != nil {
		fmt.Println(err2)
	}

}

//函数通常在最后的返回值中返回错误信息
//定义一个除法函数，并且判断了被除数为0时候，抛出错误"被除数不能为0"
func divide(i1, i2 float64) (res float64, err error) {
	if i2 == 0 {
		err = errors.New("被除数不能为0")
	}
	res = float64(i1 / i2)

	return
}

// go run errors.go

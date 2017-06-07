/**
 *
 * datetime: 2017-06-07
 *
 */
package main

import (
	"fmt"
	"os"
)

func main() {

	os.Setenv("MYSQL.PASSWD", "1")
	fmt.Println("MYSQL.PASSWD = ", os.Getenv("MYSQL.PASSWD"))

	/** 遍历环境变量
	 */
	for _, e := range os.Environ() {
		fmt.Println(e)
		// pair := strings.Split(e, "=")
		// fmt.Println(pair[0])
	}
}

//$ go run env.go

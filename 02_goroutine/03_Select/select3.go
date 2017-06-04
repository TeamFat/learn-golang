/**
 * datetime: 2017-06-04
 * select和channel 模拟超时机制
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println("res = ", res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println("res = ", res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

}

//$go run select2.go
// timeout 1
// res =  result 2

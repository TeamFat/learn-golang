package main

import (
	"fmt"
)

var chInt = make(chan int, 3)

func main() {

	go func() {
		for i := 0; i < 10; i++ {
			chInt <- i
		}

	}()
	for {
		if i, err := <-chInt; err {
			fmt.Println(i)
		}
	}

}

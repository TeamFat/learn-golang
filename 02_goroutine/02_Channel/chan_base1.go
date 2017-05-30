package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main() {

	//初始化2个通道, synChan1, synChan2
	synChan1 := make(chan struct{}, 1)
	synChan2 := make(chan struct{}, 2)

	go func() {
		//演示发送操作
		for _, elem := range []string{"a", "b", "c"} {
			strChan <- elem
			fmt.Println("Sent:", elem, "[sender]")
			if elem == "c" {
				synChan1 <- struct{}{}
				fmt.Println("Sent a sync signal . [sender]")
			}
		}
		fmt.Println("wait 2 senconds ... [sender]")
		time.Sleep(time.Second * 2)
		close(strChan)
		synChan2 <- struct{}{}
	}()

	go func() {
		//演示接收操作
		<-synChan1
		fmt.Println("Received a sync signal and wait a second...[reveiver]")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Received: ", elem, "[reveiver]")
			} else {
				break
			}
		}

		fmt.Println("Stopped . [reveiver]")
		synChan2 <- struct{}{}
	}()

	<-synChan2
	<-synChan2

}

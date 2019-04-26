package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func fn1(ch chan string) {
	for i := 0; i < 19; i++ {
		ch <- "I' am sample number:" + strconv.Itoa(i)
		time.Sleep(1 * time.Second)
	}
}

func fn2(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(10 * time.Second)
	}
}

func main() {
	ch1 := make(chan string, 3)
	ch2 := make(chan int, 5)

	for i := 0; i < 10; i++ {
		go fn1(ch1)
		go fn2(ch2)
	}
	// select 消费多个chan的形式
	for { // select 要搭配for循环使用
		select {
		case str, ok := <-ch1:
			if !ok {
				log.Printf("ch1 consume failed")
			}
			fmt.Println(str)
		case num, ok := <-ch2:
			if !ok {
				log.Printf("ch2 consume failed")
			}
			fmt.Println(num)

		}
	}
}

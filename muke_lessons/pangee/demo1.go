package main

import (
	"fmt"
	"time"
)

/*
	chan 类型遵循FIFO的原则，注意都有哪些地方消费了chan
*/

func sample(m chan string) {
	m <- "hello goroutine1"
	m <- "hello goroutine2"
	m <- "hello goroutine3"
	m <- "hello goroutine4"
}

func sample2(m chan string) {
	time.Sleep(2 * time.Second)
	str := <-m
	str = str + " I' am goroutine!"
	m <- str
	close(m) // 这里明确知道了关闭chan的时机，那要是不知道呢？应该在哪里关闭？
}

func main() {
	var message = make(chan string, 3)
	go sample(message)
	go sample2(message)

	time.Sleep(3 * time.Second)
	/*
		利用for range 消费chan，即使chan中没有了数据也会等待数据写入在消费，除非chan为nil，range会退出
		所以应当在合适的时机关闭chan
	*/
	for str := range message {
		fmt.Println(str)
	}
	fmt.Println("hello world")

}

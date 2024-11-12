/*
* Channel 3 Range, Close and Select
 */
package main

import (
	"fmt"
)

func main() {
	data := make(chan int)
	quit := make(chan int)
	/* 	c := make(chan int)

	   	go func() {
	   		defer fmt.Println("Sub gorountie close")
	   		for i := 0; i < 3; i++ {
	   			c <- i
	   			fmt.Println("Send Data : ", i)
	   		}

	   		close(c) //<---用 close(chan) 关闭一个channel，关闭后无法再向channel发送数据，但可以从chan中收到值
	   	}()

	   	time.Sleep(2 * time.Second) // 确保子goroutine先执行，打印结果更直观

	   	// channel 类型可以直接用range读队列的数据

	   	for data := range c {
	   		fmt.Println("Read data: ", data)
	   	} */

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Read Data : ", <-data) // 从channel中读数据
		}
		quit <- 0
	}()

	Fibonacii(data, quit)
	defer fmt.Println("Main goroutine finished")
}

/*
* Select 可以看作是一种IO多路复用模型
 */

func Fibonacii(data, quit chan int) {
	defer fmt.Println("Fibonacii goroutine finished")
	x, y := 1, 1
	for {
		select {
		case data <- x: // 向channel中写数据，如果这个channel有人读，就会执行这个case
			x, y = y, x+y
		case <-quit: // 从channel中读数据，如果这个channel有人写，就会执行这个case
			fmt.Println("Quit")
			return
		}
	}
}

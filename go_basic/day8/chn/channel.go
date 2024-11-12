package main

import (
	"fmt"
)

func main(){
	// 创建一个channel, 用于goroutine之间的通信
	c := make(chan int)

	go func(){
		defer fmt.Println("goroutine结束")
		fmt.Println("goroutine is running...")

		c <- 888 // 将数据通过channel发送出去
	}()

	num := <- c // 从channel中接收数据
	fmt.Println("num = ", num)
	fmt.Println("main goroutine结束")
}

/*
可以看到main func的goroutine永远是在Sub func的goroutine结束之后才结束的，
这是因为main func的goroutine会等待channel中的数据被接收，如果这个channel没设置缓冲区就会被阻塞，达到了一种同步的效果
 */
package main

import (
	"fmt"
	"time"
)

/*
we have a channel with a buffer of 3. This means that the channel can hold up to 3 values before it blocks.
The main goroutine will wait for the channel to be filled with 3 values before it can continue.
*/

func main() {
	// c := make(chan int, 3)

	// go func() {
	// 	defer fmt.Println("Sub goroutine Closed")
	// 	fmt.Println("goroutine is running...")
	// 	for i := 0; i < 3; i++ {
	// 		fmt.Println("Sub goroutine is sending -> ", i)
	// 		c <- i
	// 	}
	// }()

	// for i := 0; i < 3; i++ {
	// 	num := <-c
	// 	fmt.Println("Main goroutine is receiving <- ", num)
	// }
	// fmt.Println("main goroutine closed")
	submain()
}

/*

the following code will cause a deadlock because the main goroutine is waiting for 4 values to be sent to the channel, but the channel can only hold 3 values.

It will happen because the main goroutine is waiting for 4 values to be sent to the channel, but the channel can only hold 3 values. The main goroutine will be blocked waiting for the 4th value to be sent to the channel, but the sub goroutine has already finished sending 3 values to the channel and has closed the channel. This will cause a deadlock.



func submain() {
	c := make(chan int, 3)

	go func() {
		defer fmt.Println("Sub goroutine Closed")
		fmt.Println("goroutine is running...")
		for i := 0; i < 3; i++ {
			fmt.Println("Sub goroutine is sending -> ", i)
			c <- i
		}
	}()

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("Main goroutine is receiving <- ", num)
	}
	fmt.Println("func closed")

}
*/

func submain() {
	c := make(chan int, 3)

	go func() {
		defer fmt.Println("Sub goroutine Closed")
		fmt.Println("goroutine is running...")
		for i := 0; i < 4; i++ {
			/* 要传输的数据量超过了channel的缓冲区大小，达到最大容量时会阻塞 */
			fmt.Printf("Sub goroutine is sending -> %d\n", i)
			c <- i
			fmt.Printf("Sub goroutine sent -> %d\n", i)
			fmt.Printf("len of c: %d, cap of c: %d\n", len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Printf("Main goroutine is receiving <- %d\n", num)
	}
	fmt.Println("func closed")
}

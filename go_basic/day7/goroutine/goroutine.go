package main

import (
	"fmt"
	"time"
)

func Task() {
	i := 0
	for {
		fmt.Println("Second Task i = ", i)
		i++
		time.Sleep(1 * time.Second)
		if i > 3 {
			return
		}
	}
}

func main() {
	/*
	* 使用go Func() 去创建goroutine
	 */
	go Task()

	i := 0
	for {
		fmt.Println("Main Task i = ", i)
		i++
		time.Sleep(1 * time.Second)
		if i > 3 {
			break
		}
	}
}

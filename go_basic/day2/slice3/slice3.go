package main

import (
	"fmt"
)

func EvilCutSlice(originSlice []int) {
	fmt.Println("I WILL CHANGE YOUR SLICE BE ZERO FROM 2 to 5")
	for i := 1; i < 5; i++ {
		originSlice[i] = 0
	}
}

func main() {
	myTinySlice := make([]int, 3)                  // 开辟一个len为3的slice，初始值都是0
	fmt.Println("My Tiny Slice is: ", myTinySlice) // 0 0 0

	var myEmptySlice []int
	// 判断一个slice是否0
	if myEmptySlice == nil {
		fmt.Println("This is an empty slice")
	} else {
		fmt.Println("This is an non-empty slice ^_^")
	}

	myTinySlice = append(myTinySlice, 1, 2, 3) // 给slice追加三个元素，0 0 0 1 2 3
	fmt.Println(myTinySlice)

	// LET"S MAKE SLICE GREAT AGAIN
	for i := 0; i < len(myTinySlice); i++ {
		if i >= len(myTinySlice) {
			myTinySlice = append(myTinySlice, i+1)
			continue
		}
		myTinySlice[i] = i + 1

	}

	// Let's see our GREATER SLICE
	for index, value := range myTinySlice {
		fmt.Println("Index = ", index, " Value = ", value)
	}

	// ...
	EvilCutSlice(myTinySlice)

	fmt.Println("After Evil method, my slice is :")
	for index, value := range myTinySlice {
		fmt.Println("Index = ", index, " Value = ", value)
	}
}

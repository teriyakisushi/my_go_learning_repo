package main

import (
	"fmt"
)

func ModifyDoubleSlice(originSlice [4]int) {
	// PASS BY VALUE, A COPIED OF ORIGIN SLICE

	fmt.Println("Hahaha, Your kawaii slice now are changed by my evil method! ")
	for i := 0; i < len(originSlice); i++ {
		originSlice[i] *= 2
	}

	for i := 0; i < len(originSlice); i++ {
		fmt.Println("Index = ", i, " Value = ", originSlice[i])
		// 2,4,6,8
	}
}

func RealEvilSlice(originSlice []int) {
	// REAL REFERENCE TRANSMISSION
	fmt.Println("This is Really Evil Method, Your Slice will destory  ME HAHAHA!")
	for i := 0; i < len(originSlice); i++ {
		originSlice[i] = 0
	}
}

func main() {
	myTinySlice := [4]int{1, 2, 3, 4}
	// range keyword:

	fmt.Println("FOREACH MY TINY SLICE: ")
	for index, value := range myTinySlice {
		fmt.Println("Index = ", index, " Value = ", value)
	}

	ModifyDoubleSlice(myTinySlice)

	// It seems my evil method had changed my tiny slice, let's check
	fmt.Println("After Modified, My tiny slice now is : ")
	for index, value := range myTinySlice {
		fmt.Println("Index = ", index, " Value = ", value)
		// 1, 2, 3, 4
	}

	fmt.Println("Haha, It looks like he lie to me, my tiny slice no changed anything")
	RealEvilSlice(myTinySlice[:])

	// Let's check again
	fmt.Println("After Evil modified, my tiny slice now is: ")
	for index, value := range myTinySlice {
		fmt.Println("Inex = ", index, " Value = ", value)
		// 0 0 0 0
	}

	// en...如果定义了数组的大小后，传递给另一个func其实是值传递，不改变原始数组
	// 如果是slice的话，就是引用传递了，没写全，正确的给slice的引用传递应该是
	// mySlice := []int{1,2,3,4}
	// EvilMethod (mySlice)
}

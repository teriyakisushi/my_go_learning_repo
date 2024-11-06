package main

import (
	"fmt"
)

func main() {
	/*
	* We could use assign statements := to create a slice
	* The following statements is to declare and assign a len of 6 slice,
	 */

	mySpecialSlice := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("My Special Slice is: ", mySpecialSlice)
	fmt.Println("Its len is: ", len(mySpecialSlice))

	// Let's append something in mySpecialSlice

	mySpecialSlice = append(mySpecialSlice, 100, 200, 300)

	fmt.Println("My New Special Slice is: ", mySpecialSlice)

	/*
	* Or use keyword make,
	* the following statements is to declare & assign a non-specified slice
	* We can append value for it anytime
	 */

	myTindySlice := make([]int, 10)
	fmt.Println("At first, the TinySlice is :", myTindySlice)
	for i := 0; i < len(myTindySlice); i++ {
		myTindySlice[i] = i + 1
	}
	myTindySlice = append(myTindySlice, 11, 12, 13, 14)
	fmt.Println("Now the TindySlice is: ", myTindySlice)
}

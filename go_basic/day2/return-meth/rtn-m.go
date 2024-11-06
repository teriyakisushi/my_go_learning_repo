package main

import (
	"fmt"
)

func SortTwoNums(a int, b int) (BiggerNum int, SmallerNum int) {
	Num1 := a
	Num2 := b
	if Num1 > Num2 {
		return Num1, Num2
	}

	// We can return two values in Go

	return Num2, Num1
}

// We can also return multi values by using this way

func MakeItDouble(Num int) (origin int, double int) {
	num := Num
	origin = num     /* <------- We defined it in func name (origin int, double int)*/
	double = num * 2 /* if the type of return values are both same, we can write like this:*/
	/* origin, doubleNum int*/
	return
}

func main() {
	var (
		inputNum1 int
		inputNum2 int
		originNum int
	)

	fmt.Println("Input two Nums to sort them from Bigger -> Smaller")
	_, err := fmt.Scanf("%d %d", &inputNum1, &inputNum2)
	if err != nil {
		fmt.Println("You input is not a Integer!")
	}

	// Test default func
	fmt.Println(SortTwoNums(inputNum1, inputNum2))

	// Test SECOND FUNC
	fmt.Println("Input a Num , make it double")
	_, err = fmt.Scanf("%d", &originNum)
	if err != nil {
		fmt.Println("Your input is not a Integer!")
	}
	fmt.Println(MakeItDouble(originNum))
}

package main

/*
多个包的导入可以用import() 语句
Go的检查器比较严格，导入的包如果unused会报错
*/

import (
	"fmt"
)

/*
要执行的程序必须要有main入口，声明函数用 func
*/

func main() {
	/*
		* Go的声明&&赋值变量语句有三种，
		1. var Variable Type，  声明类型
		2. Variable := Value,  声明变量并赋值
		3. Varaible = Value , 对变量赋值
	*/
	var myVar int // 声明 myVar 为 Integer 类型但没有赋值
	myVar = 200   // 对声明的myVar赋值
	myVar2 := 100 // 声明 myVar2 的同时为它赋值100，Compiler会推断出它是Integer类型

	fmt.Println("Value of myVar is: ", myVar)                 // Println是打印一行信息
	fmt.Printf("Using Formatter print myVar is :%d\n", myVar) // Printf和C的printf类似

	// Print myVar2
	fmt.Println("Value of myVar2 is: ", myVar2)

	/* 除了Integer类型，Go能声明的类型还有：
	string, float32, float64, bool ....
	*/

	myString := "Hello world"
	myString2 := []string{"h", "e", "l", "l", "o"} // 这是一个Slice，可以随时append

	// Let's See
	fmt.Println(myString)
	fmt.Println(myString2)

	myString2 = append(myString2, "W", "W", "W")
	fmt.Println(myString2)
}

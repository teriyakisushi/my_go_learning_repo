package main

import (
	"fmt"
)

type Pet struct {
	Name string
}

// 万能类型的interface（空接口），有点像java的generic type
func SecretPetFunc(args interface{}) {
	// fmt.Printf("%s calling me!\n", args)
	// 可以预见的是如果传进来的是别的类型，上述打印出来的东西是非预期的
	// Go有断言处理，非常的convenient呀~

	value, ok := args.(string)
	if !ok {
		fmt.Println("No pet calling me")
	} else {
		fmt.Printf("%s calling me!\n", value)
	}
}

func main() {
	MyPet := Pet{"Qiqi"}
	SecretPetFunc(MyPet.Name)
	SecretPetFunc(1145.14)
}

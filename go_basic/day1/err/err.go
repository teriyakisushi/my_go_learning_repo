package main

import (
	"fmt"
)

func SayGoodbye() {
	fmt.Println("The Programme was exit with good bye~")
}

func ReturnSthInteresting() (err error) {
	fmt.Println("This is test your error happend!")
	var InputValue int
	_, err = fmt.Scanln(&InputValue)
	if err != nil {
		fmt.Println("Your input is not a Integer!")
		return err
	}

	return nil
}

func main() {
	var a int
	var err error
	fmt.Println("Please Input a integer: ")
	_, err = fmt.Scanln(&a)
	/*
	*  Error Deal, we usually use err to receive error that  func return
	 */ if err != nil {
		fmt.Println("Your input is not a integer!")
		return
	}

	/*
	*  defer语句用于在该函数体执行完后再执行的语句
	*  如果有 return 和 defer 同时存在，那会先执行return
	*  有多个defer的话，按出栈顺序执行defer语句
	*  有点像 finally
	*  不过本函数体的err触发时不会执行defer，因为没有确切返回的东西
	 */
	err = ReturnSthInteresting()
	defer SayGoodbye()
}

package main

import "fmt"

/*
* @ class的封装
*
type Human struct {
	Name string
	Kind string
	Age  int
}

func (this *Human) PrintName() {
	fmt.Println("The Human name is :", this.Name)
}

func main() {

		SuperMan := Human{
			"SuperMan",
			"Hero",
			22,
		}
		SuperMan.PrintName()

}
*/

/*
*  @class的继承
*
 */

type Human struct {
	Name string
	Kind string
	Age  int
}

type SuperHuman struct {
	Human // <----没错...打这个相当于继承了父类...绝了
	Level int
}

// 父类方法
func (this *Human) HumanGo() {
	fmt.Println("AM WALKING")
}

// 子类方法
func (this *SuperHuman) SuperHumanGo() {
	fmt.Println("SuperMan are walking!")
}

func main() {
	AHuman := Human{
		"SuperMan",
		"Hero",
		22,
	}

	// 定义一个子类，需要这样子...
	Superman := SuperHuman{
		AHuman,
		999,
	}

	// 父类方法
	Superman.HumanGo()

	// 子类的方法
	Superman.SuperHumanGo()
}

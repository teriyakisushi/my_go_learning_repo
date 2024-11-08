package main

import (
	"fmt"
)

type myint int // just like typedef long long ll in cpp

type Pets struct {
	Name string
	Kind string
	Age  int
}

/*
* 如果不加 * 的话就只是值传递，即对象的拷贝
* 加*表示其指针，是Reference传递，和CPP一样
*
 */

/*
* 在funcName前这么写可以指定是某个struct的方法
* Attention: funcName 首字母大写代表是public method，小写则是private
 */
func (p *Pets) SetName(argsName string) {
	p.Name = argsName
}

func (p *Pets) ClearName() {
	p.Name = ""
}

// 还能这样...就是清空reference的属性
func SetNameToPEPE(p *Pets) {
	p.Name = "PEPE"
}

func main() {
	MyPets := Pets{}
	MyPets.Name = "Qiqi"
	MyPets.Kind = "Neko"
	MyPets.Age = 2

	/*fmt.Println("My Pets properties: ", MyPets)
	MyPets.SetName("Qibao")
	fmt.Println("Now my Pets name is :", MyPets.Name)
	MyPets.ClearName()
	fmt.Println("Now my Pets name is :", MyPets.Name)
	*/
	SisterPets := Pets{"Mimi", "Dog", 4}

	// 传Pets指针
	SetNameToPEPE(&MyPets)
	SetNameToPEPE(&SisterPets)

	fmt.Println("now My Pet Name is: ", MyPets.Name, "\nMy Sister Pet Name is: ", SisterPets.Name)
	// PEPE
}

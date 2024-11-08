package main

import (
	"fmt"
)

// INTERFACE AND POLYMORPHISM
type AnimalIF interface {
	Run()
	GetColor()
}

type Duck struct {
	Color string
}

type Dog struct {
	Color string
}

// 被Go的写法震撼到了...直接用个指针就能implementation了

func (this *Duck) Run() {
	fmt.Println("the Duck is Running...")
}

func (this *Duck) GetColor() {
	fmt.Println("the Duck Color is :", this.Color)
}

func (this *Dog) GetColor() {
	fmt.Println("the Dog Color is :", this.Color)
}

func (this *Dog) Run() {
	fmt.Println("the Dog is Running...")
}

func showAnimal(animal AnimalIF) {
	animal.GetColor()
	animal.Run()
}

func main() {
	MyDuck := Duck{
		"Yellow",
	}

	MyDog := Dog{
		"Green",
	}

	showAnimal(&MyDuck)
	showAnimal(&MyDog)
}

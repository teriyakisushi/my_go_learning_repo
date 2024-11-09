package main

import (
	"fmt"
	"reflect" // reflect , to check type, value of args(or struct)
)

type Pet struct {
	Name string
	Age  int
	Size float64
}

func DetectSth(args interface{}) {
	// It will output the type, value of args
	fmt.Println(reflect.TypeOf(args))
	fmt.Println(reflect.ValueOf(args))
}

func ForeachArgs(args interface{}) {
	// type of args
	ArgsType := reflect.TypeOf(args)
	ArgsName := ArgsType.Name()
	ArgsValue := reflect.ValueOf(args)
	fmt.Println("Args Name = ", ArgsName, " Value = ", ArgsValue)

	// Get field of struct
	// Using method: args.NumFields to get field nums of struct
	for i := 0; i < ArgsType.NumField(); i++ {
		field := ArgsType.Field(i)
		value := ArgsValue.Field(i).Interface()

		fmt.Println(field.Name, field.Type, value)
		// It will output the Struct memebrs' name, type and value
	}
}

func main() {
	FLOAT_VAR := 3.14159
	DetectSth(FLOAT_VAR)

	MyPet := Pet{
		"Qiqi",
		12,
		50.12,
	}

	ForeachArgs(MyPet)
}

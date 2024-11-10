package main

import (
	"encoding/json"
	"fmt"
)

type Pet struct {
	// 在member属性后加`` 表示tags
	PetName string  `json:"name"`
	PetKind string  `json:"kind"`
	PetSize float64 `json:"size"`
}

func main() {
	MyPet := Pet{
		"Qiqi",
		"Dog",
		252.12,
	}

	json_str, err := json.Marshal(MyPet)
	if err != nil {
		fmt.Println("Json marshal error! ", err)
		return
	}

	fmt.Printf("Json str = %s\n", json_str)
	// name: , kind: , size:

	// Unmarshal
	// 使用json.unmarshal( Bytes, &ObjectPointer)
	mypet := Pet{}
	err = json.Unmarshal(json_str, &mypet) // <----- 是传进来的指针噢
	if err != nil {
		fmt.Println("Json unmarshal error! ", err)
		return
	}
	fmt.Println("After unmarshal : \n", mypet)
}

package main

import (
	"encoding/json"
	"fmt"
)

/*
* interface
 */

type Pet struct {
	Name string  `json:"name"`
	Kind string  `json:"kind"`
	Age  int     `json:"age"`
	Size float64 `json:"size"`
}

func (p *Pet) getName() string {
	return p.Name
}

func (p *Pet) getFull() []byte {
	petInfo, err := json.Marshal(p)
	if err != nil {
		fmt.Println("get pet infomation failed! ", err)
	}
	return petInfo
}

func (p *Pet) checkFull(src []byte) {
	dst := Pet{}
	err := json.Unmarshal(src, &dst)
	if err != nil {
		fmt.Println("check pet infomation failed! ", err)
		return
	}
	fmt.Println("Src info: ", dst)
}

func (p *Pet) RestName() {
	p.Name = ""
}

func (p *Pet) createPet(args interface{}) {
}

func main() {
}

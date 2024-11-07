package main

import "fmt"

func use_of_map() {
	/*
	* map的基本使用方式
	*
	 */

	campusMap := make(map[string]string)
	campusMap["Beijing"] = "Peking University"
	campusMap["Zhejiang"] = "Zhejiang University"

	fmt.Println("First version: ")
	for key, value := range campusMap {
		fmt.Println("Key = ", key, " Value = ", value)
	}

	fmt.Println("After Delete")
	// 可以用delete关键字删除map中的元素
	// Usage: delete(MapName, Key)
	delete(campusMap, "Beijing")

	for key, value := range campusMap {
		fmt.Println("Key = ", key, " Value = ", value)
	}

	// 如果将map传参的话，是将map的指针传过去的= =
}

func main() {
	/*
	* Map的声明是 map[key]value
	* 定义得使用make（不使用直接声明&赋值也行) 为其开辟空间
	*	不特别写明len的话也是可以随时追加的
	*
	 */

	// 声明方式1
	// map底层是hash实现的，直接打印的话按index顺序可能会乱

	CityMap := make(map[string]string)
	CityMap["Guangdong"] = "Guangzhou"
	CityMap["Zhejiang"] = "Hangzhou"
	CityMap["Guizhou"] = "Guiyang"
	CityMap["Sichuan"] = "Chengdu"

	for key, value := range CityMap {
		fmt.Println("Key = ", key, " Value = ", value)
	}

	// 声明方式2
	myPetMap := map[string]string{
		"Dog": "Qiqi",
		"Cat": "Meme",
		"Fox": "Aka",
	}

	fmt.Println(myPetMap)

	use_of_map()
}

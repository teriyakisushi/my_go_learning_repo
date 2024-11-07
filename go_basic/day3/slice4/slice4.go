package main

import "fmt"

func main() {
	// 开辟了一个Cap为5，但实际有效空间为3的Slice
	// 能访问的空间只有3个，后面2个不能访问
	// 若append的长度到达了Cap值，这个Slice会继续加cap（扩容大小为原来设置的cap值)

	myTinySlice := make([]int, 3, 5)
	fmt.Println("Len = ", len(myTinySlice), " Cap = ", cap(myTinySlice), " Slice = ", myTinySlice)

	// Append!
	myTinySlice = append(myTinySlice, 2, 3, 4)

	fmt.Println("Len = ", len(myTinySlice), " Cap = ", cap(myTinySlice), " Slice = ", myTinySlice)
	// Len = 6 , Cap = 10

	/*
	* len 表示左指针到右指针之间的距离
	* Cap 表示左指针到底层数组末尾的距离
	* cap扩容实际还是开辟了新的内存空间
	 */

	/*
	*  拷贝Slice
	*	 可以将底层数组的slice一起copy
	*  Copy(Dst, Src)
	 */

	myCopySlice := make([]int, 5)
	copy(myCopySlice, myTinySlice)

	fmt.Println("the Copy Slice is :", myCopySlice)
	// 0 0 0 2 3 4
}

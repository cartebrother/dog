package main

import (
	"fmt"
	"unsafe"
)

var arr = [6]int64{1, 2, 3, 4, 5}

var arrB [10]int64

var arrC = [...]int64{22, 23, 24}

var arrD = [...]int64{88: 1}

func foo(arr [5]int64) {

}

func main() {
	var arr1 [5]int64
	//var arr2 [6]int64
	//var arr3 [5]string

	foo(arr1)
	//foo(arr2)
	//foo(arr3)

	fmt.Println("数组长度是：", len(arr))
	fmt.Println("数组的大小是：", unsafe.Sizeof(arr), "byte")

	fmt.Printf("%T\n", arrB)

	for i, i2 := range arrB {
		fmt.Println("index:", i, " value:", i2)
	}

	fmt.Printf("%T\n", arrC)

	for i, i2 := range arrC {
		fmt.Println("index:", i, " value:", i2)
	}

	fmt.Printf("%T\n", arrD)

	fmt.Println(arrD[10], arrD[12])

	//fmt.Println(arrD[1000],arrD[-100])

}

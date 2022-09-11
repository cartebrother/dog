package main

import (
	"fmt"
)

func main() {

	//func1()

	//func2()

	func4()

	func44()

	func441()

	func442()

	func443()

	str1()

	map1()

	channel1()

}

func channel1() {
	println("====channel for range =========")

	var c = make(chan int)
	for v := range c {
		println(v)
	}
}

func map1() {
	println("============map for range===============")

	var m = map[string]int64{
		"rob":  67,
		"russ": 39,
		"john": 29,
	}
	for k, v := range m {
		println(k, "=", v)
	}
}

func str1() {
	println("=======for range string ===========")

	var s = "中华人民共和国"
	for k, v := range s {
		fmt.Printf("%d %s %x \n", k, string(v), v)
	}

	// v unicode字符码点
	// k 字符码点内存编码第一个字节在字符串内存序列中的位置
}

func func4() {
	var s = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %d\n", i, s[i])
	}
}

func func44() {
	println("=====range================")

	var s = []int{1, 2, 3, 4, 5}
	for k, v := range s {
		fmt.Printf("s[%d] = %d \n", k, v)
	}
}

func func441() {
	println("========key=============")

	var s = []int{1, 2, 3, 4, 5}
	for k := range s {
		fmt.Printf("s[%d]  \n", k)
	}
}

func func442() {
	println("========_ for v=============")

	var s = []int{1, 2, 3, 4, 5}
	for k, _ := range s {
		fmt.Printf("s[%d]  \n", k)
	}
}

func func443() {
	println("======== no care for k,v =============")

	var s = []int{1, 2, 3, 4, 5}
	for range s {
		fmt.Println("s= ", s)
	}
}

func func3() {

	for i := 0; i < 10; {
		i++
		println(i)
	}

	j := 0
	for ; j < 10; j++ {
		println(j)
	}

	k := 0
	for k < 10 {
		println(k)
		k++
	}

	z := 0
	for z < 10 {
		println(z)
	}

	for {
		fmt.Println("aaa")
	}

}

func func2() {
	var sum int
	for i, j, k := 0, 1, 2; (i < 20) && (j < 10) && (k < 30); i, j, k = i+1, j+1, k+5 {
		sum += (i + j + k)
		println(sum)
	}
}

func func1() {
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

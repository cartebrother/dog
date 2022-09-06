package main

import "fmt"

//var m11 = map[int64]string
//var m12 = map[string]string

func main() {

	//s1 := make([]int ,1)
	//s2 := make([]int ,2)
	//
	//f1:= func() {}
	//f2:= func() {}
	//
	//m1:=make(map[int]string)
	//m2:=make(map[int]string)

	//println(s1 == s2)
	//println(f1==f2)
	//println(m1==m2)
	//
	//
	//var m3 = map[string]int64

	m1 := map[int64]string{}
	fmt.Println(m1)

	m2 := map[int64][]string{
		1: []string{"a", "b"},
		2: []string{"hello", "world"},
	}

	fmt.Println(m2)

	m3 := map[Position]string{
		Position{1.0, 1.0}: "1",
		Position{2.0, 2.0}: "2",
	}
	fmt.Println(m3)

	//省略字面值的类型
	m33 := map[Position]string{
		{1.0, 1.0}: "1",
		{2.0, 2.0}: "2",
	}
	fmt.Println(m33)

	m5 := make(map[int64]string)
	println(m5)
	m6 := make(map[int64]string, 8)
	m6[1] = "hello"
	m6[2] = "world"
	m6[3] = "carter"
	m6[4] = "go"

	fmt.Println("m6:", m6)

	fmt.Println("m6的长度是：", len(m6))

	v := m6[100]
	v1 := m6[1]
	println("m6[100] value:", v)
	println("m6[1] value:", v1)

	v2, ok := m6[2]
	if !ok {
		fmt.Println("m6[2] not exist")
	}
	println("m6[2] value:", v2)

	_, ok10 := m6[10]
	println("m6[10] exist:", ok10)

	fmt.Println("before delete: ", m6)
	delete(m6, 3)
	fmt.Println("after delete:", m6)

	for key, val := range m6 {
		fmt.Println("key:", key, ",val:", val)
	}

	//多次循环输出对比

	for i := 0; i < 3; i++ {
		iter(m6)
	}

	//内部和外部使用
	println("before inner op,out")
	for key, val := range m6 {
		fmt.Println("key:", key, ",val:", val)
	}

	iter2(m6)

	println("after inner op,out")

	for key, val := range m6 {
		fmt.Println("key:", key, ",val:", val)
	}
	//var m4 map[string]int64
	//
	//m4["key"] = 1

}

func iter(m6 map[int64]string) {
	println("==========================")
	for key, val := range m6 {
		fmt.Println("key:", key, ",val:", val)
	}
}

func iter2(m6 map[int64]string) {
	println("after out op==========================")
	m6[112] = "please call 112"
	for key, val := range m6 {
		fmt.Println("key:", key, ",val:", val)
	}
}

type Position struct {
	x float64
	y float64
}

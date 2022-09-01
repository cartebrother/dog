package main

import "fmt"

type myInt int

const n myInt = 13

// Cannot use 'n+5' (type myInt) as the type int -> 做类型转换
const m int = int(n) + 5

func main() {

	var a int = 5
	//Invalid operation: a + n (mismatched types int and myInt) -> 做类型转换
	fmt.Println(a + int(n))

}

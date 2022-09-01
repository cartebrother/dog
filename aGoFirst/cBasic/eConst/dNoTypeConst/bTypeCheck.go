package main

import "fmt"

type myInt int

const n = 13

// Cannot use 'n+5' (type myInt) as the type int -> no type const
const m int = n + 5

func main() {

	var a int = 5
	//Invalid operation: a + n (mismatched types int and myInt)
	fmt.Println(a + n)

}

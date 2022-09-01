package main

import "fmt"

type myInt int

const n myInt = 13

// Cannot use 'n+5' (type myInt) as the type int
const m int = n + 5

func main() {

	var a int = 5
	//Invalid operation: a + n (mismatched types int and myInt)
	fmt.Println(a + n)

}

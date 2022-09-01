package main

import "fmt"

const m = 13333333333333333333

var k int8 = 1

func main() {
	// # command-line-arguments
	//.\eTypeConvertErr.go:11:11: m (untyped int constant 13333333333333333333) overflows int8
	j := k + m

	fmt.Println(j)

}

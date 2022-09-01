package main

import "fmt"

const (
	Apple, Banana = 11, 22
	// 	Strawberry, Grape = 11,22
	Strawberry, Grape
	//  Pear, Watermelon = 11,22
	Pear, Watermelon
)

func main() {

	fmt.Println("apple: ", Apple)
	fmt.Println("Banana: ", Banana)

	fmt.Println("Strawberry: ", Strawberry)
	fmt.Println("Grape: ", Grape)

	fmt.Println("Pear: ", Pear)
	fmt.Println("Watermelon: ", Watermelon)

}

package main

import (
	"fmt"
	"time"
)

func main() {

	continue1()

	continue2()

	break1()

	break2()

	//forE1()
	//
	//forE11()

	forE2()

	forE3()

}

func forE3() {
	println("map index ================")
	var m = map[string]int{
		"carter": 33,
		"perter": 40,
		"seven":  200,
	}

	i := 0
	for k, v := range m {
		if i == 0 {
			delete(m, "carter")
		}
		i++
		fmt.Println(k, "=", v)
	}

	fmt.Println("i = ", i)

}

func forE2() {
	println("===========loop range array copy")

	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("origin a: ", a)

	for i, v := range a {
		if i == 0 {
			a[1] = 100
			a[2] = 100
		}
		r[i] = v
	}

	fmt.Println("after loop range : a: ", a)
	fmt.Println("after loop range : r: ", r)

}

func forE1() {
	println("=========================k,v exception")
	var m = []int{1, 2, 3, 4, 5}
	for k, v := range m {
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println(k, v)
		}()
	}

	time.Sleep(time.Second * 18)
}

func forE11() {
	println("=========================k,v exception fix")
	var m = []int{1, 2, 3, 4, 5}
	for k, v := range m {
		go func(k, v int) {
			time.Sleep(time.Second * 3)
			fmt.Println(k, v)
		}(k, v)
	}

	time.Sleep(time.Second * 18)
}

func break2() {
	println("============show break loop")
	var gold = 38
	var s = [][]int{
		{1, 34, 26, 35, 78},
		{3, 45, 13, 24, 99},
		{101, 13, 38, 7, 127},
		{54, 27, 40, 83, 81},
	}

outLoop:
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == gold {
				fmt.Println("found ", gold, " at ", i, j)
				break outLoop
			}
		}
	}
}

func break1() {

	println("show break ========================")

	var s = []int{5, 19, 6, 3, 8, 12}
	var firstEven int = -1

	for i := 0; i < len(s); i++ {
		if s[i]%2 == 0 {
			firstEven = s[i]
			break
		}
	}

	println(firstEven)

}

func continue2() {
	println("============show continue loop")

	var s = [][]int{
		{1, 34, 26, 35, 78},
		{3, 45, 13, 24, 99},
		{101, 13, 38, 7, 127},
		{54, 27, 40, 83, 81},
	}

outLoop:
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == 13 {
				fmt.Println("found 13 at ", i, j)
				continue outLoop
			}
		}
	}

}

func continue1() {

	var sum int
	var s = []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(s); i++ {
		if s[i]%2 == 0 {
			continue
		}
		sum += s[i]
	}
	println(sum)
}

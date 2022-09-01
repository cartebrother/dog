package main

import "fmt"

const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexWaiterShift      = iota
	starvationThresholdNs = 1e6
)

const (
	Apple, Banana = iota, iota + 10
	Strawberry, Grape
	Pear, Watermelon
)

// 从1开始定义枚举
const (
	_ = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	_ = iota
	Pin1
	Pin2
	Pin3
	_
	Pin5
)

const (
	Blue   = 1
	Black  = 2
	Red    = 3
	Yellow = 4
)

const (
	a = iota + 1
	b
	c
)
const (
	i = iota << 1
	j
	k
)

func main() {

	fmt.Println(mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift, starvationThresholdNs)

	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	fmt.Println(Pin1, Pin2, Pin3, Pin5)

}

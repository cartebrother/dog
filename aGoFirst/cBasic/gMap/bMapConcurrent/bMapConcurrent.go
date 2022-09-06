package main

import "time"

func main() {

	m1 := make(map[int64]string, 8)
	m1[1] = "hello"
	m1[2] = "world"
	m1[3] = "carter"
	m1[4] = "go"

	p := m1[1]
	println(p)

	//p1 := &m1[1]
	//println(p1)

	go func() {
		for k, v := range m1 {
			m1[k] = v + "1"
		}
	}()

	go func() {

		for i := 0; i < 1000; i++ {

			for k, v := range m1 {
				println("k:", k, " v:", v)
			}
		}

	}()

	time.Sleep(5 * time.Second)

}

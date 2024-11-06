package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 0
	m[2] = 2
	m[3] = 3

	delete(m, 3)
	delete(m, 4)
	fmt.Println(m[3])
	fmt.Println(m[1])
	_, ok := m[4]
	if !ok {
		fmt.Println("m[4] is No")
	}
}

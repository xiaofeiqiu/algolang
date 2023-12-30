package main

import "fmt"

type T struct {
	A int
}

func main() {
	a := make([]int, 1)

	f(append(a, 10))

	fmt.Println(a)
}

func f(a []int) {
	fmt.Println(a)
}

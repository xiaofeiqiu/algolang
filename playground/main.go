package main

import "fmt"

func main() {
	t1 := []int{1}
	fmt.Println("Before modify:", t1) // Print before modification

	modify(t1)

	fmt.Println("After modify:", t1) // Print after modification
}

func modify(bbb []int) {
	bbb[0] = 100
}

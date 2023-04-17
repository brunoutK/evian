package main

import "fmt"

func main() {

	var tiger int = 10
	var ptr *int
	ptr = &tiger
	*ptr = 11
	fmt.Println("tiger的值是: ", tiger)
	fmt.Println("ptr的值是: ", *ptr)
}

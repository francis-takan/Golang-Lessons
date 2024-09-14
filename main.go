package main

import "fmt"

func main() {
	fmt.Println("Hello there...")

	fmt.Println(add(2, 5))
}

func add(x int, y int) int {
	return x + y
}

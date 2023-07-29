package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("Again")

	repeats := 5
	for repeats > 0 {
		fmt.Println(repeats)
		repeats--
	}

	val := add(1, 2)
	fmt.Println(val)

	val = mul(val, 3)
	fmt.Println(val)

	fmt.Println(sub(val, -7))
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func sub(a, b int) int {
	return a - b
}

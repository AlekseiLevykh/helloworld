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
}

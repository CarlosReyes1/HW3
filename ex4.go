package main

import "fmt"

func main() {
	var smaller int
	var sum int

	fmt.Println("Enter a small number here: ", &small)
	fmt.Scan(&smaller)
	fmt.Println("Enter a larger number here: ", &large)
	fmt.Scan(&larger)
	fmt.Println(larger, "%", smaller, "=", larger%smaller)
}

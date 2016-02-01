package main

import "fmt"

func main() {
	var name string
	fmt.Scan("Enter your name: ", &name)
	fmt.Println("Hello ", name)
}

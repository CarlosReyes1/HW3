package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	num, even := half(10)
	fmt.Println(num, even)
}

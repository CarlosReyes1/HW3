package main

import "fmt"

func greatestNum(nums ...int) int {
	var largest int
	for _, i := range numbers {
		if i > largest {
			largest = i
		}
	}
	return largest
}

func main() {
	greatest := greatestNum(2, 4, 6, 7, 9, 10, 20, 50, 100)
	fmt.Println(greatest)
}

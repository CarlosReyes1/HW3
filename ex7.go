package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
//the answer to the problem number 1 in projecteuler.net is 233168

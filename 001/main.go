package main

import "fmt"

func main() {
	var sum = 0

	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Printf("result is: %v \n", sum)
}

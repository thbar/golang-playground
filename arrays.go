package main

import "fmt"

func main() {
	var numbers [10]int

	for i := 0; i < len(numbers); i++ {
		numbers[i] = i * 3
	}

	fmt.Println("Here are some numbers:", numbers)

	for i := 0; i < len(numbers); i++ {
		var parity string
		if numbers[i]%2 == 0 {
			parity = "even"
		} else {
			parity = "odd"
		}
		fmt.Printf("Number %d is %s\n", numbers[i], parity)
	}
}

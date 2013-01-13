package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome! What's your name?")
	var name string
	fmt.Scanf("%s", &name)

	fmt.Println("Nice to have you onboard,", name)

	for i := 0; i < 10; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}

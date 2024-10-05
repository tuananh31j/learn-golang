package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Enter your String: ")
	fmt.Scanln(&input)

	isValid := strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a")

	if isValid {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
	
}
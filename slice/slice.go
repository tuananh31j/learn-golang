package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	stotedSlice := make([]int, 0, 3)

	for {
		fmt.Println("Enter an integer or X to exit: ")
		var input string
		fmt.Scanln(&input)

		
		if input == "X" {
			break;
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Invalid input: %v\n", input)
			continue
		}

		stotedSlice = append(stotedSlice, num)
		sort.Ints(stotedSlice)

		fmt.Printf("Stored: %v",stotedSlice)
	}
}
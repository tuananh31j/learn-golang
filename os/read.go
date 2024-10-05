package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define a struct to hold the first and last names
type Name struct {
	Fname string
	Lname string
}

func main() {
	// Prompt the user to enter the file name
	fmt.Print("Enter the name of the text file: ")
	var filename string
	fmt.Scan(&filename)

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a slice to hold the names
	var names []Name

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into first and last name
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			// Create a new Name struct and append it to the slice
			name := Name{
				Fname: parts[0],
				Lname: parts[1],
			}
			names = append(names, name)
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print out the names from the slice
	fmt.Println("\nNames from the file:")
	for _, name := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", name.Fname, name.Lname)
	}
}

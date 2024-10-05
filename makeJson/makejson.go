package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	user := make(map[string] string)
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter your address: ")
	fmt.Scanln(&address)
	user["name"] = name
	user["address"] = address
	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println("User JSON could not be created")
	}
	fmt.Println(string(userJson))
}
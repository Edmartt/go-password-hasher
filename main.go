package main

import (
	"fmt"
	"hashing/hasher"
)

func main() {
	var password string
	var currentPassword string

	fmt.Print("Type the word you want as password: ")
	fmt.Scanln(&password)

	hashed_password := hasher.ConvertToHash(password)

	fmt.Println(hashed_password)

	fmt.Print("Type your current password: ")
	fmt.Scanln(&currentPassword)

	response := hasher.CheckHash(hashed_password, currentPassword)

	fmt.Println("My response: ", response)
}

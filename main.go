package main

import (
	"fmt"
	"hashing/hasher"
)

func main() {
	var password string
	fmt.Print("Type the word you want as password: ")
	fmt.Scanln(&password)
	hashed_password := hasher.HashMaker([]byte(password))
	fmt.Println(hashed_password)
}

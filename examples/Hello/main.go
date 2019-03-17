package main

import (
	"fmt"
)

const (
	name = "Hello, My Name is Nikul."
)

func main() {
	var name string
	fmt.Print("Enter your name:")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello my name is: ", name)
}

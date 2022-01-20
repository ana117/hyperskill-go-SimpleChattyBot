package main

import "fmt"

func main() {
	fmt.Println("Hello! My name is Aid.")
	fmt.Println("I was created in 2020.")
	fmt.Println("Please, remind me your name.")

	// reading a name
	var name string
	fmt.Scan(&name)

	fmt.Println("What a great name you have, " + name + "!")
}

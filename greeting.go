package main

import "fmt"

func main() {
	var name string
	fmt.Println("Please enter your name")
	fmt.Scanf("%s", &name)

	if name == "Bob" || name == "Alice" {
		fmt.Printf("Greetings to %s!\n", name)
	}
}

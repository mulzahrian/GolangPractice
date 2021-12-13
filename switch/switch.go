package main

import "fmt"

func main() {
	name := "rian"

	switch name {
	case "rian":
		fmt.Println("hellow rian")
	case "ari":
		fmt.Println("hellow ari")
	default:
		fmt.Println("hiii, what ur name")
	}

	//short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name to long")
	case false:
		fmt.Println("Name was right")
	}

	//switch witout condition
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("To long")
	case length > 5:
		fmt.Println("middle")
	default:
		fmt.Println("Thats right")
	}
}

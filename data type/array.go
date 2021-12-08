package main

import "fmt"

func main() {
	//manual
	var names [3]string
	names[0] = "mulzahrian"
	names[1] = "ari"
	names[2] = "one"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// not
	var values = [3]int{
		90,
		80,
		95,
	}
	fmt.Println(values)

	fmt.Println(len(values))
	fmt.Println(len(names))
}

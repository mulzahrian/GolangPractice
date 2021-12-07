package main

import "fmt"

func main() {
	const firstname string = "mulzahrian"
	const lastname = "one"

	fmt.Println(firstname)
	fmt.Println(lastname)

	//in constant u can't change the value after u initial on the first one
	//it will be error
	//firstname = "Can't Change"
	//lastname = "Can't change"
}

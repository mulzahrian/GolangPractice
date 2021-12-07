package main

import "fmt"

func main(){
	const(
		firstname string = "Mulzahrian"
		lastname = "One"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)
	// error
	//firstname = "Can't Change"
	//lastname = "Can't Change"
}

package main

import "fmt"

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	// resultInt := result.(int) //panic
	// fmt.Println(resultInt)

	//when u don't wanna be panic we can use switch
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("Unknown")
	}
}

func random() interface{} {
	return "OK"
}

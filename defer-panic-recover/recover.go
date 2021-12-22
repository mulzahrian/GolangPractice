package main

import "fmt"

func endApp() {
	message := recover()
	fmt.Println("Error Message :", message)
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
	fmt.Println("Application Run")
}

func main() {
	runApp(true)
}

package main

import "fmt"

func main() {
	runApplication()
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func logging() {
	fmt.Println("Done Call the Function")
}

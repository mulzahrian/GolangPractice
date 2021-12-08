package main

import "fmt"

func main() {
	type KTPNumber string

	var ktpAri KTPNumber = "11111111"
	fmt.Println(ktpAri)
	fmt.Println(KTPNumber("2222222"))
}

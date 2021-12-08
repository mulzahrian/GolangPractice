package main

import "fmt"

func main() {
	var finalResult = 90
	var atendent = 80

	var pastTheFinalResult bool = finalResult > 80
	var pastTheAtendent bool = atendent > 80

	var past bool = pastTheFinalResult && pastTheAtendent

	fmt.Println(past)

	//fast way
	fmt.Println(finalResult >= 80 && atendent >= 80)
}

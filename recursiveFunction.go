package main

import "fmt"

//example is using Factorial
//example is 5 its mean 5*4*3*2*1
//this example is without recursive
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}
func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}

//using recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

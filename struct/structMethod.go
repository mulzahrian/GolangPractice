package main

import (
	"fmt"
)

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, My Name is", customer.Name)

}
func main() {
	one := Customer{Name: "one"}
	one.sayHello()
}

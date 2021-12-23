package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	person := NewMap("Rian")
	fmt.Println(person)

	var data map[string]string = NewMap("Ari")
	if data == nil {
		fmt.Println("Empty Data")
	} else {
		fmt.Println(data)
	}
}

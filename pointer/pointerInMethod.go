package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	rama := Man{"Rama"}
	rama.Married()

	fmt.Println(rama.Name)

}

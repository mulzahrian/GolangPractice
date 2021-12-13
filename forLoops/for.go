package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Looping", counter)
		counter++
	}
	//For Statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Looping", counter)
	}
	//slice
	slice := []string{"Rian", "Ari", "One"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	//for range
	names := []string{"rian", "one", "ari"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
	//map

	person := make(map[string]string)
	person["name"] = "Ari"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}

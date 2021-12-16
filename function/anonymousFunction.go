package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

// func blackListAdmin(name string) bool {
// 	return name == "admin"
// }

// func blackListRoot(name string) bool {
// 	return name == "root"
// }
func main() {
	BlackList := func(name string) bool {
		return name == "dog"
	}

	registerUser("one", BlackList)
	registerUser("dog", func(name string) bool {
		return name == "dog"
	})

}

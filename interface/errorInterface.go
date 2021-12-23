package main

import (
	"errors"
	"fmt"
)

func devision(value int, devision int) (int, error) {
	if devision == 0 {
		return 0, errors.New("Devision with zero")
	} else {
		return value / devision, nil
	}
}

func main() {
	result, err := devision(100, 0)
	if err == nil {
		fmt.Println("Result", result)
	} else {
		fmt.Println("Error", err.Error())
	}

}

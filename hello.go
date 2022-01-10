package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, world.")
	answer := "mmm"
	a, b := yesNo(answer)
	fmt.Println(a)
	fmt.Println(b)
}

func yesNo(answer string) (bool, error) {
	if answer == "yes" {
		return true, nil
	} else if answer == "no" {
		return true, nil
	}
	return false, errors.New("this is not Yes or No")
}

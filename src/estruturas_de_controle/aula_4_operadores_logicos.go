package main

import "fmt"

func main() {
	x := 6

	if x == 2 || x == 3 {
		fmt.Println("x == 2 ou x == 3")
	} else {
		fmt.Println("x != 2 e x != 3")
	}

	if x%2 == 0 && x%3 == 0 {
		fmt.Println("multiplo de 2 e 3")
	} else {
		fmt.Println("nem de 2 nem de 3")
	}
}

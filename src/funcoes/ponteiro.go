package main

import "fmt"

func inicial(xPtr *int) {
	*xPtr = 3
}

func main() {
	x := 5
	inicial(&x)

	fmt.Println(x)
}

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("O número %d é par\n", i)
		} else {
			fmt.Printf("O número %d é impar\n", i)
		}
	}
}

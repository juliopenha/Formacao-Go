package main

import "fmt"

func main() {
	i := 0

	for i < 20 {
		fmt.Println(i)
		i++
		if i == 20 {
			fmt.Println("break no if")
			break
		} else {
			fmt.Println("continue")
			continue
		}
	}
}

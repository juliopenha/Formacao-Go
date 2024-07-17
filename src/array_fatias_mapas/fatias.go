package main

import "fmt"

func main() {
	// var x[] float64
	// fatia:= make([]float64, 4)
	// fatia:=[low:high]
	// fatia:arr[0:5]

	arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	// fatia := make([]float64, 5)
	fatia := arr[2:5]
	fmt.Println(fatia)

	fatia1 := []int{1, 2, 3}
	fatia2 := append(fatia1, 4, 5)

	fmt.Println(fatia1, fatia2)

	fatia3 := []int{1, 2, 3}
	fatia4 := make([]int, 2)

	copy(fatia4, fatia3)

	fmt.Println(fatia3, fatia4)
}

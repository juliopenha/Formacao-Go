package main

import "fmt"

func main() {
	x := make(map[string]int)
	y := make(map[int]int)

	x["chave"] = 10
	y[1] = 20

	fmt.Println(x["chave"])
	fmt.Println(y[1])

	elemento := make(map[string]string)

	elemento["H"] = "Hidrogenio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "Litio"

	fmt.Println(elemento["He"])
}

package main

import "fmt"

const ebulicaoK = 373.0

func main() {
	tempK := ebulicaoK
	tempC := (tempK - 273.0) // K para C

	fmt.Printf("A temperatura de ebulição da água em ºK = %g e A temperatura de ebulição da água em ºC = %g\n", tempK, tempC)
}

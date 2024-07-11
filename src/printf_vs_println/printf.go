// declarar pacote principal
package main

// importar função fmt
import "fmt"

// declarar const
const ebulicaoF = 212.0

// declarar função principal
func main() {
	// declarar variáveis
	// operador curto só funciona dentro de funções
	tempF := ebulicaoF
	tempC := (tempF - 32) * 5 / 9 // F para C

	// mostrar resulado
	// fmt.Println("A temperatura de ebulição da água em ºF = ", tempF)
	// fmt.Println("A temperatura de ebulição da água em ºC = ", tempC)

	fmt.Printf("A temperatura de ebulição da água em ºF = %g e A temperatura de ebulição da água em ºC = %g\n", tempF, tempC)

	// deve aparecer F = 212 e C = 100
}

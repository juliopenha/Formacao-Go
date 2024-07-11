// declarar pacote principal
package main

// importar função fmt
import "fmt"

var a int = 30
var b string = "Nome"

// declarar função principal
func main() {

	var c float64 = float64(a)

	fmt.Printf("O Valor de c é: %g (%T) e o valor de b é: %s (%T)\n", c, c, b, b)

}

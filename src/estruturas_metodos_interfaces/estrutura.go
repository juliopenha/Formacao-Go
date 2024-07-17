package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {
	// estrutura é struct
	ana := pessoa{"Ana", 54}

	var joao pessoa

	joao.nome = "João"
	joao.idade = 20

	fmt.Println(ana)
	fmt.Println(pessoa{"Ana", 54})
	fmt.Println(joao)
	fmt.Println(pessoa{"João", 20})
}

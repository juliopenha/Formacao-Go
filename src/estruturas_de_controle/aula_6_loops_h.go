package main

import "fmt"

func main() {
	for horas := 0; horas <= 12; horas++ {
		for minutos := 0; minutos < 60; minutos++ {
			fmt.Printf("Horas: %02d:%02d\n", horas, minutos)
		}
	}
}

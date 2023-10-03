package main

import (
	"fmt"
)

func calcularNota(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com a nota:", nota)
	} else {
		fmt.Println("Reprovado com a nota:", nota)
	}
}

func main() {
	calcularNota(8)
	calcularNota(5.2)
}
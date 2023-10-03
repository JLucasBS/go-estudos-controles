package main

import "fmt"

func notaPraConceito(nota float64) string {
	// Apenas para teste, o melhor caso para esse exemplo seria um Switch
	if nota >= 9  && nota <= 10 {
		return "A"
	} else if nota >= 8 && nota < 9 {
		return "B"
	} else if nota >= 7 && nota < 8 {
		return "C"
	} else if nota >= 3 && nota < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Printf(notaPraConceito(7.5))
}
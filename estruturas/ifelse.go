package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7{
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Não aprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(7.0)
	imprimirResultado(5.1)

}
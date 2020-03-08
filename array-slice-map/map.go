package main

import "fmt"

func main() {
	significados := make(map[string]string)

	significados["danke"] = "obrigado"
	significados["entschuldigung"] = "com licença"
	significados["weihnacht"] = "noite de natal"

	for chave, significado := range significados {
		fmt.Println(chave, significado)
	}

}

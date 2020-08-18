package main

import "fmt"

type Circle struct {
	Raio        int
	Comprimento int
}
type Cilinder struct {
	Haltura int
	Circle
}

func main() {
	cilinder := new(Cilinder)
	cilinder.Haltura = 11
	cilinder.Raio = 12
	cilinder.Comprimento = 13

	fmt.Println(cilinder)
}

package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}
func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "lápis",
		preco:    1.65,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())
	produto2 := produto{"detergente", 1000.0, 0.230}
	fmt.Println(produto2, produto2.precoComDesconto())

}

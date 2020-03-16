package main

import "fmt"

type imprimivel interface{
	toString() string

}
type pessoa struct{
	name string
	sobrenome string
}

type produto struct{
	name string
	preco float64
}

func (p pessoa) toString() string{
	return p.name + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.preco)
}

func imprimir(x imprimivel){
	fmt.Println(x.toString())
}
func main(){
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça Jeas", 79.90}
	imprimir(coisa)



}
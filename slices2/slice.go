package main

import (
	"fmt"
	"reflect"
)

func main(){
	a1 := [3]int{1,2,3} // array tem uma estrutura fixa definida na declaração
	s1 := []int{1,2,3} //slice é dinâmico não tendo podendo ter uma estrura fixa definida ou não

	//o slice é tipado de uma maneira diferente do array por conta de sua estrutura dinamica
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(s1), reflect.TypeOf(a1))

	//o slice pode ser como um vetor que referencia valores de um outro vetor
	//em outras palavras, o slice pode ser parte de um array
	a2 := [5]int{1,2,3,4,5}
	//nesse caso s2 é uma parte de a2 onde s2 = [2,3]
	s2 := a2[1:3]
	fmt.Println(reflect.TypeOf(s2), s2)
	//se mudarmos um valor de uma posição do slice, a posição referente no array será mudada também
	// devido a estrutura do slice ser implementada em ponteiros que referencia aos endereços que são partes do array ou slice original
	s3 := a2[0:2]
	fmt.Println(a2, s3[1])
	s3[1] = 66
	// perceba que ao mudarmos o valor de s[3] para 66 o valor referente em a2 também muda
	fmt.Println(a2, s3)

	//isso não ocorre com arrays
	a3 := [3]int{1,2,77}

	a2[3] = a3[2]
	fmt.Println(a2, a3)

	a2[3] = 55

	fmt.Println(a2, a3)

	fmt.Println(len(s2))


  a4 := [][]int{{1, 2},{1},{1,2,3}}
  fmt.Println(a4)
  fmt.Println(len(a4))

}
package main

import (
	"fmt"
	m "math"
	r "reflect"
)

func main(){
	fmt.Println(1,2,1000)
	fmt.Println("Literal inteiro é", r.TypeOf(32000))
	
	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", r.TypeOf(b))

	// com sinal... int8 int16 int32 int64	

	i1 := m.MaxInt64
	fmt.Println("O valor máximo do int i1 é", i1)
	fmt.Println("O tipo de i1 é", r.TypeOf(i1))
	
	var i2 rune = 'a' //representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", r.TypeOf(i2))
	fmt.Println(i2)

	//números reais (float32, float64)

	var x float32 = 49.99
	fmt.Println("O tipo de x é", r.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", r.TypeOf(49.99))

	//boolean
	bo := true 
	fmt.Println("O tipo de bo é", r.TypeOf(bo))
	fmt.Println(!bo)

	//string 
	s1 := "Olá meu nome é Lucas"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	s2 :=  `Olá    
	meu    
	nome
	é    
	Leo ` 
	fmt.Println("O tamanho da string é", len(s2))	

	char := 'a'    
	fmt.Println("O tipo de char é", r.TypeOf(char))    
	fmt.Println(char)
}



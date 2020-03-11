package main

import "fmt"

func produzir(c chan<- int){
	c<-1
	c<-2
	c<-3

	// Basta fechar o canal ao final da exegução.=
	close(c)
}

func main(){
	c := make(chan int, 3)
	go produzir(c)

	for valor := range c{
		fmt.Println(valor)
	}
}
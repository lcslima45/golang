package main

import "fmt"

func main(){
	/*c := make(chan int)
	go produzir(c)

	valor := <-c
	fmt.Println(valor)
*/
	c := make(chan int, 3)
	go produzir(c)
	fmt.Println(<-c, <-c, <-c, <-c)

}

func produzir(c chan int){
	c<-1
	c<-2
	c<-3
}

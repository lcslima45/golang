package main

import (
	"Teste/aurora"
	"fmt"
)

func main(){
	greetMessage := hello("")
	fmt.Println(aurora.Green(greetMessage))
	greetMessage =hello("John")
	fmt.Println(aurora.Blue(greetMessage))
}


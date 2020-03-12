package main

import (
	"http://github.com/logrusorgru/aurora"
	"fmt"
)

func main(){
	greetMessage := hello("")
	fmt.Println(aurora.Green(greetMessage))
	greetMessage =hello("John")
	fmt.Println(aurora.Blue(greetMessage))
}


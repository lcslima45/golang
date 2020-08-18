package main

import (
	"fmt"
)

func main() {
	char := make(map[string]int)
	str := string("djfdjkjkdfjkdfghkdfkjdfkdfjkdfghkdhkdfhgkdfhkdfgeqpiewiioroerogiutuigergireouierguyreuieghie")
	for _, s := range str {
		char[string(s)]++
	}
	for str, char := range char {
		fmt.Println(str, char)
	}
}

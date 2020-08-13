package main

import (
	"fmt"
)

func main() {
	a := "hello, world"
	fmt.Println(a[0:3])
	fmt.Println(a[:5])
	fmt.Println(a[7:])
	fmt.Println("goodbye" + a[5:])
	fmt.Println("Hello, 人物")
	fmt.Printf("dkfsldkls\t dlskfdsklsdlsdlk\a skdlfldsfsldksdlsldsdl\\ mdslldsklsd\r kdlsldklkslsdk\b jdlskjfldsklsdk \b dlsldsfksllsd\f")
	fmt.Println("\U00004e16")

	prefix := "Carlos"
	frase := "Carlos Barks Gosta de Cantar"
	suffix := prefix
	frase2 := "Meu amigo, Carlos"
	fmt.Println(HasPrefix(frase, prefix))
	fmt.Println(HasSuffix(frase2, suffix))
	n := 0
	for range frase2 {
		n++
	}
	fmt.Println(n)
}
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
func HasSuffix(s, suffix string) bool {
	return len(s) > len(suffix) && s[len(s)-len(suffix):] == suffix
}

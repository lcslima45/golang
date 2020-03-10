package main

func encaminhar(origem <- chan string, destino string){
	for {
		destino <- <-origem
	}
}
func juntar() (<-chan string) {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com")
		html.Titulo("https://www.amazon.com", "https://www.youtube.com")
	)

}
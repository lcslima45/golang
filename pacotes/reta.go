package main

import "math"

//iniciando com letra maíscula é PUBLICO	(visível dentro e fora do pacote) 
//iniciando com letra minúscula é PRIVADO  	(visível apenas dentro do pacote)
	
//Por exemplo...
//Ponto - gerará algo público
//ponto ou _POnto - gerará algo privado

//Ponto representa uma coordenada no plano cartesiano
type Ponto struct{
  x float64
  y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
  cx = math.Abs(b.x - a.x)
  cy = math.Abs(b.y - a.y)
  return 
}

func Distancia(a,b Ponto) float64{
  cx, cy := catetos(a,b) 
  return math.Sqrt(math.Pow(cx,2) + math.Pow(cy,2))
}

func main() {
  p1 := Ponto{2.0, 2.0}
  p2 := Ponto{2.0, 4.0}

  fmt.Println(catetos(p1, p2))
  fmt.Println(Distancia(p1,p2))
}

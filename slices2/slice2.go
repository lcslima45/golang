package main

import "fmt"

func main(){
	s1 := make([]int, 10,10)
	s2 := append(s1, 1,2,3)

	fmt.Println(s1, s2)

	s1[1] = 2

	fmt.Println(s1, s2)
  s3 := append(s2, 3,4,5)
  fmt.Println(s1,s2,s3)
	s3[2] = 51
	fmt.Println(s1, s2, s3)
	fmt.Print(cap(s1), cap(s2), cap(s3))
}
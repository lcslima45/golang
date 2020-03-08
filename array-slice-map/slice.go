package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{1, 2, 3}
	fmt.Println(a1, s1)

	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[2:5]
	fmt.Println(a2, s2)

	s4 := a2[:1]
	fmt.Println(s4)

	s5 := make([]int, 10)
	fmt.Println(s5)
	s5[9] = 12
	fmt.Println(s5)
	s5 = make([]int, 10, 20)
	fmt.Println(cap(s5))
	s5 = append(s5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s5, len(s5), cap(s5))
	s5 = append(s5, 1)
	fmt.Println(s5, len(s5), cap(s5))

}

package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {"A": 123.01},
		"J": {"B": 24.03},
		"E": {"C": 120.02},
	}

	delete(funcsPorLetra, "J")
	for a, b := range funcsPorLetra {
		fmt.Println(a, b)
	}
}

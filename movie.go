package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color, omitempty`
	Actors []string
}
type Bar struct {
	Nome      string
	Open      float64 `json:"Abertura"`
	High      float64 `json:"Máxima"`
	Low       float64 `json:"Mínima"`
	Close     float64 `json:"Fechamento"`
	Encerrado bool    `json:"Encerramento"`
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
		// ...
	}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	fmt.Println(movies)
	file, err := os.Create("movie.json")
	defer file.Close()
	file.Write(data)
	var bar = []Bar{
		{Nome: "Orama", Open: 12332.0, High: 3231312, Low: 32131313.342, Close: 23121.1212, Encerrado: true},
	}
	data, err = json.Marshal(bar)
	file2, err := os.Create("bar.json")
	defer file.Close()
	file2.Write(data)
	fmt.Println(file2)
	var bar2 []Bar
	jsonFile, err := os.Open("bar.json")
	defer jsonFile.Close()
	data, err = ioutil.ReadAll(jsonFile)
	json.Unmarshal(data, &bar2)
	fmt.Println(bar2)
}

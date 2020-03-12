package main

import "testing"

func TestHello(t *testing.T){
	//testando para argumento vazio
	emptyResult := hello("") //Deve retornar "Hello Dude!"

	if emptyResult != "Hello Dude!"{
		t.Errorf("hello failed expected %v got %v", "Hello Dude!", emptyResult)
	}

	//testando para argumento válido
	result := hello("Mike") //should return "Hello Mike!"

	if result != "Hello John!"{
		t.Errorf("hello failed expected %v got %v", "Hello John!", result)
	}


}
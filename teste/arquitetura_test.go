package teste

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T){
	t.Parallel() // Teste controlado paralelamente
	//runtime.GOARCH é a arquitentura em que o programa está rodando amd64, arm, s390x etc....
	if runtime.GOARCH =="amd64"{
		t.Skip("Não funciona em arquitetura amd64")

	}
	t.Fail()
}


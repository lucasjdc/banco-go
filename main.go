package main

import (
	"fmt"
	"banco/contas"
)

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(150)
	
	fmt.Println(contaExemplo.ObterSaldo()) 
}

package main

import (
	"fmt"
	"banco/contas"
)

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	contaDoDenis.Sacar(1000)
	
	fmt.Println(contaDoDenis.ObterSaldo())
}

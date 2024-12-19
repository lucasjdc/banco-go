package main

import (
	"fmt"
	"banco/contas"
	"banco/clientes"
)

func main() {
	clienteBruno := clientes.Titular{"Bruno", "123.456.789-12", "Professor"}
	contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}
	
	fmt.Println(contaDoBruno)
 
}

package main

import (
	"fmt"
  "banco/contas"
)

func main() {
  contaDaMaria := contas.ContaCorrente{Titular:"Maria",  Saldo: 300}
  contaDoGustavo := contas.ContaCorrente{Titular:"Gustavo", Saldo: 100}

  status := contaDaMaria.Transferir(200, &contaDoGustavo)
  fmt.Println(status)
  
  fmt.Println(contaDaMaria)
  fmt.Println(contaDoGustavo)
}

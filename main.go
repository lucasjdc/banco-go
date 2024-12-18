package main

import (
	"fmt"
)

type ContaCorrente struct {
	titular 		string
	numeroAgencia 	int
	numeroConta 	int
	saldo 			float64
}

func (c *ContaCorrente) Sacar(valorDoSaque  float64) string {
  podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
  if podeSacar {
    c.saldo -= valorDoSaque
    return "Saque realizado com sucesso"
  } else {
    return "Saldo  insuficiente"
  }
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
  if valorDoDeposito > 0 {
    c.saldo += valorDoDeposito
    return "Deposito realizado com sucessso." , c.saldo
  } else {
    return "Valor do deposito menor que zero.", c.saldo
  }
}

func main() {
	fmt.Println("\033[1;34m\nConta Corrente\n\033[0m")
  
  contaDaSilva := ContaCorrente{}
  contaDaSilva.titular = "Silvia"
  contaDaSilva.saldo = 500
  fmt.Println("Saldo:", contaDaSilva.saldo)
  fmt.Println(contaDaSilva.Sacar(1))
  fmt.Println("Saldo:", contaDaSilva.saldo)
  status, valor := contaDaSilva.Depositar(600)
  fmt.Println(status, valor)

	
}

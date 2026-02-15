package main

import (
	"alura/bank-project/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valorDoSaque float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 30)
	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuiza := contas.ContaCorrente{}
	contaDaLuiza.Depositar(200)
	PagarBoleto(&contaDaLuiza, 50)
	fmt.Println(contaDaLuiza.ObterSaldo())
}

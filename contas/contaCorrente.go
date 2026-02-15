package contas

import (
	"alura/bank-project/clientes"
	"fmt"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo -= valorDoSaque
		return fmt.Sprintf("Saque de R$%.2f realizado com sucesso. saldo atual: R$%.2f", valorDoSaque, c.saldo)
	} else {
		return fmt.Sprintf("saldo insuficiente para realizar o saque de R$%.2f. saldo atual: R$%.2f", valorDoSaque, c.saldo)
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return fmt.Sprintf("Depósito de R$%.2f realizado com sucesso. saldo atual: R$%.2f", valorDoDeposito, c.saldo)
	} else {
		return fmt.Sprintf("Valor do depósito inválido: R$%.2f", valorDoDeposito)
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia > 0 && valorDaTransferencia <= c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

package main

import (
	"fmt"
	"github.com/alura/banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoAlexandre := contas.ContaPoupanca{}
	contaDoAlexandre.Depositar(100)
	PagarBoleto(&contaDoAlexandre, 60)

	fmt.Println(contaDoAlexandre.ObterSaldo())

	contaDaLuiza := contas.ContaCorrente{}
	contaDaLuiza.Depositar(500)
	PagarBoleto(&contaDaLuiza, 400)

	fmt.Println(contaDaLuiza.ObterSaldo())

}

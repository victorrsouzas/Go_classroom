package main

import (
	"fmt"

	"bancos/clientes"
	"bancos/contas"
)

/* func SemParametro() string {
	return "Exemplo de função sem parâmetro!"
}

func UmParametro(texto string) string {
	return texto
}

func DoisParametros(texto string, numero int) (string, int) {
	return texto, numero
}

// Função Variádica
func Somando(numeros ...int) int {
	resultadoDaSoma := 0
	for _, numero := range numeros {
		resultadoDaSoma += numero
	}
	return resultadoDaSoma
} */

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	fmt.Println(contas.ContaCorrente{})
	contaDoVictor := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Victor",
			CPF:       "12312371289",
			Profissao: "Educador",
		},
		NumeroAgencia: 589,
		NumeroConta:   12356,
	}
	clienteAline := clientes.Titular{
		Nome:      "Aline",
		CPF:       "12312371289",
		Profissao: "Educador",
	}
	contaDaAline := contas.ContaCorrente{
		Titular:       clienteAline,
		NumeroAgencia: 589,
		NumeroConta:   12356,
	}

	fmt.Println(contaDoVictor, contaDaAline)

	fmt.Println(contaDoVictor == contaDaAline)

	contaDoMessi := &contas.ContaCorrente{}
	contaDoMessi2 := &contas.ContaCorrente{}

	contaDoMessi = new(contas.ContaCorrente)
	contaDoMessi.Titular = clientes.Titular{
		Nome:      "Messi",
		CPF:       "12312371289",
		Profissao: "Educador",
	}

	contaDoMessi2 = new(contas.ContaCorrente)
	contaDoMessi2.Titular = clientes.Titular{
		Nome:      "Messi",
		CPF:       "12312371289",
		Profissao: "Educador",
	}

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)

	fmt.Println(*contaDoMessi)

	fmt.Println(*contaDoMessi == *contaDoMessi2)

	fmt.Println(contaDaAline.Sacar(400))

	contaDaAline.Depositar(1000)

	status := contaDaAline.Transferir(300, contaDoMessi)

	fmt.Println(status)
	fmt.Println(contaDaAline.ObterSaldo())
	fmt.Println(contaDoMessi.ObterSaldo())
	fmt.Println(contaDoDenis.ObterSaldo())

	PagarBoleto(&contaDoDenis, 800)
	fmt.Println(contaDoDenis.ObterSaldo())
}

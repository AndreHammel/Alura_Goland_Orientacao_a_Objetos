package main

import (
	"fmt"

	"../pkg/contas"
)

func main() {
	contaDoBarriga := contas.ContaCorrente{Titular: "Sr Barriga", Saldo: 1000}
	contaDoMadruga := contas.ContaCorrente{Titular: "Sr Madruga", Saldo: 1}

	// utilizando o método sacar
	valorDoSaque := 200.
	fmt.Println(contaDoBarriga.Sacar(float32(valorDoSaque)))
	fmt.Println(contaDoBarriga)

	// utilizando o método depositar
	var valorDoDeposito float32 = 110
	status, valor := contaDoBarriga.Depositar(float32(valorDoDeposito))
	fmt.Println("Status da operação:",status,"Saldo na conta:", valor)

	// utilizanod o método tranferir
	fmt.Println(contaDoBarriga)
	fmt.Println(contaDoBarriga.Transferir(300, &contaDoMadruga))
	fmt.Println(contaDoBarriga)
	fmt.Println(contaDoMadruga)

}
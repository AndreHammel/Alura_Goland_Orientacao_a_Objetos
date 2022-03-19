package main

import "fmt"

type ContaCorrente struct {
	titular 			string
	numeroAgencia int
	numeroConta 	int
	saldo 				float32
}

func (c *ContaCorrente) Sacar(valorDoSaque float32) string {
	// podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if valorDoSaque <= c.saldo && valorDoSaque > 0 {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}
	return "Saldo insuficiênte"
}

func main() {
	contaDoKiko := ContaCorrente{"Kiko", 10, 1001, 10.43}
	fmt.Println(contaDoKiko)

	contaDoBarriga := ContaCorrente{}
	contaDoBarriga.titular = "Sr. Barriga"
	contaDoBarriga.saldo = 1000
	valorDoSaque := 200.
	fmt.Println(contaDoBarriga)
	fmt.Println(contaDoBarriga.Sacar(float32(valorDoSaque)))
	fmt.Println(contaDoBarriga)

	/*
	condaDoChaves := ContaCorrente{"Chaves",11, 2002, 20.32 }
	
	var contaDoMadruga *ContaCorrente
	contaDoMadruga = new(ContaCorrente)
	contaDoMadruga.titular = "Sr. Madruga"
	contaDoMadruga.saldo = 0.0

	fmt.Println(condaDoChaves)
	fmt.Println(*contaDoMadruga)
	*/
}
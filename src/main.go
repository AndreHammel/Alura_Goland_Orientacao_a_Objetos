package main

import "fmt"

type ContaCorrente struct {
	titular 			string
	numeroAgencia int
	numeroConta 	int
	saldo 				float32
}

// método sacar
func (c *ContaCorrente) Sacar(valorDoSaque float32) string {
	if valorDoSaque <= c.saldo && valorDoSaque > 0 {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}
	return "Saldo insuficiênte"
}

func main() {
	// forma 1 de instanciar (iniciando os atributos)
	contaDoKiko := ContaCorrente{"Kiko", 10, 1001, 10.43}
	fmt.Println(contaDoKiko)
	

	// forma 2 de instanciar 
	contaDoBarriga := ContaCorrente{}
	contaDoBarriga.titular = "Sr. Barriga"
	contaDoBarriga.saldo = 1000
	valorDoSaque := 200.
	fmt.Println(contaDoBarriga)
	fmt.Println("------------------------------------------------------")
	

	// forma 2 de instanciar (quando não inicio todos os atributos)
	condaDoChaves := ContaCorrente{titular: "Chaves", numeroAgencia: 9999 }
	fmt.Println(condaDoChaves)
	fmt.Println("------------------------------------------------------")
	


	// forma 3 de instanciar - não usual ( forma "comum" em outras linguagens)
	var contaDoMadruga *ContaCorrente
	contaDoMadruga = new(ContaCorrente)
	// contaDoMadruga.titular = "Sr. Madruga"
	contaDoMadruga.saldo = 0.0
	fmt.Println(*contaDoMadruga)
	fmt.Println("------------------------------------------------------")

	

	fmt.Println("----------------Métodos---------------------")
	// utilizando o método sacar
	fmt.Println(contaDoBarriga.Sacar(float32(valorDoSaque)))
	fmt.Println(contaDoBarriga)
}
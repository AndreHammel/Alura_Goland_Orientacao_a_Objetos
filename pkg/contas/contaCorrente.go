package contas

// atributos em minúsculas
type ContaCorrente struct {
	Titular 			string
	NumeroAgencia int
	NumeroConta 	int
	Saldo 				float32
}

// método sacar
func (c *ContaCorrente) Sacar(valorDoSaque float32) string {
	if valorDoSaque <= c.Saldo && valorDoSaque > 0 {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}
	return "Saldo insuficiênte"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float32) (string, float32) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.Saldo
	}
	return "Valor do deposito inválido", c.Saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float32, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.Saldo {
		c.Sacar(valorDaTransferencia)
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}
	
	return false
}

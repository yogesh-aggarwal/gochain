package transaction

import "fmt"

type Transaction struct {
	amount int64  // number of coins
	payer  string // public key
	payee  string // private key
}

func NewTransaction(amount int64, payer string, payee string) *Transaction {
	return &Transaction{
		amount: amount,
		payer:  payer,
		payee:  payee,
	}
}

func (t *Transaction) String() string {
	return fmt.Sprintf("%s-[%d]->%s", t.payer, t.amount, t.payee)
}

func (t *Transaction) Marshal() []byte {
	return []byte(t.String())
}

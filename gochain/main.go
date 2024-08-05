package main

import (
	"gochain/gochain/chain"
	"gochain/gochain/wallet"
)

func main() {
	c := chain.NewChain()

	yogesh := wallet.NewWallet(400)
	steve := wallet.NewWallet(200)

	yogesh.SendMoney(c, 200, steve.PublicKeyStr)
	yogesh.SendMoney(c, 200, steve.PublicKeyStr)

	c.Print()
}

package chain

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"gochain/gochain/block"
	"gochain/gochain/transaction"
)

type Chain struct {
	length int64
	blocks []block.Block
}

func NewChain() *Chain {
	chain := Chain{
		length: 0,
		blocks: make([]block.Block, 10),
	}

	genesis := block.NewBlock("", *transaction.NewTransaction(100, "genesis", "me"))
	chain.blocks[0] = *genesis
	chain.length = 1

	return &chain
}

func (c *Chain) GetLastBlock() *block.Block {
	return &c.blocks[c.length-1]
}

func (c *Chain) AddBlock(transaction transaction.Transaction, senderPublicKey rsa.PublicKey, signature []byte) {
	// Verify the signature
	hash := sha256.Sum256(transaction.Marshal())
	err := rsa.VerifyPKCS1v15(&senderPublicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		panic(err)
	}

	// Prepare the new block & add it to the chain
	newBlock := block.NewBlock(c.GetLastBlock().GetHash(), transaction)
	c.blocks[c.length] = *newBlock
	c.length++
}

func (c *Chain) Print() {
	for i := range c.length {
		print(c.blocks[i].GetHash())

		if i < c.length-1 {
			print("\n👇\n")
		}
	}
}

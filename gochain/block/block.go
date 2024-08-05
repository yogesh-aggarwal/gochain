package block

import (
	"crypto/sha256"
	"fmt"
	"gochain/gochain/transaction"
	"time"
)

type Block struct {
	prevHash    string
	transaction transaction.Transaction
	timestamp   int64
}

func NewBlock(prevHash string, transaction transaction.Transaction) *Block {
	return &Block{
		prevHash:    prevHash,
		transaction: transaction,
		timestamp:   time.Now().UnixNano(),
	}
}

func (b *Block) String() string {
	return fmt.Sprintf("%s ! %s @ %d", b.prevHash, b.transaction.String(), b.timestamp)
}

func (b *Block) GetHash() string {
	blockString := b.String()

	hash := sha256.New()
	_, err := hash.Write([]byte(blockString))
	if err != nil {
		panic(err)
	}

	hashStr := fmt.Sprintf("%x", hash.Sum(nil))
	return hashStr
}

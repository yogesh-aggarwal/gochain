package wallet

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"gochain/gochain/chain"
	"gochain/gochain/transaction"
)

type Wallet struct {
	Balance int64

	PrivateKey    rsa.PrivateKey
	PrivateKeyStr string

	PublicKey    rsa.PublicKey
	PublicKeyStr string
}

func NewWallet(balance int64) *Wallet {
	// Generate Private Key
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	// Generate Public Key
	publicKey := privateKey.Public().(*rsa.PublicKey)
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})

	return &Wallet{
		Balance: balance,

		PrivateKey:    *privateKey,
		PrivateKeyStr: fmt.Sprintf("%x", privateKeyPEM),

		PublicKey:    *publicKey,
		PublicKeyStr: fmt.Sprintf("%x", publicKeyPEM),
	}
}

func (w *Wallet) SendMoney(chain *chain.Chain, amount int64, payeePublicKeyStr string) {
	if amount > w.Balance {
		panic("BRO EARN SOME MONEY FIRST")
	}
	w.Balance -= amount

	// Prepare transaction
	transaction := transaction.NewTransaction(amount, w.PublicKeyStr, payeePublicKeyStr)

	// Prepare signature
	hash := sha256.Sum256(transaction.Marshal())
	signature, err := rsa.SignPKCS1v15(rand.Reader, &w.PrivateKey, crypto.SHA256, hash[:])
	if err != nil {
		panic(err)
	}

	// Add block to the chain
	chain.AddBlock(*transaction, w.PublicKey, signature)
}

func (w *Wallet) Print() {
	println("--- PUBLIC KEY ---")
	println(w.PublicKeyStr)
	println()
	println("--- PRIVATE KEY ---")
	println(w.PrivateKeyStr)
}

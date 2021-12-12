package wallet

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"errors"

	"github.com/Sandy247/blockchain-db/internal/transaction"
	"github.com/Sandy247/blockchain-db/pkg/utils"
)

type Wallet struct {
	KeyPair *rsa.PrivateKey
}

func NewWallet() (*Wallet, error) {
	key_pair, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, errors.New("Error generating key pair: " + err.Error())
	}
	return &Wallet{
		KeyPair: key_pair,
	}, nil
}

func (w *Wallet) Sign(data interface{}) (string, error) {
	data_hash := utils.CalculateHash(data)
	sign, err := rsa.SignPKCS1v15(rand.Reader, w.KeyPair, crypto.SHA512, data_hash)
	if err != nil {
		return "", errors.New("Error signing data: " + err.Error())
	}
	return hex.EncodeToString(sign), nil
}

func (w *Wallet) CreateTransaction(receiver_public_key rsa.PublicKey, amount int64, type_ string) (*transaction.Transaction, error) {
	sender_public_key := w.KeyPair.PublicKey
	transaction := transaction.NewTransaction(sender_public_key, receiver_public_key, amount, type_)
	signature, err := w.Sign(*transaction)
	if err != nil {
		return nil, errors.New("Error signing transaction: " + err.Error())
	}
	transaction.SignTransaction(signature)
	return transaction, nil
}

func ValidateSignature(data interface{}, signature string, public_key rsa.PublicKey) error {
	signature_bytes, err := hex.DecodeString(signature)
	if err != nil {
		return errors.New("Error decoding signature: " + err.Error())
	}
	data_hash := utils.CalculateHash(data)
	err = rsa.VerifyPKCS1v15(&public_key, crypto.SHA512, data_hash, signature_bytes)
	if err != nil {
		return errors.New("Error validating signature: " + err.Error())
	}
	return nil
}

package transaction

import (
	"crypto/rsa"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                  string
	Timestamp           int64
	Sender_Public_Key   rsa.PublicKey
	Receiver_Public_Key rsa.PublicKey
	Amount              int64
	Type                string
	Signature           string
}

func NewTransaction(sender_public_key, receiver_public_key rsa.PublicKey, amount int64, type_ string) *Transaction {
	return &Transaction{
		ID:                  uuid.New().String(),
		Timestamp:           time.Now().Unix(),
		Amount:              amount,
		Sender_Public_Key:   sender_public_key,
		Receiver_Public_Key: receiver_public_key,
		Type:                type_,
		Signature:           "",
	}
}

func (t *Transaction) SignTransaction(signature string) {
	t.Signature = signature
}

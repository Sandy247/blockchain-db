package transaction

import (
	"errors"
)

type TransactionPool struct {
	transactions map[*Transaction]struct{}
}

func (t *TransactionPool) AddTransaction(transaction *Transaction) error {
	exists := struct{}{}
	val, ok := t.transactions[transaction]
	if ok {
		if val == exists {
			return errors.New("Transaction already exists")
		} else {
			return errors.New("Transaction already exists but is malformed")
		}
	}
	t.transactions[transaction] = exists
	return nil
}

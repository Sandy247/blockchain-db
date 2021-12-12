package blockchain

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type BlockChain struct {
	chain []*Block
}

func NewBlockChain() *BlockChain {
	GenesisBlock := NewBlock("Let there be light", nil)
	return &BlockChain{
		chain: []*Block{GenesisBlock},
	}
}

func (bc *BlockChain) AddBlock(data interface{}) *Block {
	prevBlock := bc.chain[len(bc.chain)-1]
	newBlock := NewBlock(data, prevBlock)
	bc.chain = append(bc.chain, newBlock)
	return newBlock
}

func (bc *BlockChain) ReplaceBlockChain(newChain []*Block) error {
	if len(newChain) <= len(bc.chain) {
		return errors.New("the new chain must be longer")
	}
	if err := IsValidBlockChain(newChain); err != nil {
		return errors.New("the new chain is not valid")
	}
	bc.chain = newChain
	return nil
}

func (bc *BlockChain) String() string {
	var result strings.Builder
	for _, block := range bc.chain {
		result.WriteString(fmt.Sprintf("%s,", block))
	}
	return result.String()[:len(result.String())-1]
}

func (bc *BlockChain) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Chain []*Block `json:"chain"`
	}{
		Chain: bc.chain,
	})
}

func IsValidBlockChain(chain []*Block) error {
	if len(chain) == 0 {
		return errors.New("blockchain is empty")
	}
	for i := 1; i < len(chain); i++ {
		if v := IsValidBlock(chain[i], chain[i-1]); v != nil {
			return errors.New(v.Error() + " at index " + fmt.Sprintf("%d. Block: \n%s", i, chain[i]))
		}
	}
	return nil
}

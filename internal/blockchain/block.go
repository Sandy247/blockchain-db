package blockchain

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/Sandy247/blockchain-db/configs"
	"github.com/Sandy247/blockchain-db/pkg/utils"
)

type Block struct {
	ts            int64
	data          []byte
	prevBlockHash []byte
	difficulty    int64
	nonce         int64
	hash          []byte
}

func (b *Block) String() string {
	return fmt.Sprintf("[%d, %s, %s, %d, %d, %s]", b.ts, b.data, utils.ByteArrayToBinary(b.prevBlockHash), b.difficulty, b.nonce, utils.ByteArrayToBinary(b.hash))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Timestamp     int64  `json:"timestamp"`
		Data          string `json:"data"`
		PrevBlockHash string `json:"prev_block_hash"`
		Difficulty    int64  `json:"difficulty"`
		Nonce         int64  `json:"nonce"`
		Hash          string `json:"hash"`
	}{
		Timestamp:     b.ts,
		Data:          string(b.data),
		PrevBlockHash: utils.ByteArrayToBinary(b.prevBlockHash),
		Difficulty:    b.difficulty,
		Nonce:         b.nonce,
		Hash:          utils.ByteArrayToBinary(b.hash),
	})
}

func NewBlock(data_in interface{}, prevBlock *Block) *Block {
	var hash, prevBlockHash []byte
	// var data bytes.Buffer
	data := []byte(fmt.Sprintf("%v", data_in))
	ts := time.Now().UnixNano()
	nonce := int64(0)
	difficulty := adjustDifficulty(ts, prevBlock)
	// enc := gob.NewEncoder(&data)
	// err := enc.Encode(data_in)
	if prevBlock == nil {
		prevBlockHash = []byte{}
		return &Block{ts: ts, data: data, prevBlockHash: prevBlockHash, difficulty: difficulty, nonce: nonce, hash: utils.CalculateHash(ts, prevBlockHash, data, difficulty, nonce)}
	}

	for hash = utils.CalculateHash(ts, prevBlock.hash, data, difficulty, nonce); utils.ByteArrayToBinary((hash))[:difficulty] != strings.Repeat("0", int(difficulty)); {
		ts = time.Now().UnixNano()
		difficulty = adjustDifficulty(ts, prevBlock)
		nonce++
		hash = utils.CalculateHash(ts, prevBlock.hash, data, difficulty, nonce)
	}

	return &Block{ts: ts, data: data, prevBlockHash: prevBlock.hash, difficulty: difficulty, nonce: nonce, hash: hash}
}

func adjustDifficulty(ts int64, prevBlock *Block) int64 {
	if prevBlock == nil {
		return 1
	}
	if ts-prevBlock.ts < configs.MiningRate {
		return prevBlock.difficulty + 1
	}
	if prevBlock.difficulty > 1 {
		return prevBlock.difficulty - 1
	}
	return 1
}

func IsValidBlock(newBlock, prevBlock *Block) error {
	if !bytes.Equal(prevBlock.hash, newBlock.prevBlockHash) {
		return errors.New("previous hash is not equal")
	}
	if !bytes.Equal(utils.CalculateHash(newBlock.ts, prevBlock.hash, newBlock.data, newBlock.difficulty, newBlock.nonce), newBlock.hash) {
		return errors.New("hash is corrupt")
	}
	if utils.ByteArrayToBinary((newBlock.hash))[:newBlock.difficulty] != strings.Repeat("0", int(newBlock.difficulty)) {
		return errors.New("hash does not have expected number of leading zeroes")
	}
	if math.Abs(float64(newBlock.difficulty)-float64(prevBlock.difficulty)) > 1 {
		return errors.New("difficulty deviates from the previous difficulty by more than expected")
	}
	return nil
}

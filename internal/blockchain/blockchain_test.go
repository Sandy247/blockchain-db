package blockchain

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/Sandy247/blockchain-db/configs"
)

func TestNewBlockChain(t *testing.T) {
	bc := NewBlockChain()
	if string(bc.chain[0].data) != "Let there be light" {
		t.Errorf("NewBlockChain() failed. Got %s", bc.chain[0].data)
	}

}

func TestAddBlock(t *testing.T) {
	bc := NewBlockChain()
	bc.AddBlock("foo")
	if !bytes.Equal(bc.chain[len(bc.chain)-1].prevBlockHash, bc.chain[len(bc.chain)-2].hash) {
		t.Errorf("AddBlock() failed. prevHash of last block is not equal to the hash of the second last block!")
	}
}

func TestQuicklyMinedBlock(t *testing.T) {
	bc := NewBlockChain()
	bc.AddBlock("foo")
	if bc.chain[len(bc.chain)-1].difficulty != bc.chain[len(bc.chain)-2].difficulty+1 {
		t.Errorf("Difficulty isn't increasing for quickly mined blocks!")
	}
}

func TestSlowlyMinedBlock(t *testing.T) {
	bc := NewBlockChain()
	bc.AddBlock("foo")
	time.Sleep(configs.MiningRate * time.Nanosecond)
	bc.AddBlock("bar")
	if bc.chain[len(bc.chain)-1].difficulty != bc.chain[len(bc.chain)-2].difficulty-1 {
		t.Errorf("Difficulty isn't increasing for quickly mined blocks!")
	}
}

func TestValidBlockChain(t *testing.T) {
	bc := NewBlockChain()
	bc.AddBlock("foo")
	bc.AddBlock("bar")
	if err := IsValidBlockChain(bc.chain); err != nil {
		t.Errorf("Blockchain validity check isn't functioning properly! " + err.Error())
	}
}
func TestInvalidBlockChain(t *testing.T) {
	bc := NewBlockChain()
	bc.AddBlock("foo")
	bc.AddBlock("bar")
	bc.chain[1].hash = []byte("badhash")
	if err := IsValidBlockChain(bc.chain); err == nil {
		t.Errorf("Blockchain validity check isn't functioning properly!")
	}
}
func TestValidReplaceBlockChain(t *testing.T) {
	bc := NewBlockChain()
	bc.AddBlock("foo")
	bc.AddBlock("bar")
	bc_new := NewBlockChain()
	bc_new.AddBlock("foo")
	bc_new.AddBlock("bar")
	bc_new.AddBlock("baz")
	if err := bc.ReplaceBlockChain(bc_new.chain); err != nil || !reflect.DeepEqual(bc.chain, bc_new.chain) {
		t.Errorf("Valid Blockchain replacement failed! " + err.Error())
	}
}
func TestInvalidBlockChainReplaceShorterChain(t *testing.T) {
	bc := NewBlockChain()
	bc.AddBlock("foo")
	bc.AddBlock("bar")
	bc.AddBlock("baz")
	bc_new := NewBlockChain()
	bc_new.AddBlock("foo")
	bc_new.AddBlock("bar")
	if err := bc.ReplaceBlockChain(bc_new.chain); err == nil || err.Error() != "the new chain must be longer" {
		t.Errorf("Invalid Blockchain replacement didn't fail!")
	}
}
func TestInvalidBlockChainReplaceInvalidChain(t *testing.T) {
	bc := NewBlockChain()
	bc.AddBlock("foo")
	bc.AddBlock("bar")
	bc_new := NewBlockChain()
	bc_new.AddBlock("foo")
	bc_new.AddBlock("bar")
	bc_new.AddBlock("baz")
	bc_new.chain[2].hash = []byte("badhash")
	if err := bc.ReplaceBlockChain(bc_new.chain); err == nil || err.Error() != "the new chain is not valid" {
		t.Errorf("Invalid Blockchain replacement didn't fail!")
	}
}

func BenchmarkMiningRate(b *testing.B) {
	bc := NewBlockChain()
	for i := 0; i < b.N; i++ {
		bc.AddBlock("benchmark")
		fmt.Printf("Last difficulty: %d\n", bc.chain[len(bc.chain)-1].difficulty)
	}
}

package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	ts := int64(1636096342536626052)
	data := []byte{102, 111, 111}
	prevBlockHash := []byte{36, 165, 21, 158, 9, 83, 105, 193, 82, 94, 81, 7, 196, 13, 128, 218, 182, 161, 205, 50, 58, 243, 219, 150, 35, 216, 56, 129, 169, 178, 115, 16, 152, 95, 124, 122, 31, 254, 83, 181, 163, 189, 126, 147, 203, 20, 146, 131, 96, 188, 156, 51, 222, 201, 169, 8, 0, 168, 202, 246, 236, 41, 134, 254}
	difficulty := int64(2)
	nonce := int64(7)
	if !bytes.Equal(CalculateHash(ts, prevBlockHash, data, difficulty, nonce), CalculateHash(ts, prevBlockHash, data, difficulty, nonce)) {
		t.Error("CalculateHash() is not working")
	}
	if !bytes.Equal(CalculateHash("foo"), []byte{247, 251, 186, 110, 6, 54, 248, 144, 229, 111, 187, 243, 40, 62, 82, 76, 111, 163, 32, 74, 226, 152, 56, 45, 98, 71, 65, 208, 220, 102, 56, 50, 110, 40, 44, 65, 190, 94, 66, 84, 216, 130, 7, 114, 197, 81, 138, 44, 90, 140, 12, 127, 126, 218, 25, 89, 74, 126, 181, 57, 69, 62, 30, 215}) {
		t.Error("CalculateHash() is not working")
	}
}

func TestByteArrayToBinary(t *testing.T) {
	test := make([]byte, 4)
	binary.BigEndian.PutUint32(test, uint32(451))
	if res, err := strconv.ParseInt(ByteArrayToBinary(test), 2, 32); err != nil || res != 451 {
		if err != nil {
			t.Error("ByteArrayToBinary() is not working : " + err.Error())
		}
		t.Error("ByteArrayToBinary() is not working: " + fmt.Sprintf("%d", res))
	}
}

package utils

import (
	"crypto/sha512"
	"fmt"
	"strings"
)

func CalculateHash(data ...interface{}) []byte {
	var block_bytes []byte
	for _, d := range data {
		d_bytes := []byte(fmt.Sprintf("%v", d))
		block_bytes = append(block_bytes, d_bytes...)
	}
	hash := sha512.Sum512(block_bytes)
	return hash[:]
}

func ByteArrayToBinary(b []byte) string {
	var result strings.Builder
	for _, v := range b {
		result.WriteString(fmt.Sprintf("%08b", v))
	}
	return result.String()
}

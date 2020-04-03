package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

const RandLen = 32

// 强随机数，字符串长度64
func Rand32() (string, error) {
	return Rand(RandLen)
}

// 强随机数，字符串长度 = size * 2
func Rand(size int) (string, error) {
	data := make([]byte, size)
	n, err := rand.Read(data)
	if err != nil {
		return "", err
	}

	if n != size {
		return "", fmt.Errorf("n %d; size %d; don't match", n, size)
	}

	return hex.EncodeToString(data), nil
}

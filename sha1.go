package crypto

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
)

func Sha1(s string) string {
	return Sha1Sum([]byte(s))
}

func Sha1File(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err == nil {
		return Sha1Sum(data), nil
	} else {
		return "", err
	}
}

func Sha1Sum(data []byte) string {
	r := sha1.Sum(data)
	return hex.EncodeToString(r[:])
}

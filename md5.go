package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
)

func Md5(s string) string {
	return Md5Sum([]byte(s))
}

func Md5File(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err == nil {
		return Md5Sum(data), nil
	} else {
		return "", err
	}
}

func Md5Sum(data []byte) string {
	r := md5.Sum(data)
	return hex.EncodeToString(r[:])
}

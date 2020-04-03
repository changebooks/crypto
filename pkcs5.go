package crypto

import (
	"bytes"
	"errors"
)

func Pkcs5Padding(data []byte, blockSize int) ([]byte, error) {
	dataSize := len(data)
	if dataSize == 0 {
		return nil, errors.New("data can't be empty")
	}

	if blockSize <= 0 {
		return nil, errors.New("block size can't be less or equal than 0")
	}

	paddingSize := blockSize - dataSize%blockSize
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(data, paddingText...), nil
}

func Pkcs5UnPadding(data []byte) ([]byte, error) {
	dataSize := len(data)
	if dataSize == 0 {
		return nil, errors.New("data can't be empty")
	}

	paddingSize := int(data[dataSize-1])
	contentSize := dataSize - paddingSize
	if contentSize < 0 {
		return nil, errors.New("data size can't be less than padding size")
	}

	return data[:contentSize], nil
}

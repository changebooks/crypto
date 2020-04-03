package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"sync"
)

type AesCbc struct {
	key     []byte // 长度，8的倍数，如：16, 24, 32
	iv      []byte // 长度，16
	block   cipher.Block
	encrypt cipher.BlockMode
	decrypt cipher.BlockMode
}

// 加密
func (x *AesCbc) Encrypt(plainText []byte) ([]byte, error) {
	if plainText == nil {
		return nil, errors.New("plain text can't be nil")
	}

	plainText, err := Pkcs5Padding(plainText, x.block.BlockSize())
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, len(plainText))
	x.encrypt.CryptBlocks(cipherText, plainText)
	return cipherText, nil
}

// 解密
func (x *AesCbc) Decrypt(cipherText []byte) ([]byte, error) {
	if cipherText == nil {
		return nil, errors.New("cipher text can't be nil")
	}

	plainText := make([]byte, len(cipherText))
	x.decrypt.CryptBlocks(plainText, cipherText)
	return Pkcs5UnPadding(plainText)
}

func (x *AesCbc) GetKey() []byte {
	return x.key
}

func (x *AesCbc) GetIv() []byte {
	return x.iv
}

func (x *AesCbc) GetBlock() cipher.Block {
	return x.block
}

func (x *AesCbc) GetEncrypt() cipher.BlockMode {
	return x.encrypt
}

func (x *AesCbc) GetDecrypt() cipher.BlockMode {
	return x.decrypt
}

type AesCbcBuilder struct {
	mu  sync.Mutex // protects following fields
	key []byte
	iv  []byte
}

func (x *AesCbcBuilder) Build() (*AesCbc, error) {
	if x.key == nil {
		return nil, errors.New("key can't be nil")
	}

	if x.iv == nil {
		return nil, errors.New("iv can't be nil")
	}

	block, err := aes.NewCipher(x.key)
	if err != nil {
		return nil, err
	}

	if len(x.iv) != block.BlockSize() {
		return nil, fmt.Errorf("iv size %d; block size %d; not equal", len(x.iv), block.BlockSize())
	}

	encrypt := cipher.NewCBCEncrypter(block, x.iv)
	decrypt := cipher.NewCBCDecrypter(block, x.iv)

	return &AesCbc{
		key:     x.key,
		iv:      x.iv,
		block:   block,
		encrypt: encrypt,
		decrypt: decrypt,
	}, nil
}

func (x *AesCbcBuilder) SetKeyStr(s string) *AesCbcBuilder {
	return x.SetKey([]byte(s))
}

func (x *AesCbcBuilder) SetKey(bucket []byte) *AesCbcBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.key = bucket
	return x
}

func (x *AesCbcBuilder) SetIvStr(s string) *AesCbcBuilder {
	return x.SetIv([]byte(s))
}

func (x *AesCbcBuilder) SetIv(bucket []byte) *AesCbcBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.iv = bucket
	return x
}

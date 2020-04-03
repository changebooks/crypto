package crypto

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type Cipher struct {
	key *Key
}

func NewCipher(k *Key) (*Cipher, error) {
	if k == nil {
		return nil, errors.New("key can't be nil")
	}

	return &Cipher{
		key: k,
	}, nil
}

// 加密
func (x *Cipher) Encrypt(s string, period int64) string {
	randKey := x.key.GetRandom()
	encryptKey := x.GetEncryptKey(randKey)
	signatureKey := x.GetSignatureKey(randKey)
	expiredAt := fmt.Sprintf("%010d", x.key.GetExpiredAt(period))

	plainText := expiredAt + x.Signature(s, signatureKey) + s
	cipherText := x.Calculate(plainText, encryptKey)
	return randKey + Base64UrlSafeEncode([]byte(cipherText))
}

// 解密
func (x *Cipher) Decrypt(s string) (plainText string, err error, isExpired bool, expiredAt int64) {
	randLen := x.key.GetRandLen()
	if len(s) <= randLen {
		err = errors.New("cipher text len can't be less or equal than rand key len")
		return
	}

	cipherText, err := Base64UrlSafeDecode(s[randLen:])
	if err != nil {
		return
	}

	randKey := s[:randLen]
	encryptKey := x.GetEncryptKey(randKey)
	signatureKey := x.GetSignatureKey(randKey)

	s = x.Calculate(string(cipherText), encryptKey)
	if len(s) <= 10 {
		err = errors.New("contain expiry's plain text len can't be less or equal than 10")
		return
	}

	expiredAt, err = strconv.ParseInt(s[:10], 10, 64)
	if err != nil {
		return
	}

	if expiredAt > 0 && expiredAt <= x.key.Now() {
		isExpired = true
	}

	if len(s) <= 26 {
		err = errors.New("contain signature's plain text len can't be less or equal than 26")
		return
	}

	if s[10:26] != x.Signature(s[26:], signatureKey) {
		err = errors.New("signature differ")
		return
	}

	plainText = s[26:]
	return
}

// 异或
func (x *Cipher) Calculate(s string, encryptKey string) string {
	txtLen := len(s)
	if txtLen == 0 {
		return ""
	}

	var r bytes.Buffer
	iv := x.InitialVector(encryptKey)
	for i, j, k := 0, 0, 0; i < txtLen; i++ {
		j = (j + 1) % 256
		k = (k + iv[j]) % 256
		iv[j], iv[k] = iv[k], iv[j]
		r.WriteByte(s[i] ^ uint8(iv[(iv[j]+iv[k])%256]))
	}

	return r.String()
}

// 签名
func (x *Cipher) Signature(s string, signatureKey string) string {
	return Md5(s + signatureKey)[:16]
}

// 通过加密密钥，初始化向量
func (x *Cipher) InitialVector(encryptKey string) [256]int {
	var r [256]int
	keyLen := len(encryptKey)
	if keyLen == 0 {
		return r
	}

	var keys [256]int
	for i := 0; i < 256; i++ {
		keys[i] = int(encryptKey[i%keyLen])
		r[i] = i
	}

	for i, j := 0, 0; i < 256; i++ {
		j = (j + r[i] + keys[i]) % 256
		r[i], r[j] = r[j], r[i]
	}

	return r
}

// 加密密钥
func (x *Cipher) GetEncryptKey(randKey string) string {
	encryptKey := Md5(x.key.GetEncrypt())
	randKey = Md5(randKey + encryptKey)
	return Md5(encryptKey[:16] + randKey)
}

// 签名密钥
func (x *Cipher) GetSignatureKey(randKey string) string {
	signatureKey := Md5(x.key.GetSignature())
	randKey = Md5(signatureKey + randKey)
	return Md5(randKey + signatureKey[16:])
}

func (x *Cipher) GetKey() *Key {
	return x.key
}

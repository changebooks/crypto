package crypto

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	MinRandKeyLen = 0
	MaxRandKeyLen = 32
)

// 密钥管理
type Key struct {
	encrypt   string         // 加密密钥
	signature string         // 签名密钥
	randLen   int            // 随机密钥长度，[MinRandKeyLen, MaxRandKeyLen]
	period    int64          // 有效期时长，单位：秒
	timezone  *time.Location // 时区
}

// 随机密钥
func (x *Key) GetRandom() string {
	if x.randLen > 0 {
		return Unique()[:x.randLen]
	} else {
		return ""
	}
}

// 到期时间
// < 0 ? use x.period
// > 0 ? use period
// = 0 ? return 0
func (x *Key) GetExpiredAt(period int64) int64 {
	if period < 0 {
		period = x.period
	}

	if period > 0 {
		return period + x.Now()
	} else {
		return 0
	}
}

func (x *Key) Now() int64 {
	return time.Now().In(x.timezone).Unix()
}

func (x *Key) GetEncrypt() string {
	return x.encrypt
}

func (x *Key) GetSignature() string {
	return x.signature
}

func (x *Key) GetRandLen() int {
	return x.randLen
}

func (x *Key) GetPeriod() int64 {
	return x.period
}

type KeyBuilder struct {
	mu        sync.Mutex // protects following fields
	encrypt   string
	signature string
	randLen   int
	period    int64
	timezone  *time.Location
}

func (x *KeyBuilder) Build() (*Key, error) {
	if x.encrypt == "" {
		return nil, errors.New("encrypt key can't be empty")
	}

	if x.signature == "" {
		return nil, errors.New("signature key can't be empty")
	}

	if x.randLen < MinRandKeyLen || x.randLen > MaxRandKeyLen {
		return nil, fmt.Errorf("rand len %d can't be less than %d or greater than %d", x.randLen, MinRandKeyLen, MaxRandKeyLen)
	}

	period := x.period
	if period < 0 {
		period = 0
	}

	timezone := x.timezone
	if timezone == nil {
		timezone = time.Local
	}

	return &Key{
		encrypt:   x.encrypt,
		signature: x.signature,
		randLen:   x.randLen,
		period:    period,
		timezone:  timezone,
	}, nil
}

func (x *KeyBuilder) SetEncrypt(s string) *KeyBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.encrypt = s
	return x
}

func (x *KeyBuilder) SetSignature(s string) *KeyBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.signature = s
	return x
}

func (x *KeyBuilder) SetRandLen(n int) *KeyBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.randLen = n
	return x
}

func (x *KeyBuilder) SetPeriod(n int64) *KeyBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.period = n
	return x
}

func (x *KeyBuilder) SetTimezone(loc *time.Location) *KeyBuilder {
	x.mu.Lock()
	defer x.mu.Unlock()

	x.timezone = loc
	return x
}

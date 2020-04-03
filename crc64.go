package crypto

import (
	"hash/crc64"
	"sync"
)

// ECMATable is the table for the ECMA polynomial, defined in ECMA 182.
var ecmaTable *crc64.Table = nil

var ecmaOnce sync.Once // guards ecmaInit

func Crc64(s string) uint64 {
	return Crc64Sum([]byte(s))
}

func Crc64Sum(data []byte) uint64 {
	return crc64.Checksum(data, ecmaInit())
}

func ecmaInit() *crc64.Table {
	ecmaOnce.Do(func() {
		ecmaTable = crc64.MakeTable(crc64.ECMA)
	})
	return ecmaTable
}

package crypto

import "hash/crc32"

func Crc32(s string) uint32 {
	return Crc32Sum([]byte(s))
}

func Crc32Sum(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

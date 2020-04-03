package crypto

import (
	"encoding/base64"
	"strings"
)

// 用于填充Base64长度的等于号
const EqualSign4 = "===="

// Base64加密，URL安全转义
// 去掉Base64密文中的=、+、/
func Base64UrlSafeEncode(b []byte) string {
	return Base64UrlSafeCover(base64.StdEncoding.EncodeToString(b))
}

// Base64解密，URL安全转义
// 还原Base64密文中的=、+、/
func Base64UrlSafeDecode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(Base64UrlSafeRestore(s))
}

// 去掉Base64密文中的=、+、/
// 去掉=、+ => -、/ => _
func Base64UrlSafeCover(s string) string {
	s = strings.Replace(s, "=", "", -1)
	s = strings.Replace(s, "+", "-", -1)
	s = strings.Replace(s, "/", "_", -1)
	return s
}

// 还原Base64密文中的=、+、/
// _ => /, - => +, 还原 =
func Base64UrlSafeRestore(s string) string {
	s = strings.Replace(s, "_", "/", -1)
	s = strings.Replace(s, "-", "+", -1)
	size := len(s) % len(EqualSign4)
	if size > 0 {
		s += EqualSign4[size:]
	}
	return s
}

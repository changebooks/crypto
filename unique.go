package crypto

import (
	"math/rand"
	"strconv"
	"time"
)

// 唯一码
func Unique() string {
	if r, err := Rand32(); err == nil {
		return r
	} else {
		return unique()
	}
}

func unique() string {
	n := rand.New(rand.NewSource(time.Now().UnixNano())).Uint64()
	t := time.Now().UnixNano()
	return Md5(strconv.FormatUint(n, 10)) + Md5(strconv.FormatInt(t, 10))
}

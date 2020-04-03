package crypto

func Hash(b []byte) int32 {
	if b == nil {
		return 0
	}

	size := len(b)
	if size == 0 {
		return 0
	}

	var r int32 = 0
	for i := 0; i < size; i++ {
		r = 31*r + int32(b[i])
	}

	return r
}

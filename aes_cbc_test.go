package crypto

import "testing"

func TestAesCbc(t *testing.T) {
	b := AesCbcBuilder{}
	a, _ := b.SetKeyStr("1234567890123456").SetIvStr("1234567890123456").Build()
	cipherText, _ := a.Encrypt([]byte("abc"))
	got, _ := a.Decrypt(cipherText)
	want := "abc"
	if string(got) != want {
		t.Errorf("got %q; want %q", got, want)
	}
}

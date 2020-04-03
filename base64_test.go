package crypto

import "testing"

func TestBase64Encode(t *testing.T) {
	got := Base64Encode(nil)
	want := ""
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := Base64Encode([]byte("123456"))
	want2 := "MTIzNDU2"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}
}

func TestBase64Decode(t *testing.T) {
	got, _ := Base64Decode("")
	want := ""
	if string(got) != want {
		t.Errorf("got %q; want %q", string(got), want)
	}

	got2, _ := Base64Decode("MTIzNDU2")
	want2 := "123456"
	if string(got2) != want2 {
		t.Errorf("got %q; want %q", string(got2), want2)
	}
}

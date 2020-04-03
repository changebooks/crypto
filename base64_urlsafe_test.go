package crypto

import "testing"

func TestBase64UrlSafeCover(t *testing.T) {
	got := Base64UrlSafeCover("")
	want := ""
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := Base64UrlSafeCover("=+/=+/=+/")
	want2 := "-_-_-_"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}
}

func TestBase64UrlSafeRestore(t *testing.T) {
	got := Base64UrlSafeRestore("")
	want := ""
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := Base64UrlSafeRestore("-_-_-_")
	want2 := "+/+/+/=="
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := Base64UrlSafeRestore("1")
	want3 := "1==="
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	got4 := Base64UrlSafeRestore("12")
	want4 := "12=="
	if got4 != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}

	got5 := Base64UrlSafeRestore("123")
	want5 := "123="
	if got5 != want5 {
		t.Errorf("got %q; want %q", got5, want5)
	}

	got6 := Base64UrlSafeRestore("1234")
	want6 := "1234"
	if got6 != want6 {
		t.Errorf("got %q; want %q", got6, want6)
	}

	got7 := Base64UrlSafeRestore("12345")
	want7 := "12345==="
	if got7 != want7 {
		t.Errorf("got %q; want %q", got7, want7)
	}
}

func TestBase64UrlSafeEncode(t *testing.T) {
	got := Base64UrlSafeEncode(nil)
	want := ""
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := Base64UrlSafeEncode([]byte("1234567890"))
	want2 := "MTIzNDU2Nzg5MA"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}
}

func TestBase64UrlSafeDecode(t *testing.T) {
	got, _ := Base64UrlSafeDecode("")
	want := ""
	if string(got) != want {
		t.Errorf("got %q; want %q", string(got), want)
	}

	got2, _ := Base64UrlSafeDecode("MTIzNDU2Nzg5MA")
	want2 := "1234567890"
	if string(got2) != want2 {
		t.Errorf("got %q; want %q", string(got2), want2)
	}
}

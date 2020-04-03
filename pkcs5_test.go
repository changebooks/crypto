package crypto

import "testing"

func TestPkcs5Padding(t *testing.T) {
	_, got := Pkcs5Padding(nil, 0)
	want := "data can't be empty"
	if got != nil && got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	_, got2 := Pkcs5Padding([]byte(""), 0)
	want2 := "data can't be empty"
	if got2 != nil && got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	_, got3 := Pkcs5Padding([]byte("abc"), -1)
	want3 := "block size can't be less or equal than 0"
	if got3 != nil && got3.Error() != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	_, got4 := Pkcs5Padding([]byte("abc"), 0)
	want4 := "block size can't be less or equal than 0"
	if got4 != nil && got4.Error() != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}

	got5, _ := Pkcs5Padding([]byte("a"), 1)
	want5 := []byte("a")
	want5 = append(want5, 1)
	if string(got5) != string(want5) {
		t.Errorf("got %q; want %q", string(got5), string(want5))
	}

	got6, _ := Pkcs5Padding([]byte("a"), 3)
	want6 := []byte("a")
	want6 = append(want6, 2, 2)
	if string(got6) != string(want6) {
		t.Errorf("got %q; want %q", string(got6), string(want6))
	}
}

func TestPkcs5UnPadding(t *testing.T) {
	_, got := Pkcs5UnPadding(nil)
	want := "data can't be empty"
	if got != nil && got.Error() != want {
		t.Errorf("got %q; want %q", got, want)
	}

	_, got2 := Pkcs5UnPadding([]byte(""))
	want2 := "data can't be empty"
	if got2 != nil && got2.Error() != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	data3 := append([]byte("abc"), 10, 10)
	_, got3 := Pkcs5UnPadding(data3)
	want3 := "data size can't be less than padding size"
	if got3 != nil && got3.Error() != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	want4 := []byte("a")
	data4 := append(want4, 2, 2)
	got4, _ := Pkcs5UnPadding(data4)
	if string(got4) != string(want4) {
		t.Errorf("got %q; want %q", string(got4), string(want4))
	}

	want5 := []byte("a")
	data5 := append(want5, 1)
	got5, _ := Pkcs5UnPadding(data5)
	if string(got5) != string(want5) {
		t.Errorf("got %q; want %q", string(got5), string(want5))
	}
}

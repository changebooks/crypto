package crypto

import "testing"

func TestHash(t *testing.T) {
	got := Hash(nil)
	var want int32 = 0
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}

	got2 := Hash([]byte(""))
	var want2 int32 = 0
	if got2 != want2 {
		t.Errorf("got %d; want %d", got2, want2)
	}

	got3 := Hash([]byte("a"))
	var want3 int32 = 97
	if got3 != want3 {
		t.Errorf("got %d; want %d", got3, want3)
	}

	got4 := Hash([]byte("1"))
	var want4 int32 = 49
	if got4 != want4 {
		t.Errorf("got %d; want %d", got4, want4)
	}

	got5 := Hash([]byte("ab"))
	var want5 int32 = 3105
	if got5 != want5 {
		t.Errorf("got %d; want %d", got5, want5)
	}

	got6 := Hash([]byte("123456"))
	var want6 int32 = 1450575459
	if got6 != want6 {
		t.Errorf("got %d; want %d", got6, want6)
	}

	got7 := Hash([]byte("abc123456def789abc123456def789abc123456def789abc123456def789abc123456def789abc123456def789"))
	var want7 int32 = -674046208
	if got7 != want7 {
		t.Errorf("got %d; want %d", got7, want7)
	}
}

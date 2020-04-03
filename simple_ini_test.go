package crypto

import "testing"

func TestSimpleIniParse(t *testing.T) {
	got := SimpleIniParse("a'b\"c d\\e=1#2\r3\n4\n#\na=1")
	want := map[string]string{"abcde": "1", "a": "1"}
	if len(got) != len(want) {
		t.Errorf("got %q; want %q", got, want)
	}
}

func TestIniParseLine(t *testing.T) {
	n, v := iniParseLine("a'b\"c d\\e=1#2\r3\n4")
	got := n + "=" + v
	want := "abcde=1"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	n, v = iniParseLine("#abc=123")
	got2 := n + "=" + v
	want2 := "="
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	n, v = iniParseLine("abc=1=23")
	got3 := n + "=" + v
	want3 := "abc=1=23"
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	n, v = iniParseLine("abc123")
	got4 := n + "=" + v
	want4 := "="
	if got4 != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}
}

func TestIniSanitiseName(t *testing.T) {
	got := iniSanitiseName("a'b\"c d\\e")
	want := "abcde"
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := iniSanitiseName("a#bc")
	want2 := "a"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := iniSanitiseName("ab\tcd")
	want3 := "ab"
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	got4 := iniSanitiseName("abc\nd")
	want4 := "abc"
	if got4 != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}
}

func TestIniSanitiseValue(t *testing.T) {
	got := iniSanitiseValue("")
	want := ""
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := iniSanitiseValue("abc")
	want2 := "abc"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := iniSanitiseValue("#abcd")
	want3 := ""
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	got4 := iniSanitiseValue("a#bcd")
	want4 := "a"
	if got4 != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}

	got5 := iniSanitiseValue("ab\rcd")
	want5 := "ab"
	if got5 != want5 {
		t.Errorf("got %q; want %q", got5, want5)
	}

	got6 := iniSanitiseValue("abc\nd")
	want6 := "abc"
	if got6 != want6 {
		t.Errorf("got %q; want %q", got6, want6)
	}
}

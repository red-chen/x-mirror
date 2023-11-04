package version

import "testing"

func TestMajorMinor(t *testing.T) {
	n := MajorMinor()
	if n != "1.7" {
		t.Error("Not equal: ", n)
	}
}

func TestFull(t *testing.T) {
	n := Full()
	if n != "2-1.7" {
		t.Error("Not equal: ", n)
	}
}

func TestCompat(t *testing.T) {
	{
		n := Compat("1.7", "1.7")
		if !n {
			t.Error("Not equal: ", n)
		}
	}
	{
		n := Compat("1.7", "1.8")
		if n {
			t.Error("Not equal: ", n)
		}
	}
}

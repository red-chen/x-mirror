package util

import (
	"testing"
)

func TestRandomSeed(t *testing.T) {
	_, err := RandomSeed()
	if err != nil {
		t.Error("Got err:", err)
	}
}

func TestRandId(t *testing.T) {
	lenght := 8
	id := RandId(lenght)
	if len(id) != lenght*2 {
		t.Error("Id len err")
	}
}

func TestSecureRandId(t *testing.T) {
	lenght := 8
	id, err := SecureRandId(lenght)
	if err != nil {
		t.Error("Got err:", err)
	}
	if len(id) != lenght*2 {
		t.Error("Id len err")
	}
}

func TestSecureRandIdOrPanic(t *testing.T) {
	lenght := 8
	id := SecureRandIdOrPanic(lenght)
	if len(id) != lenght*2 {
		t.Error("Id len err")
	}
}

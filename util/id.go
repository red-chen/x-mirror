package util

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	mrand "math/rand"
)

func RandomSeed() (seed int64, err error) {
	err = binary.Read(rand.Reader, binary.LittleEndian, &seed)
	return
}

func RandId(idLen int) string {
	if idLen <= 0 {
		panic(fmt.Sprintf("Input idLen(%d) must grater than 0", idLen))
	}
	b := make([]byte, idLen)
	var randVal uint32
	for i := 0; i < idLen; i++ {
		byteIdx := i % 4
		if byteIdx == 0 {
			randVal = mrand.Uint32()
		}
		b[i] = byte((randVal >> (8 * uint(byteIdx))) & 0xFF)
	}
	return fmt.Sprintf("%x", b)
}

func SecureRandId(idLen int) (id string, err error) {
	b := make([]byte, idLen)
	n, err := rand.Read(b)

	if n != idLen {
		err = fmt.Errorf("Only generated %d random bytes, %d requested", n, idLen)
		return
	}

	if err != nil {
		return
	}

	id = fmt.Sprintf("%x", b)
	return
}

func SecureRandIdOrPanic(idLen int) string {
	id, err := SecureRandId(idLen)
	if err != nil {
		panic(err)
	}
	return id
}

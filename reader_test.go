package bitstream

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBitReader(t *testing.T) {
	data := []byte{0x65} // 01100101

	testCases := []struct {
		readBit     bool
		count       int
		expectedArr []byte
		expected    uint64
	}{
		{
			readBit: true,
			count:   6,
		},
	}

	for _, test := range testCases {
		r := NewReader(bytes.NewBuffer(data))
		fmt.Printf("data: %b\n", data)
		var b byte
		var err error
		if test.readBit {
			b, _ := r.ReadBits(test.count)
			fmt.Printf("b: %b\n", b)
		} else {
			b, err = r.ReadByte()
		}
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("b: %b\n", b)
	}
	t.Fatal("OK")
}

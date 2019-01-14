package bitstream

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestBitReader(t *testing.T) {
	data := []byte{0xde, 0x71, 0x65} // 11011110 01110001 01100101

	testCases := []struct {
		funcName string
		count    int
		expected string
		err      error
	}{
		{
			funcName: "bits",
			count:    2,
			expected: "11",
		},
		{
			funcName: "byte",
			expected: "1111001",
		},
		{
			funcName: "bits",
			count:    13,
			expected: "1100010110010",
		},
		{
			funcName: "bit",
			expected: "1",
		},
		{
			funcName: "byte",
			expected: "0",
			err:      io.EOF,
		},
	}

	r := NewReader(bytes.NewBuffer(data))
	for _, test := range testCases {
		var a interface{}
		var err error
		switch test.funcName {
		case "bit":
			a, err = r.ReadBit()
		case "bits":
			a, err = r.ReadBits(test.count)
		case "byte":
			a, err = r.ReadByte()
		}
		b := fmt.Sprintf("%b", a)
		if b != test.expected {
			t.Errorf("expected %s, but found %s", test.expected, b)
		}
		if err != test.err {
			t.Errorf("error expected %s, but found %s", test.err, err)
		}
	}
}

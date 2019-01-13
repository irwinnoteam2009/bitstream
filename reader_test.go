package bitstream

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBitReader(t *testing.T) {
	data := []byte{0x65, 0x65, 0x65} // 01100101 01100101 01100101

	r := NewReader(bytes.NewBuffer(data))
	b, _ := r.ReadBits(2)
	fmt.Printf("2>>%b\n", b)
	ub, _ := r.ReadByte()
	fmt.Printf("8>>%b\n", ub)
	b, _ = r.ReadBits(13)
	fmt.Printf("13>>%b\n", b)
	bb, _ := r.ReadBit()
	fmt.Printf("1>>%b\n", bb)

	xb, err := r.ReadBits(4)
	fmt.Println("error?", err, xb)
	t.Fatal("OK")
}

package bitstream

import "io"

// Reader is a bit reader from io.Reader
type Reader struct {
	io.Reader
	data  [1]byte // current byte from stream
	count uint8   // readed bit count
}

// NewReader creates new bit reader from io.Reader
func NewReader(r io.Reader) *Reader {
	return &Reader{Reader: r}
}

// ReadBit reads bit from stream
func (r *Reader) ReadBit() (byte, error) {
	if r.count == 0 {
		if n, err := r.Read(r.data[:]); err != nil || n == 0 {
			return 0, err
		}
		r.count = 8
	}
	r.count--
	d := r.data[0] & 0x80 // get first bit. 0x80 = 10000000
	r.data[0] <<= 1
	return d, nil
}

// ReadByte reads byte from stream
func (r *Reader) ReadByte() (byte, error) {
	if r.count == 0 {
		n, err := r.Read(r.data[:])
		if err != nil || n == 0 {
			return 0, err
		}
		return r.data[0], err
	}

	// store old byte to variable and read new byte from stream
	byt := r.data[0]
	n, err := r.Read(r.data[:])
	if err != nil || n == 0 {
		return 0, err
	}
	// append byt by r.data bits and clear r.data bits
	byt |= r.data[0] >> r.count
	r.data[0] <<= (8 - r.count)

	return byt, nil
}

// ReadBits reads n bits from stream
func (r *Reader) ReadBits(n int) (uint64, error) {
	var u uint64

	for n >= 8 {
		byt, err := r.ReadByte()
		if err != nil {
			return 0, err
		}
		u = (u << 8) | uint64(byt)
		n -= 8
	}

	for n > 0 {
		byt, err := r.ReadBit()
		if err != nil {
			return 0, err
		}
		u <<= 1
		if byt == 1 {
			u |= 1
		}
		n--
	}

	return u, nil
}

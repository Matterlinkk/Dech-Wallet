package hash

import (
	"encoding/binary"
	"fmt"
)

var H = []uint32{
	0x67452301,
	0xefcdab89,
	0x98badcfe,
	0x10325476,
	0xc3d2e1f0,
}

func RIPEMD160(message []byte) string {

	l := len(message)
	message = append(message[:], []byte{0x80}...)

	for len(message)*8%512 != 448 {
		message = append(message, 0x00)
	}

	ms := make([]byte, 8)
	binary.LittleEndian.PutUint64(ms, uint64(l*8))
	message = append(message[:], ms[:]...)

	var blocks [][16]uint32

	messageLen := len(message) / 64

	for i := 0; i < messageLen; i++ {
		var block [16]uint32
		for j := 0; j < 16; j++ {
			block[j] = binary.LittleEndian.Uint32(message[i*64+j*4 : i*64+(j+1)*4])
		}
		blocks = append(blocks, block)
	}

	T := uint32(0)

	for _, m := range blocks {
		a := H[0]
		b := H[1]
		c := H[2]
		d := H[3]
		e := H[4]
		_a := H[0]
		_b := H[1]
		_c := H[2]
		_d := H[3]
		_e := H[4]
		for t := uint32(0); t < 80; t++ {
			T = rotate(a+f(t, b, c, d)+m[r[t]]+k(t), s[t]) + e
			a = e
			e = d
			d = rotate(c, 10)
			c = b
			b = T

			T = rotate(_a+f(79-t, _b, _c, _d)+m[_r[t]]+_k(t), _s[t]) + _e
			_a = _e
			_e = _d
			_d = rotate(_c, 10)
			_c = _b
			_b = T
		}
		T = H[1] + c + _d
		H[1] = H[2] + d + _e
		H[2] = H[3] + e + _a
		H[3] = H[4] + a + _b
		H[4] = H[0] + b + _c
		H[0] = T

	}
	var hash [20]byte
	for i, s := range H {
		hash[i*4] = byte(s)
		hash[i*4+1] = byte(s >> 8)
		hash[i*4+2] = byte(s >> 16)
		hash[i*4+3] = byte(s >> 24)
	}

	return fmt.Sprintf("%x", hash)

}

func k(j uint32) uint32 {
	if 0 <= j && j <= 15 {
		return 0x00000000
	} else if 16 <= j && j <= 31 {
		return 0x5A827999
	} else if 32 <= j && j <= 47 {
		return 0x6ED9EBA1
	} else if 48 <= j && j <= 63 {
		return 0x8F1BBCDC
	} else if 64 <= j && j <= 79 {
		return 0xA953FD4E
	}
	return 0
}

// _K ...
func _k(j uint32) uint32 {
	if 0 <= j && j <= 15 {
		return 0x50A28BE6
	} else if 16 <= j && j <= 31 {
		return 0x5C4DD124
	} else if 32 <= j && j <= 47 {
		return 0x6D703EF3
	} else if 48 <= j && j <= 63 {
		return 0x7A6D76E9
	} else if 64 <= j && j <= 79 {
		return 0x00000000
	}
	return 0
}

var r = [80]uint{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	7, 4, 13, 1, 10, 6, 15, 3, 12, 0, 9, 5, 2, 14, 11, 8,
	3, 10, 14, 4, 9, 15, 8, 1, 2, 7, 0, 6, 13, 11, 5, 12,
	1, 9, 11, 10, 0, 8, 12, 4, 13, 3, 7, 15, 14, 5, 6, 2,
	4, 0, 5, 9, 7, 12, 2, 10, 14, 1, 3, 8, 11, 6, 15, 13,
}

var _r = [80]uint{
	5, 14, 7, 0, 9, 2, 11, 4, 13, 6, 15, 8, 1, 10, 3, 12,
	6, 11, 3, 7, 0, 13, 5, 10, 14, 15, 8, 12, 4, 9, 1, 2,
	15, 5, 1, 3, 7, 14, 6, 9, 11, 8, 12, 2, 10, 0, 4, 13,
	8, 6, 4, 1, 3, 11, 15, 0, 5, 12, 2, 13, 9, 7, 10, 14,
	12, 15, 10, 4, 1, 5, 8, 7, 6, 2, 13, 14, 0, 3, 9, 11,
}

var s = [80]uint{
	11, 14, 15, 12, 5, 8, 7, 9, 11, 13, 14, 15, 6, 7, 9, 8,
	7, 6, 8, 13, 11, 9, 7, 15, 7, 12, 15, 9, 11, 7, 13, 12,
	11, 13, 6, 7, 14, 9, 13, 15, 14, 8, 13, 6, 5, 12, 7, 5,
	11, 12, 14, 15, 14, 15, 9, 8, 9, 14, 5, 6, 8, 6, 5, 12,
	9, 15, 5, 11, 6, 8, 13, 12, 5, 12, 13, 14, 11, 8, 5, 6,
}

var _s = [80]uint{
	8, 9, 9, 11, 13, 15, 15, 5, 7, 7, 8, 11, 14, 14, 12, 6,
	9, 13, 15, 7, 12, 8, 9, 11, 7, 7, 12, 7, 6, 15, 13, 11,
	9, 7, 15, 11, 8, 6, 6, 14, 12, 13, 5, 14, 13, 13, 7, 5,
	15, 5, 8, 11, 14, 14, 6, 14, 6, 9, 12, 9, 12, 5, 15, 8,
	8, 5, 12, 9, 12, 5, 14, 6, 8, 13, 6, 5, 15, 13, 11, 11,
}

func f(j, x, y, z uint32) uint32 {
	if 0 <= j && j <= 15 {
		return x ^ y ^ z
	} else if 16 <= j && j <= 31 {
		return (x & y) | ((^x) & z)
	} else if 32 <= j && j <= 47 {
		return (x | (^y)) ^ z
	} else if 48 <= j && j <= 63 {
		return (x & z) | (y & (^z))
	} else if 64 <= j && j <= 79 {
		return x ^ (y | (^z))
	}
	return 0
}

func rotate(x uint32, n uint) uint32 {
	return (x << n) | (x >> (32 - n))
}

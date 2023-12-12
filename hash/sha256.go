package hash

import (
	"encoding/binary"
	"fmt"
)

func SHA256(message []byte) string {

	var H = []uint32{
		0x6A09E667,
		0xBB67AE85,
		0x3C6EF372,
		0xA54FF53A,
		0x510E527F,
		0x9B05688C,
		0x1F83D9AB,
		0x5BE0CD19,
	}

	var k = [64]uint32{
		0x428A2F98, 0x71374491, 0xB5C0FBCF, 0xE9B5DBA5, 0x3956C25B, 0x59F111F1, 0x923F82A4, 0xAB1C5ED5,
		0xD807AA98, 0x12835B01, 0x243185BE, 0x550C7DC3, 0x72BE5D74, 0x80DEB1FE, 0x9BDC06A7, 0xC19BF174,
		0xE49B69C1, 0xEFBE4786, 0x0FC19DC6, 0x240CA1CC, 0x2DE92C6F, 0x4A7484AA, 0x5CB0A9DC, 0x76F988DA,
		0x983E5152, 0xA831C66D, 0xB00327C8, 0xBF597FC7, 0xC6E00BF3, 0xD5A79147, 0x06CA6351, 0x14292967,
		0x27B70A85, 0x2E1B2138, 0x4D2C6DFC, 0x53380D13, 0x650A7354, 0x766A0ABB, 0x81C2C92E, 0x92722C85,
		0xA2BFE8A1, 0xA81A664B, 0xC24B8B70, 0xC76C51A3, 0xD192E819, 0xD6990624, 0xF40E3585, 0x106AA070,
		0x19A4C116, 0x1E376C08, 0x2748774C, 0x34B0BCB5, 0x391C0CB3, 0x4ED8AA4A, 0x5B9CCA4F, 0x682E6FF3,
		0x748F82EE, 0x78A5636F, 0x84C87814, 0x8CC70208, 0x90BEFFFA, 0xA4506CEB, 0xBEF9A3F7, 0xC67178F2,
	}

	l := len(message)
	message = append(message[:], []byte{0x80}...)

	for len(message)*8%512 != 448 {
		message = append(message, 0x00)
	}

	ms := make([]byte, 8)
	binary.BigEndian.PutUint64(ms, uint64(l*8))
	message = append(message[:], ms[:]...)

	var blocks [][16]uint32

	messageLen := len(message) / 64

	for i := 0; i < messageLen; i++ {
		var block [16]uint32
		for j := 0; j < 16; j++ {
			block[j] = binary.BigEndian.Uint32(message[i*64+j*4 : i*64+(j+1)*4])
		}
		blocks = append(blocks, block)
	}
	//-
	W := [64]uint32{}

	for _, m := range blocks {
		a := H[0]
		b := H[1]
		c := H[2]
		d := H[3]
		e := H[4]
		f := H[5]
		g := H[6]
		h := H[7]

		for i := 0; i < 64; i++ {
			if 0 <= i && i <= 15 {
				W[i] = m[i]
			} else if 16 <= i && i <= 63 {
				s0 := Rotate(W[i-15], 7) ^ Rotate(W[i-15], 18) ^ SHR(W[i-15], 3)
				s1 := Rotate(W[i-2], 17) ^ Rotate(W[i-2], 19) ^ SHR(W[i-2], 10)
				W[i] = W[i-16] + s0 + W[i-7] + s1
			}
		}
		for i := 0; i < 64; i++ {
			e0 := Rotate(a, 2) ^ Rotate(a, 13) ^ Rotate(a, 22)
			ma := (a & b) ^ (a & c) ^ (b & c)
			t2 := e0 + ma
			e1 := Rotate(e, 6) ^ Rotate(e, 11) ^ Rotate(e, 25)
			ch := (e & f) ^ ((^e) & g)
			t1 := h + e1 + ch + k[i] + W[i]

			h = g
			g = f
			f = e
			e = d + t1
			d = c
			c = b
			b = a
			a = t1 + t2
		}
		H[0] = a + H[0]
		H[1] = b + H[1]
		H[2] = c + H[2]
		H[3] = d + H[3]
		H[4] = e + H[4]
		H[5] = f + H[5]
		H[6] = g + H[6]
		H[7] = h + H[7]
	}
	var hash string
	for _, s := range H {
		hash += fmt.Sprintf("%x", s)
	}

	return hash
}

func Rotate(x uint32, n uint) uint32 {
	return (x >> n) | (x << (32 - n))
}

func SHR(x uint32, n uint) uint32 {
	return x >> n
}

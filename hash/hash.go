package hash

import (
	"encoding/binary"
)

func SHA1(message string) [20]byte {
	h0 := uint32(0x67452301)
	h1 := uint32(0xEFCDAB89)
	h2 := uint32(0x98BADCFE)
	h3 := uint32(0x10325476)
	h4 := uint32(0xC3D2E1F0)

	data := []byte(message)

	messageLength := uint64(len(data) * 8)

	data = append(data, 0x80)

	for (len(data)+8)%64 != 0 {
		data = append(data, 0x00)
	}

	messageLengthBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(messageLengthBytes, messageLength)
	data = append(data, messageLengthBytes...)

	for i := 0; i < len(data); i += 64 {
		block := data[i : i+64]
		h0, h1, h2, h3, h4 = ProcessBlock(block, h0, h1, h2, h3, h4)
	}

	hash := [20]byte{}
	binary.BigEndian.PutUint32(hash[0:], h0)
	binary.BigEndian.PutUint32(hash[4:], h1)
	binary.BigEndian.PutUint32(hash[8:], h2)
	binary.BigEndian.PutUint32(hash[12:], h3)
	binary.BigEndian.PutUint32(hash[16:], h4)

	return hash
}

func ProcessBlock(block []byte, h0, h1, h2, h3, h4 uint32) (uint32, uint32, uint32, uint32, uint32) {

	words := make([]uint32, 80)
	for i := 0; i < 16; i++ {
		words[i] = binary.BigEndian.Uint32(block[i*4 : (i+1)*4])
	}

	for i := 16; i < 80; i++ {
		words[i] = LeftRotate(words[i-3]^words[i-8]^words[i-14]^words[i-16], 1)
	}

	a, b, c, d, e := h0, h1, h2, h3, h4

	for i := 0; i < 80; i++ {
		var f, k uint32

		if i < 20 {
			f = (b & c) | ((^b) & d)
			k = 0x5A827999
		} else if i < 40 {
			f = b ^ c ^ d
			k = 0x6ED9EBA1
		} else if i < 60 {
			f = (b & c) | (b & d) | (c & d)
			k = 0x8F1BBCDC
		} else {
			f = b ^ c ^ d
			k = 0xCA62C1D6
		}

		temp := LeftRotate(a, 5) + f + e + k + words[i]
		e = d
		d = c
		c = LeftRotate(b, 30)
		b = a
		a = temp
	}

	h0 += a
	h1 += b
	h2 += c
	h3 += d
	h4 += e

	return h0, h1, h2, h3, h4
}

func LeftRotate(value, shift uint32) uint32 {
	return (value << shift) | (value >> (32 - shift))
}

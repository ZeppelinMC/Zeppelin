package cfb8

import "crypto/cipher"

// CFB stream with 8 bit segment size
// See http://csrc.nist.gov/publications/nistpubs/800-38a/sp800-38a.pdf
type CFB8 struct {
	// the input block used to produce an output block
	//also the same as the iv, but overtime it's changed
	input []byte

	//output used to decipher and cipher plaintext and cipher segments
	output []byte

	//block is the forward crypto function
	block cipher.Block

	decrypt bool
}

func (x *CFB8) XORKeyStream(dst, src []byte) {
	//notes:
	//the plaintext and ciphertext segments consist of s bits(8 in this case)

	for i, v := range src {
		//the forward crypto operation is applied to
		//the IV to produce the first output block
		x.block.Encrypt(x.output, x.input)

		//both the cipher plain text and segments are XOR with the msb s bits of the output block
		dst[i] = v ^ x.output[0]

		//makes room to concatenate s bits of cipher text segment for
		//encryption, and plaintext segment for decryption
		copy(x.input, x.input[1:16])
		if x.decrypt {
			x.input[15] = v
		} else {
			x.input[15] = dst[i]
		}
	}
}

func NewCFB8(b cipher.Block, iv []byte, decrypt bool) *CFB8 {
	if b.BlockSize() != len(iv) {
		panic("blockSize and IV must be the same length")
	}

	cfb8 := &CFB8{
		block:   b,
		input:   make([]byte, len(iv)),
		output:  make([]byte, len(iv)),
		decrypt: decrypt,
	}

	//in CFB encryption, the first input block is the IV
	copy(cfb8.input, iv)
	return cfb8
}

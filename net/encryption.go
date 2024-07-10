package net

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/asn1"
	"fmt"
	"math/big"

	"github.com/dynamitemc/aether/net/packet/login"
)

func encrypt(sharedSecret, plaintext, dst []byte) error {
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, sharedSecret)
	stream.XORKeyStream(dst, plaintext)

	return nil
}

func decrypt(sharedSecret, ciphertext, dst []byte) error {
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return nil
	}

	stream := cipher.NewCFBDecrypter(block, sharedSecret)
	stream.XORKeyStream(dst, ciphertext)

	return nil
}

type AlgorithmIdentifier struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters asn1.RawValue `asn1:"optional"`
}

type SubjectPublicKeyInfo struct {
	Algorithm        AlgorithmIdentifier
	SubjectPublicKey asn1.BitString
}

type RSAPublicKey struct {
	Modulus        *big.Int
	PublicExponent int
}

func (c *Conn) Encrypt() error {
	pubKey := c.listener.privKey.PublicKey
	rsaPubKey := RSAPublicKey{
		Modulus:        pubKey.N,
		PublicExponent: pubKey.E,
	}

	// Marshal the RSA public key into ASN.1 DER format
	rsaPubKeyDER, err := asn1.Marshal(rsaPubKey)
	if err != nil {
		return err
	}

	spki := SubjectPublicKeyInfo{
		Algorithm: AlgorithmIdentifier{
			Algorithm:  asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}, // OID for RSA encryption
			Parameters: asn1.RawValue{Tag: 5},                             // NULL parameters
		},
		SubjectPublicKey: asn1.BitString{
			Bytes:     rsaPubKeyDER,
			BitLength: len(rsaPubKeyDER) * 8,
		},
	}

	// Marshal the SubjectPublicKeyInfo into ASN.1 DER format
	spkiDER, err := asn1.Marshal(spki)
	if err != nil {
		return err
	}

	verifyToken := make([]byte, 4)
	rand.Read(verifyToken)

	c.WritePacket(&login.EncryptionRequest{
		PublicKey:   spkiDER,
		VerifyToken: verifyToken,
	})
	p, err := c.ReadPacket()
	if err != nil {
		return err
	}
	res, ok := p.(*login.EncryptionResponse)
	if !ok {
		return fmt.Errorf("unsuccessful encryption")
	}
	c.encrypted = ok

	c.sharedSecret, err = rsa.DecryptPKCS1v15(nil, c.listener.privKey, res.SharedSecret)
	if err != nil {
		return err
	}
	c.verifyToken, err = rsa.DecryptPKCS1v15(nil, c.listener.privKey, res.VerifyToken)
	if err != nil {
		return err
	}

	if [4]byte(c.verifyToken) != [4]byte(verifyToken) {
		return fmt.Errorf("unsuccessful encryption")
	}

	return nil
}

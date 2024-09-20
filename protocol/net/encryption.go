package net

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"

	"github.com/zeppelinmc/zeppelin/protocol/net/cfb8"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/login"
)

func (c *Conn) encryptd(plaintext, dst []byte) {
	c.encrypter.XORKeyStream(dst, plaintext)
}

func (c *Conn) decryptd(ciphertext, dst []byte) {
	c.decrypter.XORKeyStream(dst, ciphertext)
}

func (c *Conn) encrypt() error {
	key, err := x509.MarshalPKIXPublicKey(&c.listener.privKey.PublicKey)
	if err != nil {
		return err
	}

	verifyToken := make([]byte, 4)
	rand.Read(verifyToken)

	c.WritePacket(&login.EncryptionRequest{
		PublicKey:          key,
		VerifyToken:        verifyToken,
		ShouldAuthenticate: c.listener.cfg.Authenticate,
	})
	p, s, err := c.ReadPacket()
	if err != nil || s {
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

	block, err := aes.NewCipher(c.sharedSecret)
	if err != nil {
		return err
	}
	c.encrypter = cfb8.NewCFB8(block, c.sharedSecret, false)
	c.decrypter = cfb8.NewCFB8(block, c.sharedSecret, true)

	return nil
}

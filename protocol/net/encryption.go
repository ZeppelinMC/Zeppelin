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

func (conn *Conn) encryptd(plaintext, dst []byte) {
	conn.encrypter.XORKeyStream(dst, plaintext)
}

func (conn *Conn) decryptd(ciphertext, dst []byte) {
	conn.decrypter.XORKeyStream(dst, ciphertext)
}

func (conn *Conn) encrypt() error {
	key, err := x509.MarshalPKIXPublicKey(&conn.listener.privKey.PublicKey)
	if err != nil {
		return err
	}

	verifyToken := make([]byte, 4)
	rand.Read(verifyToken)

	conn.WritePacket(&login.EncryptionRequest{
		PublicKey:          key,
		VerifyToken:        verifyToken,
		ShouldAuthenticate: conn.listener.cfg.Authenticate,
	})
	p, s, err := conn.ReadPacket()
	if err != nil || s {
		return err
	}
	res, ok := p.(*login.EncryptionResponse)
	if !ok {
		return fmt.Errorf("unsuccessful encryption")
	}
	conn.encrypted = ok

	conn.sharedSecret, err = rsa.DecryptPKCS1v15(nil, conn.listener.privKey, res.SharedSecret)
	if err != nil {
		return err
	}
	conn.verifyToken, err = rsa.DecryptPKCS1v15(nil, conn.listener.privKey, res.VerifyToken)
	if err != nil {
		return err
	}

	if [4]byte(conn.verifyToken) != [4]byte(verifyToken) {
		return fmt.Errorf("unsuccessful encryption")
	}

	block, err := aes.NewCipher(conn.sharedSecret)
	if err != nil {
		return err
	}
	conn.encrypter = cfb8.NewCFB8(block, conn.sharedSecret, false)
	conn.decrypter = cfb8.NewCFB8(block, conn.sharedSecret, true)

	return nil
}

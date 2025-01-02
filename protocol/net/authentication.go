package net

import (
	"crypto/sha1"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"unsafe"

	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/login"
)

func authurl(u, h string) string {
	return "https://sessionserver.mojang.com/session/minecraft/hasJoined?username=" + u + "&serverId=" + h
}

func (conn *Conn) authenticate() error {
	key, err := x509.MarshalPKIXPublicKey(&conn.listener.privKey.PublicKey)
	if err != nil {
		return err
	}
	hash := conn.sessionHash(key)
	res, err := http.Get(authurl(conn.username, hash))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("authenticated error: player not joined")
	}

	var response struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Properties []struct {
			Name      string `json:"name"`
			Value     string `json:"value"`
			Signature string `json:"signature"`
		} `json:"properties"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return err
	}

	conn.username = response.Name
	conn.uuid, err = uuid.Parse(response.ID)
	if err != nil {
		return err
	}
	conn.properties = *(*[]login.Property)(unsafe.Pointer(&response.Properties))

	return nil
}

func (conn *Conn) sessionHash(publicKey []byte) string {
	hash := sha1.New()
	hash.Write(conn.sharedSecret)
	hash.Write(publicKey)

	sum := hash.Sum(nil)

	negative := (sum[0] & 0x80) == 0x80
	if negative {
		sum = twosComplement(sum)
	}

	// Trim away zeroes
	res := strings.TrimLeft(hex.EncodeToString(sum), "0")
	if negative {
		res = "-" + res
	}
	return res
}

func twosComplement(p []byte) []byte {
	carry := true
	for i := len(p) - 1; i >= 0; i-- {
		p[i] = byte(^p[i])
		if carry {
			carry = p[i] == 0xff
			p[i]++
		}
	}
	return p
}

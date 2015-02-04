package main

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/ianmcmahon/encoding_ssh"
)

func encryptStringBySshRsaPublicKey(sshRsaPubkey string, msg string) ([]byte, error) {
	pubkey, err := ssh.DecodePublicKey(sshRsaPubkey)
	if err != nil {
		return make([]byte, 0), err
	}

	enc, err := rsa.EncryptPKCS1v15(rand.Reader, pubkey.(*rsa.PublicKey), []byte(msg))
	if err != nil {
		return make([]byte, 0), err
	}

	return enc, nil
}

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"

	"github.com/ianmcmahon/encoding_ssh"
)

func encryptStringBySSHRsaPublicKey(sshRsaPubkey string, msg string) ([]byte, error) {
	pubkey, err := ssh.DecodePublicKey(sshRsaPubkey)
	if err != nil {
		return make([]byte, 0), err
	}

	enc, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, pubkey.(*rsa.PublicKey), []byte(msg), nil)

	if err != nil {
		return make([]byte, 0), err
	}

	return enc, nil
}

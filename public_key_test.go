package main

import "testing"

func TestFetchPublicKeyThatExists(t *testing.T) {
	pubkey := FetchPublicKeyByUserName("moznion")
	if pubkey == "" {
		t.Error("Failed to fetch public key from GitHub")
	}
}

func TestFetchPublicKeyOfNotExistsUser(t *testing.T) {
	pubkey := FetchPublicKeyByUserName("moznion_not_exist")
	if pubkey != "" {
		t.Error("Something received!")
	}
}

package main

import "testing"

func TestFetchPublicKeyThatExists(t *testing.T) {
	pubkey := FetchPublicKeyByUserName("moznion")
	if pubkey == "" {
		t.Error("Failed to fetch public key from GitHub")
	}
}

func TestFetchPublicKeyOfNotExistsUser(t *testing.T) {
	pubkey := FetchPublicKeyByUserName("not_exist_asdfasdf")
	if pubkey != "" {
		t.Error("Something received!")
	}
}

func TestFetchEmailAddressThatExists(t *testing.T) {
	email := FetchEmailAddressByUserName("moznion")
	if email != "moznion@gmail.com" {
		t.Error("Failed to fetch users email")
	}
}

func TestFetchEmailAddressOfNotExistsUser(t *testing.T) {
	email := FetchEmailAddressByUserName("not_exist_asdfasdf")
	if email != "" {
		t.Error("Something received!")
	}
}

package main

import "testing"

func TestFetchPublicKeyThatExists(t *testing.T) {
	pubkey, err := FetchPublicKeyByUserName("moznion")
	if err != nil || pubkey == "" {
		t.Error("Failed to fetch public key from GitHub")
	}
}

func TestFetchPublicKeyOfNotExistsUser(t *testing.T) {
	_, err := FetchPublicKeyByUserName("not_exist_asdfasdf")

	if err == nil {
		t.Error("Failed to fetch public key from GitHub")
	}
}

func TestFetchEmailAddressThatExists(t *testing.T) {
	email, err := FetchEmailAddressByUserName("moznion")
	if err != nil || email != "moznion@gmail.com" {
		t.Error("Failed to fetch users email")
	}
}

func TestFetchEmailAddressOfNotExistsUser(t *testing.T) {
	_, err := FetchEmailAddressByUserName("not_exist_asdfasdf")
	if err == nil {
		t.Error("Something received!")
	}
}

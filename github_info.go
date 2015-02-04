package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func fetchPublicKeyByUserId(userId string) (string, error) {
	res, err := http.Get("https://github.com/" + userId + ".keys")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Failed HTTP request to fetch a public key: %d", res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	r := regexp.MustCompile("\r?\n")
	lines := r.Split(string(body), -1)
	return lines[0], nil
}

type users struct {
	Email string `json:"email"`
}

func fetchEmailAddressByUserId(userId string) (string, error) {
	res, err := http.Get("https://api.github.com/users/" + userId)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Failed HTTP request to fetch an email adress: %d", res.StatusCode))
	}

	userInfo := new(users)
	err = json.NewDecoder(res.Body).Decode(userInfo)
	if err != nil {
		return "", err
	}

	return userInfo.Email, nil
}

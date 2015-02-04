package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func FetchPublicKeyByUserName(userName string) (string, error) {
	res, err := http.Get("https://github.com/" + userName + ".keys")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Failed HTTP request: %d", res.StatusCode))
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

func FetchEmailAddressByUserName(userName string) (string, error) {
	res, err := http.Get("https://api.github.com/users/" + userName)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("Failed HTTP request: %d", res.StatusCode))
	}

	userInfo := new(users)
	err = json.NewDecoder(res.Body).Decode(userInfo)
	if err != nil {
		return "", err
	}

	return userInfo.Email, nil
}

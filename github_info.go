package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
)

func FetchPublicKeyByUserName(userName string) string {
	res, err := http.Get("https://github.com/" + userName + ".keys")
	if err != nil {
		// TODO
		return ""
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		// TODO
		return ""
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// TODO
		return ""
	}

	r := regexp.MustCompile("\r?\n")
	lines := r.Split(string(body), -1)
	return lines[0]
}

type users struct {
	Email string `json:"email"`
}

func FetchEmailAddressByUserName(userName string) string {
	res, err := http.Get("https://api.github.com/users/" + userName)
	if err != nil {
		// TODO
		return ""
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		// TODO
		return ""
	}

	userInfo := new(users)
	err = json.NewDecoder(res.Body).Decode(userInfo)
	if err != nil {
		// TODO
		return ""
	}

	return userInfo.Email
}

package main

import (
	"io/ioutil"
	"testing"
)

func TestParseYamlSuccessfully(t *testing.T) {
	tempfile, _ := ioutil.TempFile("", "")
	defer tempfile.Close()

	tempfileName := tempfile.Name()
	ioutil.WriteFile(
		tempfileName,
		[]byte("address: hoge@example.com\npassword: xxx"),
		0644,
	)

	conf, err := readYamlConfig(tempfileName)
	if err != nil || conf.Address != "hoge@example.com" || conf.Password != "xxx" {
		t.Error("Failed to read the yaml file")
	}
}

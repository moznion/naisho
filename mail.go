package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/gomail.v1"
)

type mail struct {
	fromAddr string
	pass     string
	toAddr   string
	msg      []byte
	subject  string
	body     string
}

func sendByGmail(m *mail) (err error) {
	const (
		gmailSMTPAddr = "smtp.gmail.com"
	)

	tempfile, err := ioutil.TempFile("", "secret-")
	if err != nil {
		return err
	}
	defer tempfile.Close()

	tempfileName := tempfile.Name()
	err = ioutil.WriteFile(tempfileName, m.msg, 0644)
	if err != nil {
		return err
	}

	subject := m.subject
	if subject == "" {
		subject = "Secret message"
	}

	body := m.body
	if body == "" {
		body = fmt.Sprintf("Please execute with attachment file to read: `openssl rsautl -decrypt -inkey <YOUR SECRET KEY> -in %s`", filepath.Base(tempfileName))
	}

	newMsg := gomail.NewMessage()
	newMsg.SetHeader("From", m.fromAddr)
	newMsg.SetHeader("To", m.toAddr)
	newMsg.SetHeader("Subject", subject)
	newMsg.SetBody("text/plain", body)

	f, err := gomail.OpenFile(tempfile.Name())
	if err != nil {
		return err
	}
	newMsg.Attach(f)

	smtpPort := 587
	mailer := gomail.NewMailer(gmailSMTPAddr, m.fromAddr, m.pass, smtpPort)
	return mailer.Send(newMsg)
}

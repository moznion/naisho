package main

import "net/smtp"

func sendByGmail(senderAddr string, senderPswd string, dstAddr string, body string) (err error) {
	const (
		gmailSmtpAddr = "smtp.gmail.com"
	)
	auth := smtp.PlainAuth("", senderAddr, senderPswd, gmailSmtpAddr)
	return smtp.SendMail(gmailSmtpAddr+":587", auth, senderAddr, []string{dstAddr}, []byte(body))
}

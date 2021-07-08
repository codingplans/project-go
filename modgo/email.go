package main

import (
	"github.com/jordan-wright/email"
	log "github.com/sirupsen/logrus"
	"net/smtp"
)

func TestExampleGmail() {
	println(222)
	e := email.NewEmail()
	e.From = "wwwwww@gmail.com"
	e.To = []string{"wwwwwww@126.com"}
	// e.Bcc = []string{"darrenzzy@126.com"}
	// e.Cc = []string{"darrenzzy@126.com"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!\n")
	e.HTML = []byte("<h1>Fancy Html is supported, too!</h1>\n")
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", e.From, "12345678", "smtp.gmail.com"))
	if err != nil {
		log.Info(err.Error())

	}
	println(333333)
}

package gomailer

import (
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

/*
export envs:

	EMAIL_GMAIL_FROM
	EMAIL_GMAIL_PASSWORD
*/
type DataEmail struct {
	EmailTo      string  `json:"email_to"`
	EmailSubject string  `json:"email_subject"`
	EmailBody    string  `json:"email_body"`
	Username     string  `json:"user_name"`
	SenderName   string  `json:"sender_name"`
	Css          *string `json:"css,omitempty"`
	Logo         *string `json:"logo,omitempty"`
}

func SendEmailGmail(data DataEmail) error {
	// Informações da conta Gmail
	imgLogo := ""
	from := os.Getenv("EMAIL_GMAIL_FROM")
	password := os.Getenv("EMAIL_GMAIL_PASSWORD")
	if from == "" || password == "" {
		fmt.Printf("the EMAIL_GMAIL_FROM or EMAIL_GMAIL_PASSWORD is not set! Please, export the enviroment variables.")
		return errors.New("the EMAIL_GMAIL_FROM or EMAIL_GMAIL_PASSWORD is not set! Please, export the enviroment variables")
	}

	if data.Logo != nil {
		imgLogo = `<img  src="` + *data.Logo + `"  style="width:100px;position:relative; margin:10px auto; display:block;"/>`
	}
	css := `p{ font-size: 14px;}`
	if data.Css != nil {
		css = *data.Css
	}

	to := []string{data.EmailTo}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587" // Porta TLS/STARTTLS
	subject := data.EmailSubject
	body := `
	<!DOCTYPE html>
		<html>
		<head>
		  <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
	  <meta name="color-scheme" content="dark">
    <meta name="supported-color-schemes" content="dark">
			<style>
			` + css + `			   
			.container {  padding: 20px; border-radius: 8px; }
			</style>
		</head>
		<body>
			<div class="container">
			 ` + imgLogo + `
				` + data.EmailBody + `				
			</div>
		</body>
		</html>`
	fromHeader := fmt.Sprintf("%s <%s>", data.SenderName, from)
	msg := []byte("From: " + fromHeader + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"utf-8\"\r\n" +
		"\r\n" +
		body)
	// Autenticação
	auth := smtp.PlainAuth("", from, password, smtpHost)
	// Envia o e-mail
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		fmt.Printf("Erro ao enviar e-mail: %s\n", err)
		return err
	}
	fmt.Println("Email sent with success!")
	return nil
}

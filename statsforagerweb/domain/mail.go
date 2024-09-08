package domain

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"
	"time"
)

type SmtpConfig struct {
	From, User, Host, Port, Password string
}

type Mail struct {
	config SmtpConfig
}

func NewMail(config SmtpConfig) (Mail, error) {
	var mail Mail

	if config.User == "" {
		return mail, errors.New("Need SmtpConfig with value for User.")
	}

	if config.From == "" {
		return mail, errors.New("Need SmtpConfig with value for From.")
	}

	if config.Password == "" {
		return mail, errors.New("Need SmtpConfig with value for Password.")
	}

	if config.Host == "" {
		return mail, errors.New("Need SmtpConfig with value for Host.")
	}

	if config.Port == "" {
		return mail, errors.New("Need SmtpConfig with value for Port.")
	}

	return Mail{config: config}, nil
}

func (mail *Mail) SendMailWithTls(to, subject, body string) {
	message := []byte("To: " + to + "\r\n" +
		"From: StatsForager <" + mail.config.From + ">\r\n" +
		"Subject: " + subject + "\r\n" +
		"Date: " + time.Now().Format(time.RFC1123Z) + "\r\n" +
		"\r\n" + // Blank line separating headers from body
		body)

	fmt.Println(string(message))
	auth := smtp.PlainAuth(
		"",
		mail.config.User,
		mail.config.Password,
		mail.config.Host,
	)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         mail.config.Host,
	}

	connection, err := tls.Dial("tcp", mail.config.Host+":"+mail.config.Port, tlsConfig)
	if err != nil {
		fmt.Println("SendMailWithTls error:", err)
		return
	}

	smtpClient, err := smtp.NewClient(connection, mail.config.Host)
	if err != nil {
		fmt.Println("SendMailWithTls error:", err)
		return
	}

	if err = smtpClient.Auth(auth); err != nil {
		fmt.Println("SendMailWithTls error:", err)
		return
	}

	if err = smtpClient.Mail(mail.config.From); err != nil {
		fmt.Println("SendMailWithTls error:", err)
		return
	}

	if err = smtpClient.Rcpt(to); err != nil {
		fmt.Println("SendMailWithTls error:", err)
		return
	}

	writer, err := smtpClient.Data()
	if err != nil {
		fmt.Println("SendMailWithTls error:", err)
		return
	}

	_, err = writer.Write([]byte(message))
	if err != nil {
		fmt.Println("SendMailWithTls error:", err)
		return
	}

	err = writer.Close()
	if err != nil {
		fmt.Println("SendMailWithTls error:", err)
		return
	}

	err = smtpClient.Quit()
	if err != nil {
		fmt.Println("SendMailWithTls error:", err)
		return
	}
}

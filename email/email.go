package email

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"net/textproto"
	"strings"

	"gopkg.in/gomail.v2"
)

type SMTPError string

func (s SMTPError) Error() string {
	return string(s)
}

const (
	SmtpServerAddrError SMTPError = "SMTP服务器地址有误"
)

type SmtpServerConfig struct {
	Username string
	Password string
	Host     string
	Port     int
}

func SendViaSmtp(from, subject, content string, receivers []string, smtpServerConfig SmtpServerConfig) error {

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", receivers...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	d := gomail.NewDialer(smtpServerConfig.Host, smtpServerConfig.Port, smtpServerConfig.Username, smtpServerConfig.Password)
	err := d.DialAndSend(m)
	fmt.Printf("err:%v", err)
	var dnsError *net.DNSError
	if errors.As(err, &dnsError) {
		return fmt.Errorf("%w: %v", SmtpServerAddrError, err)
	}
	var opError *net.OpError
	if errors.As(err, &opError) {
		return fmt.Errorf("%w: %v", SmtpServerAddrError, err)
	}
	var perr *textproto.Error
	if errors.As(err, &perr) {
		println()
		println("p", err.Error())
		println("p", perr.Code)
		return fmt.Errorf("%w: %v", SmtpServerAddrError, err)
	}
	//if err.Error() == "gomail: could not send email 1: 550 5.1.1 recipient is not exist" {
	//	println("11111111111")
	//}

	return nil
}

// SendEmailParams 邮件发送参数
type SendEmailParams struct {
	From     string
	To       []string
	Message  string
	Password string
	SMTPHost string
	SMTPPort string
}

// invokeSendEmail 调用golang库发送邮件
func invokeSendEmail(in *SendEmailParams) error {
	auth := smtp.PlainAuth("", in.From, in.Password, in.SMTPHost)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: false,
		MinVersion:         tls.VersionTLS12,
		ServerName:         in.SMTPHost,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", net.JoinHostPort(in.SMTPHost, in.SMTPPort), tlsconfig)
	if err != nil {
		return err
	}
	c, err := smtp.NewClient(conn, in.SMTPHost)
	if err != nil {
		return err
	}

	if err := c.Auth(auth); err != nil {
		return err
	}
	if err := c.Mail(in.From); err != nil {
		return err
	}
	for _, to := range in.To {
		if err := c.Rcpt(to); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(in.Message))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}

// SendHtmlEmail 发送邮箱
func SendHtmlEmail(smtpHost, smtpPort, username, password string, to []string, subject, body string) error {
	if len(to) == 0 {
		return fmt.Errorf("mail receiver required")
	}

	headers := make(map[string]string)
	headers["From"] = username
	tostr := strings.Join(to, ",")
	headers["To"] = tostr
	sub := base64.StdEncoding.EncodeToString([]byte(subject))
	headers["Subject"] = "=?UTF-8?B?" + sub + "?="
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message += "\r\n" + body

	return invokeSendEmail(&SendEmailParams{
		From:     username,
		To:       to,
		Message:  message,
		Password: password,
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
	})
}

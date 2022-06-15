package email

import "testing"

var username = ""
var password = ""
var smtpHost = ""
var smtpPort = 25
var smtpSslPort = "465"

var to = []string{""}

func TestSendViaSmtp(t *testing.T) {

	err := SendViaSmtp(username,
		"title",
		"hello world",
		to,
		SmtpServerConfig{
			Username: username,
			Password: password,
			Host:     smtpHost,
			Port:     smtpPort,
		},
	)
	if err != nil {
		t.Log(err)
	}
}

func TestSendHtmlEmail(t *testing.T) {
	subject := "测试title"
	body := "测试body"
	err := SendHtmlEmail(smtpHost, smtpSslPort, username, password, to, subject, body)
	if err != nil {
		t.Log(err)
	}
}

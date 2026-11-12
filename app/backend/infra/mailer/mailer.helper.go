package mailer

import (
	"strings"
)

func CreateEnvelope(mail Mail) []byte {
	var msg []byte
	msg = append(msg, []byte("From: "+mail.From+"\r\n")...)
	msg = append(msg, []byte("To: "+strings.Join(mail.To, ", ")+"\r\n")...)
	if len(mail.Cc) > 0 {
		msg = append(msg, []byte("Cc: "+strings.Join(mail.Cc, ", ")+"\r\n")...)
	}
	if len(mail.Bcc) > 0 {
		msg = append(msg, []byte("Bcc: "+strings.Join(mail.Bcc, ", ")+"\r\n")...)
	}

	return msg
}

func CreateHeader(mail Mail) []byte {
	var msg []byte
	msg = append(msg, []byte("Subject: "+mail.Subject+"\r\n")...)
	msg = append(msg, []byte("\r\n")...)
	return msg
}

func CreateBody(mail Mail) []byte {
	var msg []byte
	msg = append(msg, []byte(mail.Body+"\r\n")...)
	return msg
}

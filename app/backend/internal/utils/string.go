package utils

import "net/mail"

func IsEmailValid(email string) bool {
	addr, err := mail.ParseAddress(email)
	return err == nil && addr.Address == email
}

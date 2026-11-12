package config

import "crypto/rand"

type AuthConfig struct {
	SecretKeyJWT string
	Issuer       string
}

func LoadAuthConfig() AuthConfig {
	return AuthConfig{
		SecretKeyJWT: getEnv("SECRET_KEY_JWT", rand.Text()),
		Issuer:       getEnv("ISSUER", "SGINTERN"),
	}
}

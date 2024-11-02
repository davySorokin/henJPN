package main

import (
	"encoding/base32"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateTOTP(email string) (string, error) {
	secret := email + "HENNGECHALLENGE003"
	encodedSecret := base32.StdEncoding.EncodeToString([]byte(secret))

	otp, err := totp.GenerateCodeCustom(encodedSecret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    10,
		Algorithm: otp.AlgorithmSHA512,
	})
	if err != nil {
		return "", err
	}

	return otp, nil
}

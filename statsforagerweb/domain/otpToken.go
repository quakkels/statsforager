package domain

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type OtpToken struct {
	Otp           string
	ExpirationUtc time.Time
}

func NewOtpToken(expiration time.Duration) (OtpToken, error) {
	length := 15
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return OtpToken{}, err
	}

	otpToken := OtpToken{
		Otp:           base64.URLEncoding.EncodeToString(buffer)[:length],
		ExpirationUtc: time.Now().UTC().Add(expiration),
	}

	return otpToken, nil
}

func (otpToken *OtpToken) IsValid(otpToTest string) bool {
	nowUtc := time.Now().UTC()
	if otpToTest == otpToken.Otp && otpToken.ExpirationUtc.After(nowUtc) {
		return true
	}
	return false
}

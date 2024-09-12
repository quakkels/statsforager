package domain

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"
)

type OtpToken struct {
	AccountCode   string
	Otp           string
	ExpirationUtc time.Time
	Thumbprint    string
}

func NewOtpToken(accountCode string, expiration time.Duration) (OtpToken, error) {
	length := 15
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return OtpToken{}, err
	}

	otpToken := OtpToken{
		AccountCode:   accountCode,
		Otp:           base64.URLEncoding.EncodeToString(buffer)[:length],
		ExpirationUtc: time.Now().UTC().Add(expiration),
	}
	otpToken.setHashAsHex()

	return otpToken, nil
}

func (otpToken *OtpToken) IsValid(otpToTest string) bool {
	nowUtc := time.Now().UTC()
	hashToTest := otpToken.hashOtpAsHex(otpToTest)
	if otpToken.Otp == otpToTest &&
		otpToken.ExpirationUtc.After(nowUtc) &&
		otpToken.Thumbprint == hashToTest {
		return true
	}
	fmt.Println("originalHash:", otpToken.Thumbprint, "\nhashToTest:", hashToTest)
	return false
}

func (otpToken *OtpToken) setHashAsHex() {
	otpToken.Thumbprint = otpToken.hashOtpAsHex(otpToken.Otp)
}

func (otpToken *OtpToken) hashOtpAsHex(otp string) string {
	data := fmt.Sprintf("%s%s%d", otpToken.AccountCode, otp, otpToken.ExpirationUtc.Unix())
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
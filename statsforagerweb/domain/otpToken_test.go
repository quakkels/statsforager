package domain_test

import (
	"statsforagerweb/domain"
	"testing"
	"time"
)

func TestNewOtpToken(t *testing.T) {
	// arrange
	email := "email@example.com"
	validFor := 10 * time.Millisecond

	// act
	otp, err := domain.NewOtpToken(email, validFor)

	// assert
	if err != nil {
		t.Error(err)
	}
	if len(otp.Otp) != 15 {
		t.Error("Otp has incorrect length. Expected 15. Got:", len(otp.Otp))
	}
	if otp.IsValid("wrong") {
		t.Error("Otp.IsValid() succeeded when it should have failed due to wrong otp")
	}
	if !otp.IsValid(otp.Otp) {
		t.Error("Otp.IsValid() check failed when it should have succeeded.")
	}

	time.Sleep(validFor + time.Millisecond)
	if otp.IsValid(otp.Otp) {
		t.Error("Otp.IsValid() succeeded when it should have failed due to exceeding expiration.")
	}
}

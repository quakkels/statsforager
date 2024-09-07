package domain_test

import (
	"statsforagerweb/domain"
	"testing"
)

func testValidAccount(t *testing.T) {
	account := domain.Account{Email: "valid@example.com"}

	// act
	result := account.ValidateAccount()

	// assert
	if !result.IsSuccess {
		t.Fatal("Account validation expected success")
	}
}

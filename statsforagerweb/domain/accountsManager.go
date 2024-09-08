package domain

import (
	"context"
	"fmt"
	"net/mail"
	"net/url"
	"strings"
)

type AccountsConfig struct {
	AppRoot string
}

type Account struct {
	Email    string
	IsActive bool
}

func (account *Account) ValidateAccount() validationResult {
	messages := make(map[string]string)
	_, err := mail.ParseAddress(account.Email)
	if err != nil {
		messages["email"] = "Invalid email"
	}
	return *NewValidationResult(messages)
}

type AccountsRepository interface {
	GetAccountByEmail(context.Context, string) (Account, error)
	RegisterAccount(context.Context, string) error
	SaveAccount(context.Context, Account) error
}

type AccountsManager struct {
	config       AccountsConfig
	accountsRepo AccountsRepository
	mail         Mail
}

func NewAccountsManager(config AccountsConfig, accountsRepo AccountsRepository, mail Mail) AccountsManager {
	return AccountsManager{
		config:       config,
		accountsRepo: accountsRepo,
		mail:         mail,
	}
}

func (manager *AccountsManager) RegisterEmail(context context.Context, email string) (validationResult, error) {
	account := Account{email, false}

	validationResult := account.ValidateAccount()
	if !validationResult.IsSuccess {
		return validationResult, nil
	}

	err := manager.accountsRepo.RegisterAccount(context, account.Email)
	if err != nil {
		return validationResult, err
	}

	return validationResult, nil
}

func (manager *AccountsManager) SendLoginMail(
	context context.Context,
	email string,
	otp string,
) (validationResult, error) {
	successResult := *NewValidationResult(nil)
	if len(email) == 0 {
		errors := make(map[string]string)
		errors["email"] = "Missing email"
		return *NewValidationResult(errors), nil
	}
	account, err := manager.accountsRepo.GetAccountByEmail(context, email)
	if err != nil {
		fmt.Println("GetAccountByEmail")
		fmt.Println(err)
		return successResult, err // return success to avoid leaking account info
	}

	if account.Email != strings.ToLower(email) {
		fmt.Println(
			"account.Email:", account.Email+
				"\nstrings.ToLower(email):", strings.ToLower(email))
		fmt.Println("emails don't match")
		return successResult, nil // return success to avoid leaking account info... todo: may be unreachable code
	}

	go manager.mail.SendMailWithTls(
		account.Email,
		"StatsForager Login Confirmation",
		"Complete your passwordless log in to StatsForager by following this link:\r\n\r\n\t"+manager.config.AppRoot+"/login/confirm/"+url.PathEscape(otp)+"\r\nThank you,\r\nStatsForager")

	return successResult, err
}

package domain_test

import (
	"statsforagerweb/domain"
	"testing"
)

func TestNewValidationResultSucceedsWithNilMessages(t *testing.T) {
	// arrange
	var messages map[string]string

	// act
	result := domain.NewValidationResult(messages)

	// assert
	if !result.IsSuccess {
		t.Fatal("Expected true but got: ", result.IsSuccess)
	}
}

func TestNewValidationResultSucceedsWithEmptyMessages(t *testing.T) {
	// arrange
	messages := make(map[string]string)
	// act
	result := domain.NewValidationResult(messages)

	// assert
	if !result.IsSuccess {
		t.Fatal("Expected true but got: ", result.IsSuccess)
	}
}

func TestNewValidationResultFailsWithMessages(t *testing.T) {
	// arrange
	messages := make(map[string]string)
	messages["PropertyName"] = "FailMessage"

	// act
	result := domain.NewValidationResult(messages)

	// assert
	if result.IsSuccess {
		t.Fatal("Expected false but got: ", result.IsSuccess)
	}

	if result.Messages["PropertyName"] != "FailMessage" {
		t.Fatal("Expected 'FailMessage' but got: ", result.Messages["PropertyName"])
	}
}

func TestToMessageSliceSucceedsWhenMissingMessages(t *testing.T) {
	results := domain.NewValidationResult(make(map[string]string))

	results.ToMessagesSlice()
}

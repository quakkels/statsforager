package domain_test

import (
	"statsforagerweb/domain"
	"testing"
)

func TestNilSite(t *testing.T) {
	// arrange
	var site domain.Site

	// act
	result := site.HasLocation("location")

	// assert
	if result {
		t.Fatal("Expected false but got: ", result)
	}
}

func TestSiteHasLocation(t *testing.T) {
	// arrange
	expected := true
	site := domain.Site{Domain: "example.com"}

	// act
	result := site.HasLocation("https://example.com/long/path?field=value")
	result2 := site.HasLocation("http://example.com/long/path?field=value")

	// assert
	if result != expected {
		t.Fatal("Expected", expected, "but got:", result)
	}
	if result2 != expected {
		t.Fatal("Expected", expected, "but got:", result2)
	}
}

func TestSiteDoesNotHaveLocation(t *testing.T) {
	// arrange
	expected := false
	site := domain.Site{Domain: "example.com"}

	// act
	result := site.HasLocation("https://www.example.com/long/path?field=value")

	// assert
	if result != expected {
		t.Fatal("Expected", expected, "but got:", result)
	}
}

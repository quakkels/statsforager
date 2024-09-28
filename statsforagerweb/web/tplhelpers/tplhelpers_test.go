package tplhelpers_test

import (
	"html/template"
	"statsforagerweb/web/tplhelpers"
	"testing"
)

func TestSelect(t *testing.T) {
	expected := template.HTML(
		`<select name="n" class="one two"><option value="key1">one</option><option value="key2" selected="selected">two</option></select>`,
	)
	options := map[string]string{
		"key1": "one",
		"key2": "two",
	}

	attributes := map[string]string{
		"name":  "n",
		"class": "one two",
	}

	result := tplhelpers.Select("key2", options, attributes)

	if expected != result {
		t.Error("Expected select HTML does not match result. \nExpected:\n", expected, "\nResult:\n", result)
	}
}

func TestMakeMap(t *testing.T) {
	t.Error("test is unwritten.")
}

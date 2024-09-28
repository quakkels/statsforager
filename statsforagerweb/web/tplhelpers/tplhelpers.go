package tplhelpers

import (
	"errors"
	"fmt"
	"html/template"
)

func Select(selected string, options, attributes map[string]string) template.HTML {
	attributesHtml := makeAttributesHtml(attributes)
	optionsHtml := ""
	for value, label := range options {
		selectedHtml := ""
		if selected == value {
			selectedHtml = ` selected="selected"`
		}
		optionsHtml += fmt.Sprintf(
			`<option value="%s"%s>%s</option>`,
			template.HTMLEscapeString(value),
			selectedHtml,
			template.HTMLEscapeString(label),
		)
	}

	return template.HTML(
		fmt.Sprintf(
			`<select%s>%s</select>`,
			attributesHtml,
			optionsHtml,
		),
	)
}

// {{ $mymap := makeMap "key1" "value1" "key2" "value2" }}
func MakeMap(pairs ...interface{}) (map[string]interface{}, error) {
	if len(pairs)%2 != 0 {
		return nil, errors.New("pairs must be even in number")
	}
	m := make(map[string]interface{})
	for i := 0; i < len(pairs); i += 2 {
		key, ok := pairs[i].(string)
		if !ok {
			return nil, errors.New("keys must be strings")
		}
		m[key] = pairs[i+1]
	}
	return m, nil
}

func RegisterFuncMap() {
}

func makeAttributesHtml(attributes map[string]string) string {
	attributesHtml := ""
	for key, value := range attributes {
		attributesHtml += fmt.Sprintf(
			` %s="%s"`,
			template.HTMLEscapeString(key),
			template.HTMLEscapeString(value),
		)
	}
	return attributesHtml
}

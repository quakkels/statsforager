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

// Especially useful for select options templates
// {{ $mymap := makeMap "key1" "value1" "key2" "value2" }}
func MakeMap(pairs ...string) (map[string]string, error) {
	if len(pairs)%2 != 0 {
		return nil, errors.New("pairs must be even in number")
	}
	m := make(map[string]string)
	for i := 0; i < len(pairs); i += 2 {
		key := pairs[i]
		m[key] = pairs[i+1]
	}
	return m, nil
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

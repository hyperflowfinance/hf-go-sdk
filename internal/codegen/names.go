package codegen

import (
	"strings"
	"unicode"
)

var initialisms = map[string]string{
	"api": "API", "graphql": "GraphQL", "http": "HTTP", "https": "HTTPS",
	"id": "ID", "ip": "IP", "sdk": "SDK", "url": "URL", "uri": "URI",
}

func goName(name string) string {
	var words []string
	start := 0
	runes := []rune(name)
	for i, current := range runes {
		if current == '_' {
			words = append(words, string(runes[start:i]))
			start = i + 1
		} else if i > start && unicode.IsUpper(current) &&
			(unicode.IsLower(runes[i-1]) || unicode.IsDigit(runes[i-1])) {
			words = append(words, string(runes[start:i]))
			start = i
		}
	}
	words = append(words, string(runes[start:]))
	var result strings.Builder
	for _, word := range words {
		if word == "" {
			continue
		}
		if acronym, ok := initialisms[strings.ToLower(word)]; ok {
			result.WriteString(acronym)
			continue
		}
		word = strings.ToLower(word)
		result.WriteRune(unicode.ToUpper([]rune(word)[0]))
		result.WriteString(word[1:])
	}
	return result.String()
}

func inputName(name string) string {
	result := goName(name)
	if strings.HasSuffix(result, "Input") {
		return result
	}
	return result + "Input"
}

package convert

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const (
	base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func IntToBase62(n int) string {
	if n == 0 {
		return string(base62[0])
	}

	var result []byte
	for n > 0 {
		result = append(result, base62[n%62])
		n /= 62
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func RemoveAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		return s
	}
	return output
}

func RemoveEmojis(input string) string {
	var output []rune
	for _, r := range input {
		if r <= unicode.MaxASCII {
			output = append(output, r)
		}
	}
	return string(output)
}

func Clean(input string) string {
	output := RemoveAccents(input)
	output = RemoveEmojis(output)
	return output
}

func ReplaceSlashes(input string, replace string) string {
	rule := []string{}
	rule = append(rule, "/", replace, "_", replace, "-", replace, "%", replace, "(", replace, ")", replace)
	replacer := strings.NewReplacer(rule...)
	input = replacer.Replace(input)
	return input
}

func RemoveSlashes(input string) string {
	return ReplaceSlashes(input, "")
}

func KebabCase(input string) string {
	rule := []string{}
	rule = append(rule, ".", " ", "_", " ", "-", " ")
	replacer := strings.NewReplacer(rule...)
	input = replacer.Replace(input)
	words := strings.Fields(input)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words, "-")
}

func CamelCase(input string) string {
	caser := cases.Title(language.BrazilianPortuguese)
	rule := []string{}
	rule = append(rule, ".", " ", "_", " ", "-", " ")
	replacer := strings.NewReplacer(rule...)
	input = replacer.Replace(input)
	words := strings.Fields(input)
	for i, word := range words {
		words[i] = caser.String(word)
	}
	return strings.Join(words, "")
}

func PascalCase(input string) string {
	caser := cases.Title(language.BrazilianPortuguese)
	rule := []string{}
	rule = append(rule, ".", " ", "_", " ", "-", " ")
	replacer := strings.NewReplacer(rule...)
	input = replacer.Replace(input)
	words := strings.Fields(input)
	for i, word := range words {
		words[i] = caser.String(word)
	}
	return strings.Join(words, "")
}

func SnakeCase(input string) string {
	rule := []string{}
	rule = append(rule, ".", " ", "_", " ", "-", " ")
	replacer := strings.NewReplacer(rule...)
	input = replacer.Replace(input)
	words := strings.Fields(input)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words, "_")
}

func UpperSnakeCase(input string) string {
	rule := []string{}
	rule = append(rule, ".", " ", "_", " ", "-", " ")
	replacer := strings.NewReplacer(rule...)
	input = replacer.Replace(input)
	words := strings.Fields(input)
	for i, word := range words {
		words[i] = strings.ToUpper(word)
	}
	return strings.Join(words, "_")
}

func UpperCamelCase(input string) string {
	caser := cases.Title(language.BrazilianPortuguese)
	rule := []string{}
	rule = append(rule, ".", " ", "_", " ", "-", " ")
	replacer := strings.NewReplacer(rule...)
	input = replacer.Replace(input)
	words := strings.Fields(input)
	for i, word := range words {
		words[i] = caser.String(word)
	}
	return strings.Join(words, "")
}

func LowerCamelCase(input string) string {
	caser := cases.Title(language.English)
	rule := []string{}
	rule = append(rule, ".", " ", "_", " ", "-", " ")
	replacer := strings.NewReplacer(rule...)
	input = replacer.Replace(input)
	words := strings.Fields(input)
	for i, word := range words {
		words[i] = caser.String(word)
	}
	return strings.Join(words, "")
}

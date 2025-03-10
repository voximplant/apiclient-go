package misc

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func StructToMap(s interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct but got %s", val.Kind())
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		result[field.Name] = value.Interface()
	}
	return result, nil
}

func ToCamelCase(s string) string {
	words := strings.Fields(s)
	for i := range words {
		words[i] = strings.ToUpper(string(words[i][0])) + strings.ToLower(words[i][1:])
	}
	return strings.Join(words, "")
}

func ToSnakeCase(s string) string {
	var result []string
	var word []rune
	for _, r := range s {
		if unicode.IsUpper(r) && len(word) > 0 {
			result = append(result, string(word))
			word = nil // Reset the word
		}
		word = append(word, unicode.ToLower(r)) // Always append the lowercase version
	}
	if len(word) > 0 {
		result = append(result, string(word))
	}
	return strings.Join(result, "_")
}

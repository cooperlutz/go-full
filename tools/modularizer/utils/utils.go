package utils

import (
	"path"
	"runtime"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// SnakeToPascal converts a snake_case string to PascalCase (e.g. "user_account" -> "UserAccount")
func SnakeToPascal(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = cases.Title(language.English).String(parts[i])
	}
	return strings.Join(parts, "")
}

// SnakeToFlat converts a snake_case string to flat case (e.g. "user_account" -> "UserAccount")
func SnakeToFlat(s string) string {
	return strings.ReplaceAll(s, "_", "")
}

// SnakeToKebab converts a snake_case string to kebab-case (e.g. "user_account" -> "user-account")
func SnakeToKebab(s string) string {
	return strings.ReplaceAll(s, "_", "-")
}

// SnakeToCamel converts a snake_case string to camelCase (e.g. "user_account" -> "userAccount")
func SnakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		if i == 0 {
			parts[i] = strings.ToLower(parts[i])
		} else {
			parts[i] = cases.Title(language.English).String(parts[i])
		}
	}
	return strings.Join(parts, "")
}

// FirstLetter returns the first letter of the string in lowercase (e.g. "UserAccount" -> "u")
func FirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0])
}

// GetDirectoryOfCurrentFile returns the directory of the current file
func GetDirectoryOfCurrentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if ok {
		return path.Dir(file)
	}
	return ""
}

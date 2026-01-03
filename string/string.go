// Package string provides string manipulation utilities
package stringutil

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"strings"
	"unicode"
)

// IsEmpty checks if a string is empty or contains only whitespace
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsNotEmpty checks if a string is not empty
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// Trim removes leading and trailing whitespace
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// Contains checks if a string contains a substring (case-insensitive)
func Contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// ContainsCaseSensitive checks if a string contains a substring (case-sensitive)
func ContainsCaseSensitive(s, substr string) bool {
	return strings.Contains(s, substr)
}

// StartsWith checks if a string starts with a prefix
func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// EndsWith checks if a string ends with a suffix
func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// RemovePrefix removes the prefix from a string if it exists
func RemovePrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

// RemoveSuffix removes the suffix from a string if it exists
func RemoveSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

// Reverse reverses a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CamelToSnake converts camelCase to snake_case
func CamelToSnake(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			result.WriteByte('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}

// SnakeToCamel converts snake_case to camelCase
func SnakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

// RandomString generates a random string of specified length
func RandomString(length int) (string, error) {
	bytes := make([]byte, length/2+1)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes)[:length], nil
}

// RemoveAll removes all occurrences of a substring
func RemoveAll(s, substr string) string {
	return strings.ReplaceAll(s, substr, "")
}

// Replace replaces the first n occurrences of old with new
func Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// ReplaceAll replaces all occurrences of old with new
func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// Split splits a string by separator
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// Join joins strings with a separator
func Join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

// Substring returns a substring from start to end
func Substring(s string, start, end int) string {
	if start < 0 {
		start = 0
	}
	if end > len(s) {
		end = len(s)
	}
	if start >= end {
		return ""
	}
	return s[start:end]
}

// IsNumeric checks if a string contains only numeric characters
func IsNumeric(s string) bool {
	matched, _ := regexp.MatchString(`^\d+$`, s)
	return matched
}

// IsAlpha checks if a string contains only alphabetic characters
func IsAlpha(s string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z]+$`, s)
	return matched
}

// IsAlphanumeric checks if a string contains only alphanumeric characters
func IsAlphanumeric(s string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, s)
	return matched
}

// Truncate truncates a string to a maximum length
func Truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}

// TruncateWithEllipsis truncates a string and adds ellipsis if truncated
func TruncateWithEllipsis(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}


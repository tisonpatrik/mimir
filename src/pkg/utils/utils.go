package utils

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// DecodeURL replaces encoded characters in a URL with their decoded equivalents
func DecodeURL(rawURL string) string {
	// First, unescape JSON-escaped Unicode characters
	var unescaped string
	err := json.Unmarshal([]byte(`"`+rawURL+`"`), &unescaped)
	if err != nil {
		// If unescaping fails, return the original URL
		unescaped = rawURL
	}

	// Then, decode URL-encoded characters
	decoded, err := url.QueryUnescape(unescaped)
	if err != nil {
		// If decoding fails, return the unescaped URL
		return unescaped
	}
	return decoded
}

// ToInt converts a string to an integer. If the string is empty or invalid, it returns 0.
func ToInt(value string) int {
	if value == "" {
		return 0
	}
	var result int
	fmt.Sscanf(strings.TrimSpace(value), "%d", &result)
	return result
}

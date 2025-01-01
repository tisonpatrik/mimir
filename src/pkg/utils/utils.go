package utils

import (
	"encoding/json"
	"net/url"
	"os"
)

// EnsureDir creates a directory if it doesn't exist
func EnsureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}

// SaveToFile saves a JSON-serializable object to a file
func SaveToFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")  // Pretty-print JSON
	encoder.SetEscapeHTML(false) // Disable HTML escaping

	return encoder.Encode(data)
}

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

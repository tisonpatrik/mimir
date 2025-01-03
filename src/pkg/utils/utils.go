package utils

import (
	"encoding/json"
	"net/url"
	"os"
	"regexp"
	"strings"
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

// SanitizeFileName sanitizes a string to be safe for use as a filename.
func SanitizeFileName(name string) string {
	// Replace invalid characters with underscores
	re := regexp.MustCompile(`[<>:"/\\|?*\x00-\x1F]`)
	sanitized := re.ReplaceAllString(name, "_")

	// Trim leading and trailing spaces
	sanitized = strings.TrimSpace(sanitized)

	// Optionally limit the filename length
	if len(sanitized) > 255 {
		sanitized = sanitized[:255]
	}

	return sanitized
}

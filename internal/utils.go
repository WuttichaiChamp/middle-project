package internal

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateID generates a random ID
func GenerateID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// StringPtr returns a pointer to a string
func StringPtr(s string) *string {
	return &s
}

// StringValue returns the value of a string pointer or empty string if nil
func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

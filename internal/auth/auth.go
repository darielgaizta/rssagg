package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts an API key from request headers
// e.g. Authorization: Bearer {api-key-goes-here}
func GetAPIKey(headers http.Header) (string, error) {
	authorization := headers.Get("Authorization")
	if authorization == "" {
		return "", errors.New("no authentication found.")
	}

	values := strings.Split(authorization, " ")
	if len(values) != 2 {
		return "", errors.New("malformed authentication headers.")
	}

	// Notice that Bearer is explicitly hard-coded
	if values[0] != "Bearer" {
		return "", errors.New("malformed API key in the authentication headers.")
	}

	return values[1], nil
}

package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an api key from
// the headers of an http request
// exemple:
// Authorization : ApiKey {api_key}
func GetAPIKey(headers http.Header) (string, error) {
	header := headers.Get("Authorization")
	if header == "" {
		return "", errors.New("No authentication info found")
	}
	vals := strings.Split(header, " ")
	if len(vals) != 2 {
		return "", errors.New("Invalid authentication header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("Invalid authentication header")
	}
	return vals[1], nil
}

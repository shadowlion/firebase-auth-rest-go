package firebase_auth

import (
	"net/http"
	"time"
)

type client struct {
	httpClient *http.Client
	apiKey     string
}

// NewClient returns a client for Firebase Auth REST API calls.
func NewClient(apiKey string) *client {
	return &client{
		httpClient: &http.Client{
			Timeout: time.Duration(10 * time.Second),
		},
		apiKey: apiKey,
	}
}

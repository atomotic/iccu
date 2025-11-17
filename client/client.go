package client

import (
	"context"
	"errors"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

const tokenURL = "https://api.iccu.sbn.it/oauth2/token"

// Client wraps the authenticated HTTP client for ICCU API
type Client struct {
	http *http.Client
}

// New creates a new authenticated client with the provided credentials
func New(ctx context.Context, key, secret string) (*Client, error) {
	if key == "" || secret == "" {
		return nil, errors.New("client ID and secret are required")
	}

	config := clientcredentials.Config{
		ClientID:     key,
		ClientSecret: secret,
		TokenURL:     tokenURL,
	}

	return &Client{
		http: config.Client(ctx),
	}, nil
}

// HTTP returns the underlying HTTP client for making authenticated requests
func (c *Client) HTTP() *http.Client {
	return c.http
}

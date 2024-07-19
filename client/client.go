package client

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

var client *http.Client
var tokenURL = "https://api.iccu.sbn.it/oauth2/token"

func New(key, secret string) {
	config := clientcredentials.Config{
		ClientID:     key,
		ClientSecret: secret,
		TokenURL:     tokenURL,
	}

	ctx := context.Background()
	client = config.Client(ctx)
}

func GetClient() *http.Client {
	if client == nil {
		panic("Client not initialized. Call client.New() first.")
	}
	return client
}

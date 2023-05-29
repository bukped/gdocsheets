package client

import (
	"context"
	"net/http"

	"github.com/gibs952/gdocsheets/token"
	"golang.org/x/oauth2"
)

func GetClient(config *oauth2.Config, filename string) *http.Client {
	// The file token.json stores the user's access and refresh t
	tok, err := token.TokenFromFile(filename)
	if err != nil {
		tok = token.GetTokenFromWeb(config)
		token.SaveToken(filename, tok)
	}
	return config.Client(context.Background(), tok)
}
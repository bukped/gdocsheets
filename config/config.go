package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewConfigGoogle(filePath string, scope ...string) (cfg *oauth2.Config, err error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return		
	}

	cfg, err = google.ConfigFromJSON(b, scope...)
	if err == nil {
		return
	}
	return
}
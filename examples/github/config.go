package main

import (
	"fmt"
	"os"
)

type Config struct {
	ClientID     string
	ClientSecret string
}

func NewConfig() (Config, error) {
	var cfg Config
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	if clientID == "" {
		return cfg, fmt.Errorf("GITHUB_CLIENT_ID is required")
	}
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	if clientSecret == "" {
		return cfg, fmt.Errorf("GITHUB_CLIENT_SECRET is required")
	}
	cfg = Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
	return cfg, nil
}

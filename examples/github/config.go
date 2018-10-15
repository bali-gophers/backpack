package main

import (
	"fmt"
	"os"
)

type Config struct {
	ClientID      string
	ClientSecret  string
	MysqlHost     string
	MysqlUser     string
	MysqlPassword string
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
	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost == "" {
		mysqlHost = "localhost"
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser == "" {
		mysqlUser = "raka"
	}
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	if mysqlPassword == "" {
		mysqlPassword = "raka-is-not-used"
	}
	cfg = Config{
		ClientID:      clientID,
		ClientSecret:  clientSecret,
		MysqlHost:     mysqlHost,
		MysqlUser:     mysqlUser,
		MysqlPassword: mysqlPassword,
	}
	return cfg, nil
}

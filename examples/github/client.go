package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Profile struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	AvatarURL   string    `json:"avatar_url"`
	Bio         string    `json:"bio"`
	PublicRepos int       `json:"public_repos"`
	CreatedAt   time.Time `json:"created_at"`
}

type AccessTokenRequestModel struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

type AccessTokenResponseModel struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func (m AccessTokenResponseModel) String() string {
	return fmt.Sprintf("%s %s", m.TokenType, m.AccessToken)
}

type Client struct {
	clientID     string
	clientSecret string
	scopes       []string

	apiBaseURL       string
	accessTokenURL   string
	authorizationURL string
}

func NewClient(cfg Config) *Client {
	return &Client{
		clientID:         cfg.ClientID,
		clientSecret:     cfg.ClientSecret,
		scopes:           []string{"public_repo"},
		apiBaseURL:       "https://api.github.com",
		accessTokenURL:   "https://github.com/login/oauth/access_token",
		authorizationURL: "https://github.com/login/oauth/authorize",
	}
}

func (c *Client) GetAuthorizationURL() (string, error) {
	params := []string{
		"client_id=" + c.clientID,
		"client_secret=" + c.clientSecret,
		"allow_signup=false",
		"scope=" + strings.Join(c.scopes, ","),
	}
	return fmt.Sprintf("%s?%s", c.authorizationURL, strings.Join(params, "&")), nil
}

func (c *Client) GetAccessToken(code string) (res AccessTokenResponseModel, err error) {
	model := AccessTokenRequestModel{
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
		Code:         code,
	}
	modelByte, err := json.Marshal(model)
	if err != nil {
		log.Printf("[ERROR] %#v \n", err)
		return res, err
	}
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}
	req, err := http.NewRequest("POST", c.accessTokenURL, bytes.NewBuffer(modelByte))
	if err != nil {
		log.Printf("[ERROR] %#v \n", err)
		return res, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] %#v \n", err)
		return res, err
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(responseData, &res)
	if err != nil {
		log.Printf("[ERROR] %#v \n", err)
		return res, err
	}
	return res, nil
}

func (c *Client) GetAuthenticated(token string) (res Profile, err error) {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user", c.apiBaseURL), nil)
	if err != nil {
		log.Printf("[ERROR] %#v \n", err)
		return res, err
	}
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] %#v \n", err)
		return res, err
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[ERROR] %#v \n", err)
		return res, err
	}
	err = json.Unmarshal(respData, &res)
	if err != nil {
		log.Printf("[ERROR] %#v \n", err)
		return res, err
	}
	return res, nil
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	HelloTemp = `
Hello %s, you have awesome %d public repos
	`
)

type Handler struct {
	Config Config
	Client *Client
}

func NewHandler(cfg Config, client *Client) Handler {
	return Handler{
		Config: cfg,
		Client: client,
	}
}

func (h Handler) Home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Welcome"))
}

func (h Handler) Health(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Pong!"))
}

func (h Handler) Auth(w http.ResponseWriter, req *http.Request) {
	url, err := h.Client.GetAuthorizationURL()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	log.Printf("Redirecting to %s \n", url)
	http.Redirect(w, req, url, 301)
}

func (h Handler) Callback(w http.ResponseWriter, req *http.Request) {
	queryValues := req.URL.Query()
	codes, exists := queryValues["code"]
	if !exists {
		w.Write([]byte("Code couldn't be found"))
		return
	}
	res, err := h.Client.GetAccessToken(codes[0])
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	profile, err := h.Client.GetAuthenticated(res.String())
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(fmt.Sprintf(HelloTemp, profile.Name, profile.PublicRepos)))
}

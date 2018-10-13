# Github authentication page

## Pre-installation
- Create Github OAuth apps, go to `https://github.com/settings/developers` with following config
```
Authorization callback URL = http://localhost:9000/auth/callback

```
- Keep client ID & client Secret, it will use later

## Installation
- Run `go get github.com/bali-gophers/backpack`
- Go to `examples/github` directory
- Run `go build -o goBinary`

## How to run
- Run `go build -o goBinary`
- Setup environment variable, provide `GITHUB_CLIENT_ID` & `GITHUB_CLIENT_SECRET` that you get before when creating Github Oauth apps
- Run `./goBinary`
- Open `localhost:9000/auth` in your favorite browser
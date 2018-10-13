# Github authentication page

## Pre-installation
- Create Github OAuth apps, go to `https://github.com/settings/developers` with following config
```
Authorization callback URL = http://localhost:9000/auth/callback

```
- Keep client ID & client Secret, it will be used later

## Installation
- Run `go get github.com/bali-gophers/backpack`
- Go to `examples/github` directory
- Run `make build`

## How to run
- Setup environment variables by putting the values into `.env` file, see `env.example` for the examples
- Run `make run`
- Open `localhost:9000/auth` in your favorite browser
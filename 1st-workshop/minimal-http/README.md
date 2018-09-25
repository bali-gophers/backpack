# Minimal http application

## Installation
```
$ go build -o minimal-server main.go
$ ./minimal-server
```

## Endpoints
- /hello -> render plain hello message
- /helloHtml -> render static html
- /helloVars -> render dynamic html (values from memory)
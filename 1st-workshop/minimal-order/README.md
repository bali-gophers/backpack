# Minimal Go program with Database access

## Requirements
- Go
- MySQL

## Installation
```
$ go get .
$ mysql -h localhost -u root -p < migration.sql
$ go build -o minimal-order main.go
$ ./minimal-order
```

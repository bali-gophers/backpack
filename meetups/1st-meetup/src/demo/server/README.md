# Order mini-application
This is just a small application to demonstrate how to write http application using Go.

## Requirements
- Go
- MySQL

## Installation
```
$ go get .
$ mysql -h localhost -u root -p < migration.sql
$ go build -o order-server
$ ./order-server
```

## Endpoints
- [POST]    /order
```
Request Body:
{
    "fullName": "Made Raka",
    "email": "maderakateja@gmail.com",
    "items": [
        {
            "title": "Bunga Gemitir",
            "count": 3,
            "price": 10000
        },
        {
            "title": "Ceper",
            "count": 5,
            "price": 50000
        }
    ]
}
```
- [GET]     /order/:orderId
- [GET]     /order/:orderId/html

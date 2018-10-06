package main

import (
	"errors"
	"fmt"
	"time"
)

type Order struct {
	OrderID     int64     `json:"orderId"`
	OrderNumber string    `json:"orderNumber"`
	FullName    string    `json:"fullName"`
	Email       string    `json:"email"`
	Total       int64     `json:"total"`
	Items       []Item    `json:"items"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (o Order) Validate() (Order, error) {
	if o.FullName == "" {
		return Order{}, errors.New("fullName is required")
	}
	if o.Email == "" {
		return Order{}, errors.New("email is required")
	}
	if len(o.Items) == 0 {
		return Order{}, errors.New("order items is required")
	}
	for _, item := range o.Items {
		err := item.Validate()
		if err != nil {
			return Order{}, err
		}
	}
	var total int64
	for _, item := range o.Items {
		total += (item.Price * int64(item.Count))
	}
	now := time.Now()
	o.OrderNumber = fmt.Sprintf("%d%d%d%d", now.Day(), now.Hour(), now.Minute(), now.Second())
	o.Total = total
	o.CreatedAt = time.Now()
	return o, nil
}

type Item struct {
	ItemID  int64  `json:"itemId"`
	OrderID int64  `json:"orderId"`
	Title   string `json:"title"`
	Count   int    `json:"count"`
	Price   int64  `json:"price"`
}

func (i Item) Validate() error {
	if i.Title == "" {
		return errors.New("title is required")
	}
	if i.Count <= 0 {
		return errors.New("count couldn't be zero")
	}
	if i.Price <= 0 {
		return errors.New("price couldn't be zero")
	}
	return nil
}

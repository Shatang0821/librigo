package model

import "fmt"

type Book struct {
	ID        int
	Title     string
	Author    string
	Price     int
	CreatedAt string
}

func NewBook(title, author string, price int) (*Book, error) {
	if title == "" || author == "" || price < 0 {
		return nil, fmt.Errorf("invalid book parameters")
	}
	return &Book{
		Title:  title,
		Author: author,
		Price:  price,
	}, nil
}

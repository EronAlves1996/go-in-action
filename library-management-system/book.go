package main

import "fmt"

type Book struct {
	ISBN       string
	Title      string
	Author     string
	CheckedOut bool
}

func (b *Book) CheckOut() {
	b.CheckedOut = true
}

func (b *Book) Return() {
	b.CheckedOut = false
}

func (b Book) String() string {
	return fmt.Sprintf("%s by %s (ISBN: %s)", b.Title, b.Author, b.ISBN)
}

func (b Book) Description() string {
	return fmt.Sprintf("{title: %s, author: %s}", b.Title, b.Author)
}

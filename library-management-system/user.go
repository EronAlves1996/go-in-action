package main

import "fmt"

type User struct {
	ID            int
	Name          string
	BorrowedBooks []*Book
}

func (u *User) Borrow(b *Book) {
	u.BorrowedBooks = append(u.BorrowedBooks, b)
}

func (u *User) ReturnBook(b *Book) {
	i := -1
	for idx, v := range u.BorrowedBooks {
		if v == b {
			i = idx
			break
		}
	}
	if i != -1 {
		u.BorrowedBooks = append(u.BorrowedBooks[:i], u.BorrowedBooks[i+1:]...)
	}
}

func (u User) ListBorrowed() {
	for _, b := range u.BorrowedBooks {
		fmt.Println(b.String())
	}
}

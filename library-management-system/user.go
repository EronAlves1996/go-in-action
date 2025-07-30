package main

import "fmt"

type User struct {
	ID            int
	Name          string
	BorrowedBooks []*Book
}

func (u *User) Borrow(b *Book) {
	if b.CheckedOut {
		return
	}
	b.CheckOut()
	u.BorrowedBooks = append(u.BorrowedBooks, b)
}

func (u *User) ReturnBook(b *Book) {
	if !b.CheckedOut {
		return
	}

	i := -1
	for idx, v := range u.BorrowedBooks {
		if v == b {
			i = idx
			break
		}
	}
	if i != -1 {
		b.Return()
		u.BorrowedBooks = append(u.BorrowedBooks[:i], u.BorrowedBooks[i+1:]...)
	}
}

func (u User) ListBorrowed() {
	if len(u.BorrowedBooks) == 0 {
		fmt.Println("empty")
	}
	for _, b := range u.BorrowedBooks {
		fmt.Println(b.String())
	}
}

func (u User) Description() string {
	return fmt.Sprintf("{name: %s, books borrowed: %d}", u.Name, len(u.BorrowedBooks))
}

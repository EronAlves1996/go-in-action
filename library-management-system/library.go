package main

type Library struct {
	Books map[string]*Book
	Users map[int]*User
}

func (l *Library) AddBook(b *Book) {
	l.Books[b.ISBN] = b
}

func (l *Library) AddUser(u *User) {
	l.Users[u.ID] = u
}

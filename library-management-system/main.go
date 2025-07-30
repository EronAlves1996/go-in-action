package main

import "fmt"

func PrintDescription(item Describable) {
	fmt.Println(item.Description())
}

func main() {
	book := Book{
		ISBN:   "12121412",
		Title:  "Morro dos ventos uivantes",
		Author: "José da Silva",
	}
	book2 := Book{
		ISBN:   "121214122",
		Title:  "Biografia Marilia mendonça",
		Author: "Murilo Huff",
	}
	user := User{
		ID:            1,
		Name:          "Camus",
		BorrowedBooks: []*Book{},
	}
	library := Library{
		Books: make(map[string]*Book),
		Users: make(map[int]*User),
	}
	library.AddBook(&book)
	library.AddBook(&book2)
	library.AddUser(&user)

	user.ListBorrowed()
	user.Borrow(&book)
	user.ListBorrowed()
	user.ReturnBook(&book)
	user.ListBorrowed()
}

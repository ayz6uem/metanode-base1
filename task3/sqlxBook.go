package task3

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var schemaBook = `
CREATE TABLE IF NOT EXISTS books (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    price integer NOT NULL
)`

func SqlxBook() {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	db.MustExec(schemaBook)
	db.MustExec(`INSERT INTO books(id, title, author, price) VALUES (1, '钢铁是怎样练成的', 'jack', 49)`)
	db.MustExec(`INSERT INTO books(id, title, author, price) VALUES (2, 'golang开发指南', 'hary', 50)`)
	db.MustExec(`INSERT INTO books(id, title, author, price) VALUES (3, 'sql入门', 'john', 51)`)

	var books []Book
	db.Select(&books, "select * from books where price>$1", 50)
	fmt.Println(books)
}

type Book struct {
	ID     int64
	Title  string
	Author string
	Price  int
}

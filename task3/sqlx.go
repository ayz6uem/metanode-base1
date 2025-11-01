package task3

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE IF NOT EXISTS employees (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    department VARCHAR(255) NOT NULL,
    salary integer NOT NULL
)`

func Sqlx() {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	db.MustExec(schema)
	db.MustExec(`INSERT INTO employees(id, name, department, salary) VALUES (1, 'jack', '技术部', 10000)`)
	db.MustExec(`INSERT INTO employees(id, name, department, salary) VALUES (2, 'hary', '技术部', 12000)`)
	db.MustExec(`INSERT INTO employees(id, name, department, salary) VALUES (3, 'john', '人事部', 9000)`)

	var employees []Employee
	db.Select(&employees, "select * from employees where department=$1", "技术部")
	fmt.Println(employees)

	var emp Employee
	db.Get(&emp, "select * from employees order by salary desc limit 1")
	fmt.Println(emp)
}

type Employee struct {
	ID         int64
	Name       string
	Department string
	Salary     int
}

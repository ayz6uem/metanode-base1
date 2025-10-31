package task2

import (
	"fmt"
	"math"
)

func Object() {
	rect := Rectangle{Width: 3, Height: 4}
	fmt.Printf("task2 面向对象 rect area %v perimeter %v\n", rect.Area(), rect.Perimeter())
	ci := Circle{radius: 5}
	fmt.Printf("task2 面向对象 circle %v perimeter %v\n", ci.Area(), ci.Perimeter())

	e := Employee{Person: Person{Name: "wz", Age: 18}, EmployeeId: 1001}
	e.PrintInfo()
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * c.radius * math.Pi
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeId int
}

func (e Employee) PrintInfo() {
	fmt.Println("task2 面向对象 ", e)
}

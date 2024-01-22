package entity

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Firstname string
	Lastname  string
	Tel       string
	Email     string
	Username  string
	Password  string

	PositionID uint
	Position *Position 

	PrecedeID uint
	Precede *Precede

	GenderID uint
	Gender  *Gender

}

type Precede struct {
	gorm.Model
	Name string

	Employee []Employee
}

type Position struct {
	gorm.Model
	Name        string
	Description string
	Salary      string

	Employee []Employee
}

type Gender struct {
	gorm.Model
	Name string

	Employee []Employee
	User []User
}

type User struct {
	gorm.Model
	Firstname string
	Lastname  string
	Age       int
	Phone     string
	Username  string
	Email     string
	Password  string

	GenderID uint
	Gender *Gender

	RoleID uint
	Role *Role

	Payment []Payment

}

type Role struct {
	gorm.Model
	Name string

	User []User
}

type Ticket struct {
	gorm.Model
	Price  string
	seat   int
	Detail string

	CarID uint
	Car *Car

	DepartureID uint
	Departure *Departure

	Payment []Payment
}

type Departure struct {
	gorm.Model
	DepsrtureStation  string
	DestinationStaion string
	Date              time.Time

	Ticket []Ticket
}

type Car struct {
	gorm.Model
	Name     string
	CarModel string
	Route    string

	Ticket []Ticket
}

type Payment struct {
	gorm.Model
	Total  int
	
	TicketID uint
	Ticket *Ticket

	UserID uint
	User *User
}


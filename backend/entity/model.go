package entity

//"time"

type Employee struct {
	gorm.model
	Firstname string
	Lastname  string
	Tel       string
	Email     string
	Username  string
	Password  string
}

type User struct {
	gorm.model
	Firstname string
	Lastname  string
	Tel       string
	Email     string
	Username  string
	Password  string
	Details   string

}

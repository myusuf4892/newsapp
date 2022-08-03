package users

import "time"

type Core struct {
	ID        int
	FullName  string
	Phone     string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Auth struct {
	ID       int
	FullName string
	Phone    string
	Email    string
}

type Business interface {
	Auth(dataUser Core) (token, fullName string, userID int, err error)
	AddUser(dataUser Core) (err error)
}

type Data interface {
	FindUser(email string) (response Core, err error)
	Insert(dataUser Core) (err error)
}

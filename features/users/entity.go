package users

import "time"

type Core struct {
	ID       int
	FullName string
	Phone    string
	Email    string
	Password string
	CreateAt time.Time
	UpdateAt time.Time
}

type Business interface {
	AddUser(dataUser Core) (err error)
}

type Data interface {
	Insert(dataUser Core) (err error)
}

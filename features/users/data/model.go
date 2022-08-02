package data

import (
	"news/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string
	Phone    string
	Email    string
	Password string
}

func (dataUser *User) toCore() users.Core {
	return users.Core{
		ID:       int(dataUser.ID),
		FullName: dataUser.FullName,
		Phone:    dataUser.Phone,
		Email:    dataUser.Email,
		Password: dataUser.Password,
	}
}

func fromCore(dataUser users.Core) User {
	return User{
		FullName: dataUser.FullName,
		Phone:    dataUser.Phone,
		Email:    dataUser.Email,
		Password: dataUser.Password,
	}
}

func toCore(dataUser User) users.Core {
	return dataUser.toCore()
}

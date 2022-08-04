package data

import (
	"news/features/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName  string
	Phone     string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (dataUser *User) toCore() users.Core {
	return users.Core{
		ID:        int(dataUser.ID),
		FullName:  dataUser.FullName,
		Phone:     dataUser.Phone,
		Email:     dataUser.Email,
		Password:  dataUser.Password,
		CreatedAt: dataUser.CreatedAt,
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

// func toCore(dataUser User) users.Core {
// 	return dataUser.toCore()
// }

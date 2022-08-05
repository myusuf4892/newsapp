package data

import (
	"encoding/json"
	"errors"
	"news/features/users"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository) Insert(dataReq users.Core) (err error) {
	data := fromCore(dataReq)
	var newError map[string]interface{}

	srv := repo.db.Create(&data)

	if srv.Error != nil {
		errByte, _ := json.Marshal(srv.Error)
		json.Unmarshal((errByte), &newError)

		if newError["Number"] == float64(1062) {
			return errors.New("email already used")
		}
		return srv.Error
	}
	return srv.Error
}

func (repo *mysqlUserRepository) FindUser(email string) (response users.Core, err error) {
	dataUser := User{}
	srv := repo.db.Where("email = ?", email).Find(&dataUser)
	if srv.Error != nil {
		return users.Core{}, errors.New("no data user")
	}
	return dataUser.toCore(), srv.Error
}

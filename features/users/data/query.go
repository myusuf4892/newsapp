package data

import (
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

	srv := repo.db.Create(&data)

	if srv.Error != nil {
		return errors.New("error server")
	}
	return nil
}

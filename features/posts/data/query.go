package data

import (
	"errors"
	"news/features/posts"

	"gorm.io/gorm"
)

type mysqlArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(conn *gorm.DB) posts.Data {
	return &mysqlArticleRepository{
		db: conn,
	}
}

func (repo *mysqlArticleRepository) Insert(dataReq posts.Core) (err error) {
	data := fromCore(&dataReq)

	srv := repo.db.Create(data)
	if srv.Error != nil {
		return errors.New("error server")
	}
	return nil
}

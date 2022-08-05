package data

import (
	"errors"
	"news/features/posts"

	"gorm.io/gorm"
)

type mysqlPostRepository struct {
	db *gorm.DB
}

func NewPostRepository(conn *gorm.DB) posts.Data {
	return &mysqlPostRepository{
		db: conn,
	}
}

func (repo *mysqlPostRepository) Insert(dataReq posts.Core) (row int, err error) {
	data := fromCore(&dataReq)

	srv := repo.db.Create(&data)
	if srv.Error != nil {
		return int(srv.RowsAffected), errors.New("error server")
	}
	return int(srv.RowsAffected), srv.Error
}

func (repo *mysqlPostRepository) Get() (response []posts.Core, err error) {
	data := []Post{}

	srv := repo.db.Preload("PostType").Preload("User").Find(&data)
	if srv.Error != nil {
		return []posts.Core{}, srv.Error
	}

	response = ToCoreList(data)

	return response, nil
}

func (repo *mysqlPostRepository) Update(dataReq posts.Core, ID int) (row int, err error) {
	Model := Post{}
	Model.ID = uint(ID)

	srv := repo.db.Model(&Model).Updates(&Post{
		Tittle:      dataReq.Tittle,
		Description: dataReq.Description,
		PostTypeID:  dataReq.PostTypeID,
		UserID:      dataReq.UserID,
	})
	if srv.Error != nil {
		return int(srv.RowsAffected), errors.New("error server")
	}

	return int(srv.RowsAffected), srv.Error
}

func (repo *mysqlPostRepository) Destroy(ID int) (row int, err error) {
	Model := Post{}

	srv := repo.db.Delete(&Model, ID)
	if srv.Error != nil {
		return int(srv.RowsAffected), errors.New("error server")
	}

	return int(srv.RowsAffected), srv.Error
}

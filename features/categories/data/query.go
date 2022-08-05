package data

import (
	"errors"
	"news/features/categories"

	"gorm.io/gorm"
)

type mysqlCategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(conn *gorm.DB) categories.Data {
	return &mysqlCategoryRepository{
		db: conn,
	}
}

func (repo *mysqlCategoryRepository) Insert(dataCategory categories.Core) (row int, err error) {
	data := fromPostType(&dataCategory)

	srv := repo.db.Create(&data)
	if srv.Error != nil {
		return int(srv.RowsAffected), errors.New("error server")
	}
	return int(srv.RowsAffected), srv.Error
}

func (repo *mysqlCategoryRepository) Get() (response []categories.Core, err error) {
	data := []PostType{}

	srv := repo.db.Find(&data)
	if srv.Error != nil {
		return []categories.Core{}, errors.New("no data post type")
	}

	response = ToCoreList(data)

	return response, srv.Error
}

func (repo *mysqlCategoryRepository) Update(dataReq categories.Core, ID int) (row int, err error) {
	Model := PostType{}
	Model.ID = uint(ID)
	srv := repo.db.Model(&Model).Updates(&PostType{
		Category: dataReq.Category,
		Type:     dataReq.Type,
	})
	if srv.Error != nil {
		return int(srv.RowsAffected), errors.New("error server")
	}

	return int(srv.RowsAffected), srv.Error
}

func (repo *mysqlCategoryRepository) Destroy(ID int) (row int, err error) {
	Model := PostType{}
	srv := repo.db.Delete(&Model, ID)
	if srv.Error != nil {
		return int(srv.RowsAffected), errors.New("error server")
	}

	return int(srv.RowsAffected), srv.Error
}

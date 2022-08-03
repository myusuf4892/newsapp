package business

import (
	"errors"
	"news/features/categories"
)

type categoryUseCase struct {
	categoryData categories.Data
}

func NewCategoryBusiness(dataPost categories.Data) categories.Business {
	return &categoryUseCase{
		categoryData: dataPost,
	}
}

func (uc *categoryUseCase) AddCategories(dataReq categories.Core) (err error) {
	res := uc.categoryData.Insert(dataReq)
	if res != nil {
		return errors.New("error server")
	}
	return res
}

func (uc *categoryUseCase) GetCategories() (response []categories.Core, err error) {
	res, err := uc.categoryData.Get()
	if err != nil {
		return nil, errors.New("error server")
	}

	return res, nil
}

func (uc *categoryUseCase) UpdateCategories(dataReq categories.Core, ID int) (err error) {
	res := uc.categoryData.Update(dataReq, ID)
	if res != nil {
		return errors.New("error server")
	}

	return nil
}

func (uc *categoryUseCase) DestroyCategories(ID int) (err error) {
	res := uc.categoryData.Destroy(ID)
	if res != nil {
		return errors.New("error server")
	}

	return nil
}

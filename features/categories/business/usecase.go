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

func (uc *categoryUseCase) AddCategories(dataReq categories.Core) (row int, err error) {
	if dataReq.Category == "" || dataReq.Type == "" {
		return 0, errors.New("all must be filled")
	}
	res, err := uc.categoryData.Insert(dataReq)
	if err != nil {
		return 0, errors.New("error server")
	}
	if res == 0 {
		return 0, err
	}

	return res, nil
}

func (uc *categoryUseCase) GetCategories() (response []categories.Core, err error) {
	res, err := uc.categoryData.Get()
	if err != nil {
		return nil, errors.New("error server")
	}

	return res, err
}

func (uc *categoryUseCase) UpdateCategories(dataReq categories.Core, ID int) (row int, err error) {
	if dataReq.Category == "" || dataReq.Type == "" {
		return 0, errors.New("all must be filled")
	}
	res, err := uc.categoryData.Update(dataReq, ID)
	if err != nil {
		return 0, errors.New("error server")
	}
	if res == 0 {
		return 0, err
	}

	return res, nil
}

func (uc *categoryUseCase) DestroyCategories(ID int) (row int, err error) {
	res, err := uc.categoryData.Destroy(ID)
	if err != nil {
		return 0, errors.New("error server")
	}
	if res == 0 {
		return 0, err
	}

	return res, nil
}

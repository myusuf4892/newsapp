package business

import (
	"errors"
	"news/features/posts"
)

type postUseCase struct {
	postData posts.Data
}

func NewPostBusiness(dataPost posts.Data) posts.Business {
	return &postUseCase{
		postData: dataPost,
	}
}

func (uc *postUseCase) AddPost(dataReq posts.Core) (row int, err error) {
	if dataReq.Tittle == "" || dataReq.Description == "" {
		return 0, errors.New("all must be filled")
	}
	res, err := uc.postData.Insert(dataReq)
	if err != nil {
		return 0, errors.New("error server")
	}
	if res == 0 {
		return 0, err
	}
	return res, nil
}

func (uc *postUseCase) GetPost() (response []posts.Core, err error) {
	res, err := uc.postData.Get()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *postUseCase) UpdatePost(dataReq posts.Core, ID int) (row int, err error) {
	res, err := uc.postData.Update(dataReq, ID)
	if err != nil {
		return 0, errors.New("error server")
	}
	if res == 0 {
		return 0, err
	}
	return res, nil
}

func (uc *postUseCase) DestroyPost(ID int) (row int, err error) {
	res, err := uc.postData.Destroy(ID)
	if err != nil {
		return 0, errors.New("error server")
	}
	if res == 0 {
		return 0, err
	}
	return res, nil
}

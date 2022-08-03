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

func (uc *postUseCase) AddPost(dataReq posts.Core) (err error) {
	res := uc.postData.Insert(dataReq)
	if res != nil {
		return errors.New("error server")
	}
	return res
}

func (uc *postUseCase) GetPost() (response []posts.Core, err error) {
	res, err := uc.postData.Get()
	if err != nil {
		return nil, errors.New("error server")
	}

	return res, nil
}

func (uc *postUseCase) UpdatePost(dataReq posts.Core, ID int) (err error) {
	res := uc.postData.Update(dataReq, ID)
	if res != nil {
		return errors.New("error server")
	}

	return nil
}

func (uc *postUseCase) DestroyPost(ID int) (err error) {
	res := uc.postData.Destroy(ID)
	if res != nil {
		return errors.New("error server")
	}

	return nil
}

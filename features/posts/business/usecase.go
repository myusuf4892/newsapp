package business

import (
	"errors"
	"news/features/posts"
)

type postUseCase struct {
	userData posts.Data
}

func NewArticleBusiness(dataPost posts.Data) posts.Business {
	return &postUseCase{
		userData: dataPost,
	}
}

func (uc *postUseCase) AddArticles(dataReq posts.Core) (err error) {
	res := uc.userData.Insert(dataReq)
	if res != nil {
		return errors.New("error server")
	}
	return res
}

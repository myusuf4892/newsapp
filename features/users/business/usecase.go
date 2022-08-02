package business

import (
	"errors"
	"news/features/users"
)

type userUseCase struct {
	userData users.Data
}

func NewUserBusiness(dataUser users.Data) users.Business {
	return &userUseCase{
		userData: dataUser,
	}
}

func (uc *userUseCase) AddUser(dataReq users.Core) (err error) {
	res := uc.userData.Insert(dataReq)
	if res != nil {
		return errors.New("error server")
	}
	return res
}

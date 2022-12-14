package business

import (
	"errors"
	_middlewares "news/app/middlewares"
	"news/features/users"

	"golang.org/x/crypto/bcrypt"
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
	if dataReq.FullName == "" || dataReq.Phone == "" || dataReq.Email == "" || dataReq.Password == "" {
		return errors.New("all data must be filled")
	}

	errEmailFormat := emailFormatValidation(dataReq.Email)
	if errEmailFormat != nil {
		return errors.New(errEmailFormat.Error())
	}

	passWillBcrypt := []byte(dataReq.Password)
	hash, _ := bcrypt.GenerateFromPassword(passWillBcrypt, bcrypt.DefaultCost)

	dataReq.Password = string(hash)

	res := uc.userData.Insert(dataReq)
	if res != nil {
		return errors.New("error server")
	}
	return nil
}

func (uc *userUseCase) Auth(dataReq users.Core) (token, fullName string, userID int, err error) {
	res, errRes := uc.userData.FindUser(dataReq.Email)
	if errRes != nil {
		return "", "", 0, errors.New("no data user")
	}
	compare := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(dataReq.Password))
	if compare != nil {
		return "", "", 0, errors.New("wrong password")
	}
	token, err = _middlewares.CreateToken(int(res.ID))
	if err != nil {
		return "", "", 0, errors.New("can't create token")
	}
	return token, res.FullName, res.ID, nil
}

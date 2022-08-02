package app

import (
	_userBusiness "news/features/users/business"
	_userData "news/features/users/data"
	_userPresentation "news/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter *_userPresentation.UserHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {

	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	return Presenter{
		UserPresenter: userPresentation,
	}
}

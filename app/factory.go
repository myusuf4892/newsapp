package app

import (
	_userBusiness "news/features/users/business"
	_userData "news/features/users/data"
	_userPresentation "news/features/users/delivery"

	_articleBusiness "news/features/posts/business"
	_articleData "news/features/posts/data"
	_articlePresentation "news/features/posts/delivery"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter    *_userPresentation.UserHandler
	ArticlePresenter *_articlePresentation.ArticleHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {

	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	articleData := _articleData.NewArticleRepository(dbConn)
	articleBusiness := _articleBusiness.NewArticleBusiness(articleData)
	articlePresentation := _articlePresentation.NewArticleHandler(articleBusiness)

	return Presenter{
		UserPresenter:    userPresentation,
		ArticlePresenter: articlePresentation,
	}
}

package app

import (
	_userBusiness "news/features/users/business"
	_userData "news/features/users/data"
	_userPresentation "news/features/users/delivery"

	_postBusiness "news/features/posts/business"
	_postData "news/features/posts/data"
	_postPresentation "news/features/posts/delivery"

	_categoryBusiness "news/features/categories/business"
	_categoryData "news/features/categories/data"
	_categoryPresentation "news/features/categories/delivery"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter     *_userPresentation.UserHandler
	PostPresenter     *_postPresentation.PostHandler
	CategoryPresenter *_categoryPresentation.CategoryHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {

	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	categoryData := _categoryData.NewCategoryRepository(dbConn)
	categoryBusiness := _categoryBusiness.NewCategoryBusiness(categoryData)
	categoryPresentation := _categoryPresentation.NewCategoryHandler(categoryBusiness)

	postData := _postData.NewPostRepository(dbConn)
	postBusiness := _postBusiness.NewPostBusiness(postData)
	postPresentation := _postPresentation.NewPostHandler(postBusiness)

	return Presenter{
		UserPresenter:     userPresentation,
		PostPresenter:     postPresentation,
		CategoryPresenter: categoryPresentation,
	}
}

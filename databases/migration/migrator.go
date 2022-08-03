package migration

import (
	_mCategory "news/features/categories/data"
	_mPost "news/features/posts/data"
	_mUser "news/features/users/data"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(_mUser.User{})
	db.AutoMigrate(_mCategory.PostType{})
	db.AutoMigrate(_mPost.Post{})
}

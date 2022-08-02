package migration

import (
	_mPost "news/features/posts/data"
	_mUser "news/features/users/data"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(_mUser.User{})
	db.AutoMigrate(_mPost.PostType{})
	db.AutoMigrate(_mPost.Post{})
}

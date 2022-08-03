package data

import (
	_category "news/features/categories/data"
	"news/features/posts"
	_user "news/features/users/data"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Tittle      string
	Description string
	PostTypeID  int
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PostType    _category.PostType
	User        _user.User
}

func (post *Post) toCore() posts.Core {
	return posts.Core{
		ID:          int(post.ID),
		Tittle:      post.Tittle,
		Description: post.Description,
		PostTypeID:  post.PostTypeID,
		UserID:      post.UserID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		PostType: posts.PostType{
			Category: post.PostType.Category,
			Type:     post.PostType.Type,
		},
		User: posts.User{
			FullName: post.User.FullName,
		},
	}
}

func ToCoreList(data []Post) []posts.Core {
	result := []posts.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(post *posts.Core) Post {
	return Post{
		Tittle:      post.Tittle,
		Description: post.Description,
		PostTypeID:  post.PostTypeID,
		UserID:      post.UserID,
		PostType: _category.PostType{
			Category: post.PostType.Category,
			Type:     post.PostType.Type,
		},
		User: _user.User{
			FullName: post.User.FullName,
		},
	}
}

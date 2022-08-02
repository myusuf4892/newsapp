package data

import (
	"news/features/posts"
	"news/features/users/data"
	"time"

	"gorm.io/gorm"
)

type PostType struct {
	gorm.Model
	Jenis string
	Type  string
	Post  Post
}

type Post struct {
	gorm.Model
	Tittle      string
	Description string
	PostTypeID  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        data.User
}

func (post *Post) toCore() posts.Core {
	return posts.Core{
		ID:          int(post.ID),
		Tittle:      post.Tittle,
		Description: post.Description,
		PostTypeID:  post.PostTypeID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}
}

func fromCore(post *posts.Core) Post {
	return Post{
		Tittle:      post.Tittle,
		Description: post.Description,
		PostTypeID:  post.PostTypeID,
	}
}

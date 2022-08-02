package posts

import "time"

type Core struct {
	ID          int
	Tittle      string
	Description string
	PostTypeID  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User
}

type User struct {
	ID       int
	FullName string
	Phone    string
	Email    string
}

type Business interface {
	AddArticles(dataArticles Core) (err error)
}

type Data interface {
	Insert(dataArticles Core) (err error)
}

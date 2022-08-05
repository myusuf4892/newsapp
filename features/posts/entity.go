package posts

import "time"

type Core struct {
	ID          int
	Tittle      string
	Description string
	PostTypeID  int
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PostType    PostType
	User        User
}

type PostType struct {
	ID       int
	Category string
	Type     string
}

type User struct {
	ID       int
	FullName string
	Phone    string
	Email    string
}

type Business interface {
	AddPost(dataPost Core) (row int, err error)
	GetPost() (dataPost []Core, err error)
	UpdatePost(dataPost Core, ID int) (row int, err error)
	DestroyPost(ID int) (row int, err error)
}

type Data interface {
	Insert(dataPost Core) (row int, err error)
	Get() (dataPost []Core, err error)
	Update(dataPost Core, ID int) (row int, err error)
	Destroy(ID int) (row int, err error)
}

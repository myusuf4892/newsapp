package request

import "news/features/posts"

type Post struct {
	Tittle      string
	Description string
	PostTypeID  int
}

func ToCore(dataReq Post) posts.Core {
	core := posts.Core{
		Tittle:      dataReq.Tittle,
		Description: dataReq.Description,
		PostTypeID:  dataReq.PostTypeID,
	}
	return core
}

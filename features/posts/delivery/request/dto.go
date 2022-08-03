package request

import "news/features/posts"

type Post struct {
	Tittle      string `json:"tittle" form:"tittle"`
	Description string `json:"description" form:"description"`
	PostTypeID  int    `json:"post_type_id" form:"post_type_id"`
}

func ToCore(dataReq Post) posts.Core {
	core := posts.Core{
		Tittle:      dataReq.Tittle,
		Description: dataReq.Description,
		PostTypeID:  dataReq.PostTypeID,
	}
	return core
}

package response

import "news/features/posts"

type Post struct {
	ID          int      `json:"post_id"`
	Tittle      string   `json:"tittle"`
	Description string   `json:"description"`
	PostTypeID  int      `json:"post_type_id"`
	UserID      int      `json:"user_id"`
	User        User     `json:"user"`
	PostType    PostType `json:"post_type"`
}

type PostType struct {
	Category string `json:"category"`
	Type     string `json:"type"`
}

type User struct {
	FullName string `json:"author"`
}

func FromCore(post posts.Core) Post {
	return Post{
		ID:          post.ID,
		Tittle:      post.Tittle,
		Description: post.Description,
		PostTypeID:  post.PostTypeID,
		UserID:      post.UserID,
		PostType: PostType{
			Category: post.PostType.Category,
			Type:     post.PostType.Type,
		},
		User: User{
			FullName: post.User.FullName,
		},
	}
}

func FromCoreToList(dataReq []posts.Core) []Post {
	result := []Post{}
	for key := range dataReq {
		result = append(result, FromCore(dataReq[key]))
	}
	return result
}

package data

import (
	"news/features/categories"

	"gorm.io/gorm"
)

type PostType struct {
	gorm.Model
	Category string
	Type     string
}

func (post *PostType) toCore() categories.Core {
	return categories.Core{
		ID:       int(post.ID),
		Category: post.Category,
		Type:     post.Type,
	}
}

func ToCoreList(data []PostType) []categories.Core {
	result := []categories.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromPostType(post *categories.Core) PostType {
	return PostType{
		Category: post.Category,
		Type:     post.Type,
	}
}

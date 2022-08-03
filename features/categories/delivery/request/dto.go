package request

import (
	"news/features/categories"
)

type PostType struct {
	Category string `json:"category" form:"category"`
	Type     string `json:"type" form:"type"`
}

func ToCore(dataReq PostType) categories.Core {
	category := categories.Core{
		Category: dataReq.Category,
		Type:     dataReq.Type,
	}
	return category
}

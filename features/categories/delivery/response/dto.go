package response

import "news/features/categories"

type PostType struct {
	ID       int    `json:"post_type_id"`
	Category string `json:"category"`
	Type     string `json:"type"`
}

func FromCore(dataReq categories.Core) PostType {
	return PostType{
		ID:       dataReq.ID,
		Category: dataReq.Category,
		Type:     dataReq.Type,
	}
}

func FromCoreToList(dataReq []categories.Core) []PostType {
	result := []PostType{}
	for key := range dataReq {
		result = append(result, FromCore(dataReq[key]))
	}
	return result
}

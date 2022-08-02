package request

import "news/features/users"

type User struct {
	FullName string `json:"full_name" form:"full_name"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(dataReq User) users.Core {
	core := users.Core{
		FullName: dataReq.FullName,
		Phone:    dataReq.Phone,
		Email:    dataReq.Email,
		Password: dataReq.Password,
	}
	return core
}

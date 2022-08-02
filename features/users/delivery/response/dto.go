package response

type User struct {
	ID       int    `json:"user_id" form:"user_id"`
	FullName string `json:"full_name" form:"full_name"`
	Token    string `json:"token" form:"token"`
}

func ToResponse(token, fullName string, userID int) User {
	return User{
		ID:       userID,
		FullName: fullName,
		Token:    token,
	}
}

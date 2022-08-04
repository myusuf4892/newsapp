package business

import (
	"errors"
	"news/features/users"
	_mockUser "news/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddUser(t *testing.T) {
	repo := new(_mockUser.UserData)
	dataReq := users.Core{
		FullName: "Muhamad Yusup",
		Phone:    "08570001000",
		Email:    "myusup@gmail.com",
		Password: "123456",
	}

	t.Run("Test Register User Success", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(nil).Once()
		srv := NewUserBusiness(repo)

		err := srv.AddUser(dataReq)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Test Register User Failed", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(errors.New("failed to register")).Once()
		srv := NewUserBusiness(repo)

		err := srv.AddUser(dataReq)
		assert.NotNil(t, err)
	})
}

// func TestLogin(t *testing.T) {
// 	repo := new(_mockUser.UserData)
// 	dataReq := users.Core{
// 		Email:    "myusup@gmail.com",
// 		Password: "123456",
// 	}

// 	returnData := users.Core{
// 		ID:       1,
// 		FullName: "Muhamad Yusup",
// 		Password: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTk1OTk2NzIsInVzZXJJRCI6Mn0.pXhJLBW0suQt_u4HHjr8lVQDMkUBxUH2SLCEXKa8PxA",
// 	}

// 	t.Run("Login User", func(t *testing.T) {
// 		repo.On("FindUser", mock.Anything).Return(returnData, nil).Once()
// 		srv := NewUserBusiness(repo)

// 		token, name, userID, err := srv.Auth(dataReq)

// 		assert.Nil(t, err)
// 		assert.Equal(t, returnData.ID, userID)
// 		assert.Equal(t, returnData.FullName, name)
// 		assert.Equal(t, returnData.Password, token)
// 		repo.AssertExpectations(t)
// 	})
// }

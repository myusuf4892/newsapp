package business

import (
	"errors"
	"news/features/posts"
	_mockPost "news/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddPost(t *testing.T) {
	repo := new(_mockPost.PostData)
	dataReq := posts.Core{
		ID:          1,
		Tittle:      "Kita Bisa",
		Description: "Indonesia Bisa Bangkit",
		PostTypeID:  1,
		UserID:      1,
	}

	t.Run("Test Posting Success", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(nil).Once()
		srv := NewPostBusiness(repo)

		err := srv.AddPost(dataReq)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Test Posting Failed", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(errors.New("failed to insert Post")).Once()
		srv := NewPostBusiness(repo)

		err := srv.AddPost(dataReq)
		assert.NotNil(t, err)
	})
}

func TestGetPost(t *testing.T) {
	repo := new(_mockPost.PostData)
	returnData := []posts.Core{
		{
			ID:          1,
			Tittle:      "Kita Bisa",
			Description: "Indonesia Bisa Updated",
			PostTypeID:  1,
			UserID:      1,
			PostType: posts.PostType{
				Category: "Opini",
				Type:     "Artikel",
			},
			User: posts.User{
				FullName: "Muhamad Yusup",
			},
		},
	}

	t.Run("Test Get Posting Success", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(returnData, nil).Once()
		srv := NewPostBusiness(repo)

		res, err := srv.GetPost()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Test Get Posting Failed", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return([]posts.Core{}, errors.New("error server")).Once()
		srv := NewPostBusiness(repo)

		_, err := srv.GetPost()
		assert.NotNil(t, err)
	})
}

func TestUpdatePost(t *testing.T) {
	repo := new(_mockPost.PostData)
	dataReq := posts.Core{
		ID:          1,
		Tittle:      "Kita Bisa",
		Description: "Indonesia Bisa Updated",
		PostTypeID:  1,
		UserID:      1,
		PostType: posts.PostType{
			Category: "Opini",
			Type:     "Artikel",
		},
		User: posts.User{
			FullName: "Muhamad Yusup",
		},
	}
	ID := 1

	t.Run("Test Update Post Success", func(t *testing.T) {
		repo.On("Update", dataReq, ID).Return(nil).Once()
		srv := NewPostBusiness(repo)

		err := srv.UpdatePost(dataReq, ID)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Test Update Post Failed", func(t *testing.T) {
		repo.On("Update", dataReq, ID).Return(errors.New("failed to insert post")).Once()
		srv := NewPostBusiness(repo)

		err := srv.UpdatePost(dataReq, ID)
		assert.NotNil(t, err)
	})
}

func TestDestroyPost(t *testing.T) {
	repo := new(_mockPost.PostData)
	ID := 1

	t.Run("Test Delete Post Success", func(t *testing.T) {
		repo.On("Destroy", mock.Anything).Return(nil).Once()
		srv := NewPostBusiness(repo)

		err := srv.DestroyPost(ID)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Test Delete Post Failed", func(t *testing.T) {
		repo.On("Destroy", mock.Anything).Return(errors.New("failed delete post")).Once()
		srv := NewPostBusiness(repo)

		err := srv.DestroyPost(ID)
		assert.NotNil(t, err)
	})
}

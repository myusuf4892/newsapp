package business

import (
	"errors"
	"news/features/categories"
	_mockPost "news/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddCategories(t *testing.T) {
	repo := new(_mockPost.CategoryData)
	dataReq := categories.Core{
		ID:       1,
		Category: "Opini",
		Type:     "Artikel",
	}
	dataReqFailed := categories.Core{
		ID:       1,
		Category: "",
		Type:     "Artikel",
	}
	noDataReq := categories.Core{
		ID:       1,
		Category: "",
		Type:     "Artikel",
	}

	t.Run("Test Add Post Type Success", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(1, nil).Once()
		srv := NewCategoryBusiness(repo)

		row, err := srv.AddCategories(dataReq)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
		repo.AssertExpectations(t)
	})

	t.Run("Test Input Post Type Failed", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(0, errors.New("all must be filled")).Once()
		srv := NewCategoryBusiness(repo)

		row, err := srv.AddCategories(dataReqFailed)
		if dataReqFailed.Category == "" || dataReqFailed.Type == "" {
			assert.NotNil(t, err)
			assert.Equal(t, 0, row)
		}
	})

	t.Run("Test No Data Input Post Type", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(0, errors.New("can't input data")).Once()
		srv := NewCategoryBusiness(repo)

		row, err := srv.AddCategories(noDataReq)
		if row == 0 {
			assert.Error(t, err)
			assert.Equal(t, 0, row)
		}
	})

	t.Run("Test Add Post Type Failed", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(0, errors.New("failed to insert Post Type")).Once()
		srv := NewCategoryBusiness(repo)

		row, err := srv.AddCategories(dataReq)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
}

func TestGetCategories(t *testing.T) {
	repo := new(_mockPost.CategoryData)
	returnData := []categories.Core{
		{
			ID:       1,
			Category: "Opini",
			Type:     "Artikel",
		},
	}

	t.Run("Test Get Post Type", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(returnData, nil).Once()
		srv := NewCategoryBusiness(repo)

		res, err := srv.GetCategories()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Test Get Post Type Failed", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return([]categories.Core{}, errors.New("error server")).Once()
		srv := NewCategoryBusiness(repo)

		_, err := srv.GetCategories()
		assert.NotNil(t, err)
	})
}

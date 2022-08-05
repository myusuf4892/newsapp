package mocks

import (
	"news/features/categories"

	"github.com/stretchr/testify/mock"
)

type CategoryUseCase struct {
	mock.Mock
}

func (_m *CategoryUseCase) AddPost(dataPost categories.Core) (row int, err error) {
	ret := _m.Called(dataPost)

	var r0 int
	if rf, ok := ret.Get(0).(func(categories.Core) int); ok {
		r0 = rf(dataPost)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(categories.Core) error); ok {
		r1 = rf(dataPost)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *CategoryUseCase) GetPost() (dataPost []categories.Core, err error) {
	ret := _m.Called()

	var r0 []categories.Core
	if rf, ok := ret.Get(0).(func() []categories.Core); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).([]categories.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *CategoryUseCase) UpdatePost(dataPost categories.Core, ID int) (row int, err error) {
	ret := _m.Called(dataPost, ID)

	var r0 int
	if rf, ok := ret.Get(0).(func(categories.Core, int) int); ok {
		r0 = rf(dataPost, ID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(categories.Core, int) error); ok {
		r1 = rf(dataPost, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *CategoryUseCase) DestroyPost(ID int) (row int, err error) {
	ret := _m.Called(ID)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

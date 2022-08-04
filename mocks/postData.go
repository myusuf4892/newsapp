package mocks

import (
	"news/features/posts"

	"github.com/stretchr/testify/mock"
)

type PostData struct {
	mock.Mock
}

func (_m *PostData) Insert(dataPost posts.Core) (err error) {
	ret := _m.Called(dataPost)

	var r0 error
	if rf, ok := ret.Get(0).(func(posts.Core) error); ok {
		r0 = rf(dataPost)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *PostData) Get() (dataPost []posts.Core, err error) {
	ret := _m.Called()

	var r0 []posts.Core
	if rf, ok := ret.Get(0).(func() []posts.Core); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).([]posts.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *PostData) Update(dataPost posts.Core, ID int) (err error) {
	ret := _m.Called(dataPost, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(posts.Core, int) error); ok {
		r0 = rf(dataPost, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *PostData) Destroy(ID int) (err error) {
	ret := _m.Called(ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

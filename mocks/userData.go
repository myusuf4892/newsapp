package mocks

import (
	"news/features/users"

	"github.com/stretchr/testify/mock"
)

type UserData struct {
	mock.Mock
}

func (_m *UserData) Insert(dataUser users.Core) (err error) {
	ret := _m.Called(dataUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(dataUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *UserData) FindUser(email string) (response users.Core, err error) {
	ret := _m.Called(email)

	var r0 users.Core
	if rf, ok := ret.Get(0).(func(string) users.Core); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(users.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(0)
	}

	return r0, r1
}

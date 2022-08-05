package mocks

import (
	"news/features/users"

	"github.com/stretchr/testify/mock"
)

type UserUseCase struct {
	mock.Mock
}

func (_m *UserUseCase) AddUser(dataUser users.Core) (err error) {
	ret := _m.Called(dataUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Core) error); ok {
		r0 = rf(dataUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *UserUseCase) Auth(dataUser users.Core) (token, fullName string, userID int, err error) {
	ret := _m.Called(dataUser)

	var r0 string
	if rf, ok := ret.Get(0).(func(users.Core) string); ok {
		r0 = rf(dataUser)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(users.Core) string); ok {
		r1 = rf(dataUser)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(users.Core) int); ok {
		r2 = rf(dataUser)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(users.Core) error); ok {
		r3 = rf(dataUser)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

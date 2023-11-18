// Code generated by mockery v2.36.0. DO NOT EDIT.

package repository

import (
	models "CardozoCasariegoLuciano/StudyNotes/Models"

	mock "github.com/stretchr/testify/mock"
)

// MockIStorage is an autogenerated mock type for the IStorage type
type MockIStorage struct {
	mock.Mock
}

type MockIStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIStorage) EXPECT() *MockIStorage_Expecter {
	return &MockIStorage_Expecter{mock: &_m.Mock}
}

// FindUserByEmail provides a mock function with given fields: email
func (_m *MockIStorage) FindUserByEmail(email string) models.User {
	ret := _m.Called(email)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(string) models.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	return r0
}

// MockIStorage_FindUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindUserByEmail'
type MockIStorage_FindUserByEmail_Call struct {
	*mock.Call
}

// FindUserByEmail is a helper method to define mock.On call
//   - email string
func (_e *MockIStorage_Expecter) FindUserByEmail(email interface{}) *MockIStorage_FindUserByEmail_Call {
	return &MockIStorage_FindUserByEmail_Call{Call: _e.mock.On("FindUserByEmail", email)}
}

func (_c *MockIStorage_FindUserByEmail_Call) Run(run func(email string)) *MockIStorage_FindUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockIStorage_FindUserByEmail_Call) Return(_a0 models.User) *MockIStorage_FindUserByEmail_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIStorage_FindUserByEmail_Call) RunAndReturn(run func(string) models.User) *MockIStorage_FindUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// ListAllUsers provides a mock function with given fields: _a0
func (_m *MockIStorage) ListAllUsers(_a0 *[]models.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*[]models.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIStorage_ListAllUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAllUsers'
type MockIStorage_ListAllUsers_Call struct {
	*mock.Call
}

// ListAllUsers is a helper method to define mock.On call
//   - _a0 *[]models.User
func (_e *MockIStorage_Expecter) ListAllUsers(_a0 interface{}) *MockIStorage_ListAllUsers_Call {
	return &MockIStorage_ListAllUsers_Call{Call: _e.mock.On("ListAllUsers", _a0)}
}

func (_c *MockIStorage_ListAllUsers_Call) Run(run func(_a0 *[]models.User)) *MockIStorage_ListAllUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*[]models.User))
	})
	return _c
}

func (_c *MockIStorage_ListAllUsers_Call) Return(_a0 error) *MockIStorage_ListAllUsers_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIStorage_ListAllUsers_Call) RunAndReturn(run func(*[]models.User) error) *MockIStorage_ListAllUsers_Call {
	_c.Call.Return(run)
	return _c
}

// SaveUser provides a mock function with given fields: user
func (_m *MockIStorage) SaveUser(user *models.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIStorage_SaveUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveUser'
type MockIStorage_SaveUser_Call struct {
	*mock.Call
}

// SaveUser is a helper method to define mock.On call
//   - user *models.User
func (_e *MockIStorage_Expecter) SaveUser(user interface{}) *MockIStorage_SaveUser_Call {
	return &MockIStorage_SaveUser_Call{Call: _e.mock.On("SaveUser", user)}
}

func (_c *MockIStorage_SaveUser_Call) Run(run func(user *models.User)) *MockIStorage_SaveUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.User))
	})
	return _c
}

func (_c *MockIStorage_SaveUser_Call) Return(_a0 error) *MockIStorage_SaveUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIStorage_SaveUser_Call) RunAndReturn(run func(*models.User) error) *MockIStorage_SaveUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIStorage creates a new instance of MockIStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIStorage {
	mock := &MockIStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

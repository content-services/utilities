// Code generated by mockery v2.36.1. DO NOT EDIT.

package yum

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockYumRepository is an autogenerated mock type for the YumRepository type
type MockYumRepository struct {
	mock.Mock
}

// Clear provides a mock function with given fields:
func (_m *MockYumRepository) Clear() {
	_m.Called()
}

// Comps provides a mock function with given fields: ctx
func (_m *MockYumRepository) Comps(ctx context.Context) (*Comps, int, error) {
	ret := _m.Called(ctx)

	var r0 *Comps
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) (*Comps, int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *Comps); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Comps)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) int); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Configure provides a mock function with given fields: settings
func (_m *MockYumRepository) Configure(settings YummySettings) {
	_m.Called(settings)
}

// Environments provides a mock function with given fields: ctx
func (_m *MockYumRepository) Environments(ctx context.Context) ([]Environment, int, error) {
	ret := _m.Called(ctx)

	var r0 []Environment
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]Environment, int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []Environment); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) int); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PackageGroups provides a mock function with given fields: ctx
func (_m *MockYumRepository) PackageGroups(ctx context.Context) ([]PackageGroup, int, error) {
	ret := _m.Called(ctx)

	var r0 []PackageGroup
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]PackageGroup, int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []PackageGroup); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]PackageGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) int); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Packages provides a mock function with given fields: ctx
func (_m *MockYumRepository) Packages(ctx context.Context) ([]Package, int, error) {
	ret := _m.Called(ctx)

	var r0 []Package
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]Package, int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []Package); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Package)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) int); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Repomd provides a mock function with given fields: ctx
func (_m *MockYumRepository) Repomd(ctx context.Context) (*Repomd, int, error) {
	ret := _m.Called(ctx)

	var r0 *Repomd
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) (*Repomd, int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *Repomd); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Repomd)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) int); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Signature provides a mock function with given fields: ctx
func (_m *MockYumRepository) Signature(ctx context.Context) (*string, int, error) {
	ret := _m.Called(ctx)

	var r0 *string
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) (*string, int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) int); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewMockYumRepository creates a new instance of MockYumRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockYumRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockYumRepository {
	mock := &MockYumRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

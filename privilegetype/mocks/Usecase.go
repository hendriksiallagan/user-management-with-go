// Code generated by mockery v1.0.0
package mocks

import (
	"context"
)
import "github.com/stretchr/testify/mock"
import "github.com/user-management-with-go/models"

type Usecase struct {
	mock.Mock
}


func (_m *Usecase) Fetch(ctx context.Context, page int64) ([]*models.Role, error) {
	ret := _m.Called(ctx, page)

	var r0 []*models.Role
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*models.Role); ok {
		r0 = rf(ctx, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Role)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, int64) error); ok {
		r2 = rf(ctx, page)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r2
}

func (_m *Usecase) FetchByID(ctx context.Context, id int64) (*models.Role, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Role
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Role); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Usecase) Store(_a0 context.Context, _a1 *models.Role) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Role) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Usecase) Update(ctx context.Context, ar *models.Role, id int64) error {
	ret := _m.Called(ctx, ar)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Role) error); ok {
		r0 = rf(ctx, ar)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

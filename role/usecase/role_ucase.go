package usecase

import (
	"context"
	"time"
	"github.com/user-management-with-go/role"
	"github.com/user-management-with-go/models"
)

type roleUsecase struct {
	roleRepo    role.Repository
	contextTimeout time.Duration
}

func NewRoleUsecase(a role.Repository, timeout time.Duration) role.Usecase {
	return &roleUsecase{
		roleRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *roleUsecase) Fetch(c context.Context, page int64, limit int64, search string, status string) (res []*models.Role, total int, err error) {
	if page <= 0 {
		page = 1		
	}

	page = (page - 1) * limit

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listRole, total, err := a.roleRepo.Fetch(ctx, page, limit, search, status)

	if err != nil {
		return nil, 0, err
	}

	return listRole, total, nil
}

func (a *roleUsecase) Store(c context.Context, m *models.Role) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.roleRepo.Store(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (a *roleUsecase) Update(c context.Context, m *models.Role, id int64) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.roleRepo.Update(ctx, m, id)
	if err != nil {
		return err
	}
	return nil
}
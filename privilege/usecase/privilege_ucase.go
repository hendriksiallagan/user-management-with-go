package usecase

import (
	"context"
	"time"
	"github.com/user-management-with-go/privilege"
	"github.com/user-management-with-go/models"
	"fmt"
)

type privilegeUsecase struct {
	privilegeRepo    privilege.Repository
	contextTimeout 		time.Duration
}

func NewPrivilegeUsecase(a privilege.Repository, timeout time.Duration) privilege.Usecase {
	return &privilegeUsecase{
		privilegeRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *privilegeUsecase) Fetch(c context.Context, page int64, limit int64, search string) (res []*models.Privilege, total int, err error) {
	if page <= 0 {
		page = 1		
	}

	page = (page - 1) * limit

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listPage, total, err := a.privilegeRepo.Fetch(ctx, page, limit, search)

	if err != nil {
		return nil, 0, err
	}

	return listPage, total, nil
}

func (a *privilegeUsecase) Store(c context.Context, m *models.Privilege, code string) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.privilegeRepo.Store(ctx, m, code)
	if err != nil {
		return err
	}
	return nil
}

func (a *privilegeUsecase) Update(c context.Context, m *models.Privilege, id int64) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.privilegeRepo.Update(ctx, m, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *privilegeUsecase) Generate(c context.Context) (code string, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	total, err := a.privilegeRepo.Count(ctx)
	if err != nil {
		return "failed", err
	}

	var newIncre = (total + 1)

 	digit := fmt.Sprintf("%04d\n", newIncre)

 	prefix := "PRV"

 	code = prefix+digit

	return code, nil
}
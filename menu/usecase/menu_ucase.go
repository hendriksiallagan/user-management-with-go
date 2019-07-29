package usecase

import (
	"context"
	"time"
	"github.com/user-management-with-go/menu"
	"github.com/user-management-with-go/models"
	"fmt"
)

type menuUsecase struct {
	menuRepo    menu.Repository
	contextTimeout 		time.Duration
}

func NewMenuUsecase(a menu.Repository, timeout time.Duration) menu.Usecase {
	return &menuUsecase{
		menuRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *menuUsecase) Fetch(c context.Context, page int64, limit int64, search string) (res []*models.Menu, total int, err error) {
	if page <= 0 {
		page = 1		
	}

	page = (page - 1) * limit

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listPage, total, err := a.menuRepo.Fetch(ctx, page, limit, search)

	if err != nil {
		return nil, 0, err
	}

	return listPage, total, nil
}

func (a *menuUsecase) Store(c context.Context, m *models.Menu, code string) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.menuRepo.Store(ctx, m, code)
	if err != nil {
		return err
	}
	return nil
}

func (a *menuUsecase) Update(c context.Context, m *models.Menu, id int64) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.menuRepo.Update(ctx, m, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *menuUsecase) Generate(c context.Context) (code string, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	total, err := a.menuRepo.Count(ctx)
	if err != nil {
		return "failed", err
	}

	var newIncre = (total + 1)

 	digit := fmt.Sprintf("%04d\n", newIncre)

 	prefix := "M"

 	code = prefix+digit

	return code, nil
}
package usecase

import (
	"context"
	"time"
	"github.com/user-management-with-go/page"
	"github.com/user-management-with-go/models"
	"fmt"
)

type pageUsecase struct {
	pageRepo    page.Repository
	contextTimeout time.Duration
}

func NewPageUsecase(a page.Repository, timeout time.Duration) page.Usecase {
	return &pageUsecase{
		pageRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *pageUsecase) Fetch(c context.Context, page int64, limit int64, search string) (res []*models.Page, total int, err error) {
	if page <= 0 {
		page = 1		
	}

	page = (page - 1) * limit

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listPage, total, err := a.pageRepo.Fetch(ctx, page, limit, search)

	if err != nil {
		return nil, 0, err
	}

	return listPage, total, nil
}

func (a *pageUsecase) Store(c context.Context, m *models.Page, code string) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.pageRepo.Store(ctx, m, code)
	if err != nil {
		return err
	}
	return nil
}

func (a *pageUsecase) Update(c context.Context, m *models.Page, id int64) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.pageRepo.Update(ctx, m, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *pageUsecase) Generate(c context.Context) (code string, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	total, err := a.pageRepo.Count(ctx)
	if err != nil {
		return "failed", err
	}

	var newIncre = (total + 1)

 	digit := fmt.Sprintf("%04d\n", newIncre)

 	prefix := "P"

 	code = prefix+digit

	return code, nil
}
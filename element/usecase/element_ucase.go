package usecase

import (
	"context"
	"time"
	"github.com/user-management-with-go/element"
	"github.com/user-management-with-go/models"
	"fmt"
)

type elementUsecase struct {
	elementRepo    		element.Repository
	contextTimeout 		time.Duration
}

func NewElementUsecase(a element.Repository, timeout time.Duration) element.Usecase {
	return &elementUsecase{
		elementRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *elementUsecase) Fetch(c context.Context, page int64, limit int64, search string) (res []*models.Element, total int, err error) {
	if page <= 0 {
		page = 1		
	}

	page = (page - 1) * limit

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listPage, total, err := a.elementRepo.Fetch(ctx, page, limit, search)

	if err != nil {
		return nil, 0, err
	}

	return listPage, total, nil
}

func (a *elementUsecase) Store(c context.Context, m *models.Element, code string) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.elementRepo.Store(ctx, m, code)

	if err != nil {
		return err
	}
	return nil
}

func (a *elementUsecase) Update(c context.Context, m *models.Element, id int64) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.elementRepo.Update(ctx, m, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *elementUsecase) Generate(c context.Context) (code string, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	total, err := a.elementRepo.Count(ctx)
	if err != nil {
		return "failed", err
	}

	var newIncre = (total + 1)

 	digit := fmt.Sprintf("%04d\n", newIncre)

 	prefix := "E"

 	code = prefix+digit

	return code, nil
}
package usecase

import (
	"context"
	"time"
	"github.com/user-management-with-go/actionelement"
	"github.com/user-management-with-go/models"
	"fmt"
)

type actionelementUsecase struct {
	actionelementRepo    actionelement.Repository
	contextTimeout 		time.Duration
}

func NewActionelementUsecase(a actionelement.Repository, timeout time.Duration) actionelement.Usecase {
	return &actionelementUsecase{
		actionelementRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *actionelementUsecase) Fetch(c context.Context, page int64, limit int64, search string) (res []*models.Actionelement, total int, err error) {
	if page <= 0 {
		page = 1		
	}

	page = (page - 1) * limit

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listPage, total, err := a.actionelementRepo.Fetch(ctx, page, limit, search)

	if err != nil {
		return nil, 0, err
	}

	return listPage, total, nil
}

func (a *actionelementUsecase) Store(c context.Context, m *models.Actionelement, code string) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.actionelementRepo.Store(ctx, m, code)
	if err != nil {
		return err
	}
	return nil
}

func (a *actionelementUsecase) Update(c context.Context, m *models.Actionelement, id int64) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.actionelementRepo.Update(ctx, m, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *actionelementUsecase) Generate(c context.Context) (code string, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	total, err := a.actionelementRepo.Count(ctx)
	if err != nil {
		return "failed", err
	}

	var newIncre = (total + 1)

 	digit := fmt.Sprintf("%04d\n", newIncre)

 	prefix := "AE"

 	code = prefix+digit

	return code, nil
}
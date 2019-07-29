package usecase

import (
	"context"
	"time"
	"github.com/user-management-with-go/privilegetype"
	"github.com/user-management-with-go/models"
	//"golang.org/x/sync/errgroup"
	//"fmt"
)

type privilegetypeUsecase struct {
	privilegetypeRepo    privilegetype.Repository
	contextTimeout time.Duration
}

func NewPrivilegeTypeUsecase(a privilegetype.Repository, timeout time.Duration) privilegetype.Usecase {
	return &privilegetypeUsecase{
		privilegetypeRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *privilegetypeUsecase) Fetch(c context.Context) ([]*models.PrivilegeType, error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listPT, err := a.privilegetypeRepo.Fetch(ctx)

	if err != nil {
		return nil, err
	}

	return listPT, nil
}
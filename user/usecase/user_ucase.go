package usecase

import (
	"context"
	"time"
	"github.com/user-management-with-go/user"
	"github.com/user-management-with-go/models"
	//"golang.org/x/sync/errgroup"
	"fmt"
)

type userUsecase struct {
	userRepo    user.Repository
	contextTimeout time.Duration
}

func NewUserUsecase(a user.Repository, timeout time.Duration) user.Usecase {
	return &userUsecase{
		userRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *userUsecase) Fetch(c context.Context, page uint64, limit uint64, search string) (res []*models.User, total int, err error) {
	if page <= 0 {
		return nil, 0, err
	}

	page = (page - 1) * limit

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listUser, total, err := a.userRepo.Fetch(ctx, page, limit, search)

	if err != nil {
		return nil, 0, err
	}

	return listUser, total,nil
}

func (a *userUsecase) FetchPasswordByID(c context.Context, page uint64, limit uint64, id uint64) (res []*models.UserPassword, total int, err error) {
	if page <= 0 {
		return nil, 0, err
	}

	page = (page - 1) * limit

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listPass, total, err := a.userRepo.FetchPasswordByID(ctx, uint64(page), uint64(limit), uint64(id))

	if err != nil {
		return nil, 0, err
	}

	return listPass, total, nil
}

func (a *userUsecase) FetchPinByID(c context.Context, page uint64, limit uint64, id uint64) (res []*models.UserPIN, total int, err error) {
	if page <= 0 {
		return nil, 0, err
	}

	page = (page - 1) * limit

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listPass, total, err := a.userRepo.FetchPinByID(ctx, uint64(page), uint64(limit), uint64(id))

	if err != nil {
		return nil, 0, err
	}

	return listPass, total, nil
}

func (a *userUsecase) FetchRoleByID(c context.Context, page uint64, limit uint64, id uint64) (res []*models.UserRole, total int, err error) {
	if page <= 0 {
		return nil, 0, err
	}

	page = (page - 1) * limit

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listPass, total, err := a.userRepo.FetchRoleByID(ctx, uint64(page), uint64(limit), uint64(id))

	if err != nil {
		return nil, 0, err
	}

	return listPass, total, nil
}

func (a *userUsecase) FetchStatusInfoByID(c context.Context, page uint64, limit uint64, id uint64) (res []*models.UserStatusInfo, total int, err error) {
	if page <= 0 {
		return nil, 0, err
	}

	page = (page - 1) * limit

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listPass, total, err := a.userRepo.FetchStatusInfoByID(ctx, uint64(page), uint64(limit), uint64(id))

	if err != nil {
		return nil, 0, err
	}

	return listPass, total, nil
}

func (a *userUsecase) FetchOtpByID(c context.Context, page uint64, limit uint64, id uint64) (res []*models.UserOtp, total int, err error) {
	if page <= 0 {
		return nil, 0, err
	}

	page = (page - 1) * limit

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listPass, total, err := a.userRepo.FetchOtpByID(ctx, uint64(page), uint64(limit), uint64(id))

	if err != nil {
		return nil, 0, err
	}

	return listPass, total, nil
}

func (a *userUsecase) FetchTokenByID(c context.Context, page uint64, limit uint64, id uint64) (res []*models.UserToken, total int, err error) {
	if page <= 0 {
		return nil, 0, err
	}

	page = (page - 1) * limit

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	listPass, total, err := a.userRepo.FetchTokenByID(ctx, uint64(page), uint64(limit), uint64(id))

	if err != nil {
		return nil, 0, err
	}

	return listPass, total, nil
}

func (a *userUsecase) Store(c context.Context, m *models.User, code string) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.userRepo.Store(ctx, m, code)
	if err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) StoreRole(c context.Context, m *models.UserRole, code string) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.userRepo.StoreRole(ctx, m, code)
	if err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) FetchByID(c context.Context, id int64) (m *models.User, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err := a.userRepo.FetchByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *userUsecase) Update(c context.Context, m *models.User, id int64) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.userRepo.Update(ctx, m, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) UpdateRole(c context.Context, m *models.UserRole, id uint64) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.userRepo.UpdateRole(ctx, m, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) UpdatePassword(c context.Context, m *models.ResetPassword, id uint64) error {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.userRepo.UpdatePassword(ctx, m, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *userUsecase) Generate(c context.Context) (code string, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	total, err := a.userRepo.Count(ctx)
	if err != nil {
		return "failed", err
	}

	var newIncre = (total + 1)

 	digit := fmt.Sprintf("%06d\n", newIncre)

 	dateTime := time.Now()
 	dateCode := dateTime.Format("060201")

 	code = dateCode+digit

	return code, nil
}

func (a *userUsecase) GenerateRole(c context.Context) (code string, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	total, err := a.userRepo.CountRole(ctx)
	if err != nil {
		return "failed", err
	}

	var newIncre = (total + 1)

	prefix := "UR"

	digit := fmt.Sprintf("%04d\n", newIncre)

	dateTime := time.Now()
	dateCode := dateTime.Format("2006")

	code = prefix+dateCode+digit

	return code, nil
}
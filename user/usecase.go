package user

import (
	"context"
	"github.com/user-management-with-go/models"
)

type Usecase interface {
	Fetch(ctx context.Context, page uint64, limit uint64, search string) (res []*models.User, total int, err error)
	FetchPasswordByID(ctx context.Context, page uint64, limit uint64, id uint64) (res []*models.UserPassword, total int, err error)
	FetchPinByID(ctx context.Context, page uint64, limit uint64, id uint64) (res []*models.UserPIN, total int, err error)
	FetchRoleByID(ctx context.Context, page uint64, limit uint64, id uint64) (res []*models.UserRole, total int, err error)
	FetchStatusInfoByID(ctx context.Context, page uint64, limit uint64, id uint64) (res []*models.UserStatusInfo, total int, err error)
	FetchOtpByID(ctx context.Context, page uint64, limit uint64, id uint64) (res []*models.UserOtp, total int, err error)
	FetchTokenByID(ctx context.Context, page uint64, limit uint64, id uint64) (res []*models.UserToken, total int, err error)
	Store(ctx context.Context, a *models.User, code string) error
	StoreRole(ctx context.Context, a *models.UserRole, code string) error
	FetchByID(ctx context.Context, id int64) (*models.User, error)
	Update(ctx context.Context, a *models.User, id int64) error
	UpdateRole(ctx context.Context, a *models.UserRole, id uint64) error
	UpdatePassword(ctx context.Context, a *models.ResetPassword, id uint64) error
	Generate(ctx context.Context) (code string, err error)
	GenerateRole(ctx context.Context) (code string, err error)
}

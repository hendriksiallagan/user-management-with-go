package actionelement

import (
	"context"
	"github.com/user-management-with-go/models"
)

type Repository interface {
	Fetch(ctx context.Context, offset int64, limit int64, search string) (res []*models.Actionelement, total int, err error)
	Store(ctx context.Context, a *models.Actionelement, code string) error
	Update(ctx context.Context, a *models.Actionelement, id int64) error
	Count(ctx context.Context) (count int, err error)
}

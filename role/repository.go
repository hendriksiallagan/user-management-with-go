package role

import (
	"context"
	"github.com/user-management-with-go/models"
)

type Repository interface {
	Fetch(ctx context.Context, offset int64, limit int64, search string, status string) (res []*models.Role, total int, err error)
	Store(ctx context.Context, a *models.Role) error
	Update(ctx context.Context, a *models.Role, id int64) error
	Count(ctx context.Context) (count int, err error)
}

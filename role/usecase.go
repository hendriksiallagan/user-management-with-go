package role

import (
	"context"
	"github.com/user-management-with-go/models"
)

type Usecase interface {
	Fetch(ctx context.Context, page int64, limit int64, search string, status string) (a []*models.Role, total int, err error)
	Store(ctx context.Context, a *models.Role) error
	Update(ctx context.Context, a *models.Role, id int64) error
}

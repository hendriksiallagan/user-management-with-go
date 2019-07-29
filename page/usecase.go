package page

import (
	"context"
	"github.com/user-management-with-go/models"
)

type Usecase interface {
	Fetch(ctx context.Context, page int64, limit int64, search string) (res []*models.Page, total int, err error)
	Store(ctx context.Context, a *models.Page, code string) error
	Update(ctx context.Context, a *models.Page, id int64) error
	Generate(ctx context.Context) (code string, err error)
}

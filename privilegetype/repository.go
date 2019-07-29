package privilegetype

import (
	"context"
	"github.com/user-management-with-go/models"
)

type Repository interface {
	Fetch(ctx context.Context) (res []*models.PrivilegeType, err error)
}

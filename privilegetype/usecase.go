package privilegetype

import (
	"context"
	"github.com/user-management-with-go/models"
)

type Usecase interface {
	Fetch(ctx context.Context) ([]*models.PrivilegeType, error)
}

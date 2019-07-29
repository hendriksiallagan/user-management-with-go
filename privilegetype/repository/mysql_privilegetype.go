package repository

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/user-management-with-go/privilegetype"
	"github.com/user-management-with-go/models"
	//"fmt"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" 
)

type mysqlPrivilegeTypeRepository struct {
	Conn *sql.DB
}

func NewMysqlPrivilegeTypeRepository(Conn *sql.DB) privilegetype.Repository {

	return &mysqlPrivilegeTypeRepository{Conn}
}

func (m *mysqlPrivilegeTypeRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.PrivilegeType, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.PrivilegeType, 0)
	for rows.Next() {
		s := new(models.PrivilegeType)
		err = rows.Scan(
			&s.MptID,
			&s.MptName,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

func (m *mysqlPrivilegeTypeRepository) Fetch(ctx context.Context) ([]*models.PrivilegeType, error) {
	query := `select mpt_id, mpt_name from mr_privilege_types `

	res, err := m.fetch(ctx, query)

	if err != nil {
		return nil, err
	}

	return res, err

}


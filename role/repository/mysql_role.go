package repository

import (
	"context"
	"database/sql"
	"github.com/user-management-with-go/models"
	"github.com/user-management-with-go/role"
	"github.com/sirupsen/logrus"
	"reflect"

	//"fmt"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" 
)

type mysqlRoleRepository struct {
	Conn *sql.DB
}

func NewMysqlRoleRepository(Conn *sql.DB) role.Repository {

	return &mysqlRoleRepository{Conn}
}

func (m *mysqlRoleRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Role, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Role, 0)
	for rows.Next() {
		s := new(models.Role)
		err = rows.Scan(
			&s.MroID,
			&s.MroCode,
			&s.MroName,
			&s.MroDescription,
			&s.MroStatus,
			&s.MroCreatedBy,
			&s.MroUpdatedBy,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

func (m *mysqlRoleRepository) count(ctx context.Context, query string, args ...interface{}) (count int, err error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {   
	    if err := rows.Scan(&count); err != nil {
	        logrus.Error(err)
	        return 0, err
	    }
	}

	return count, nil
}

func (m *mysqlRoleRepository) in_array(val interface{}, array interface{}) (exists bool) {
	exists = false
	//index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				//index = i
				exists = true
				return
			}
		}
	}

	return
}

func (m *mysqlRoleRepository) Fetch(ctx context.Context, offset int64, limit int64, search string, status string) (res []*models.Role, total int, err error) {

	query := `select mro_id, mro_code, mro_name, mro_description, mro_status, mro_created_by, mro_updated_by from mr_roles where mro_name like ? `

	if status != "" {
		query += `AND mro_status = ? ORDER BY mro_id desc limit ? offset ?`
		res, err = m.fetch(ctx, query, search, status, limit, offset)
	} else {
		query += `ORDER BY mro_id desc limit ? offset ?`
		res, err = m.fetch(ctx, query, search, limit, offset)
	}

	if err != nil {
		return nil, 0, err
	}

	query2 := `select count(mro_id) from mr_roles where mro_name like ? `

	if status != "" {
		query2 += `AND mro_status = ? `
		total, err = m.count(ctx, query2, search, status)
	} else {
		total, err = m.count(ctx, query2, search)
	}

	if err != nil {
		return nil, 0, err
	}

	return res, total, err

}

func (m *mysqlRoleRepository) Store(ctx context.Context, a *models.Role) error {

	query := `INSERT mr_roles SET mro_code=? , mro_name=? ,  mro_description=? , mro_created_by=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, a.MroCode, a.MroName, a.MroDescription, a.MroCreatedBy)
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	a.MroID = lastId
	return nil
}

func (m *mysqlRoleRepository) Update(ctx context.Context, a *models.Role, id int64) error {

	query := `UPDATE mr_roles SET mro_name=? , mro_description=?,  mro_status=?, mro_updated_by=? WHERE mro_id =?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, a.MroName, a.MroDescription, a.MroStatus, a.MroUpdatedBy, id)

	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlRoleRepository) Count(ctx context.Context) (count int, err error) {
	query := `SELECT count(mro_id) as total from mr_roles`

	total, err := m.count(ctx, query)

	if err != nil {
		return 0, err
	}

	return total, nil
}


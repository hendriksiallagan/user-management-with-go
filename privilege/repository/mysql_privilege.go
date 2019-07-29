package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/user-management-with-go/privilege"
	"github.com/user-management-with-go/models"
	"reflect"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" 
)

type mysqlPrivilegeRepository struct {
	Conn *sql.DB
}

func NewMysqlPrivilegeRepository(Conn *sql.DB) privilege.Repository {

	return &mysqlPrivilegeRepository{Conn}
}

func (m *mysqlPrivilegeRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Privilege, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Privilege, 0)
	for rows.Next() {
		s := new(models.Privilege)
		err = rows.Scan(
			&s.MrprID,
			&s.MrprCode,
			&s.MrprName,
			&s.MrprTypeID,
			&s.MrprRoleID,
			&s.MrprMenuID,
			&s.MrprApiID,
			&s.MrprPageID,
			&s.MrprElementID,
			&s.MrprActionElementID,
			&s.MrprStatus,
			&s.MrprCreatedBy,
			&s.MrprUpdatedBy,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

func (m *mysqlPrivilegeRepository) count(ctx context.Context, query string, args ...interface{}) (count int, err error) {

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

func (m *mysqlPrivilegeRepository) Fetch(ctx context.Context, offset int64, limit int64, search string) (a []*models.Privilege, total int, err error) {
	query := `select mrpr_id, mrpr_code, mrpr_name, mrpr_type_id, mrpr_role_id, mrpr_menu_id, mrpr_api_id, mrpr_page_id, mrpr_element_id, mrpr_action_element_id, mrpr_status, mrpr_created_by, mrpr_updated_by  from mr_privileges where mrpr_name like ? ORDER BY mrpr_id desc limit ? offset ? `

	res, err := m.fetch(ctx, query, search, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	query2 := `select count(mrpr_id) from mr_privileges where mrpr_name like ?  `

	total, err = m.count(ctx, query2, search)

	if err != nil {
		return nil, 0, err
	}

	return res, total, err

}

func (m *mysqlPrivilegeRepository) Store(ctx context.Context, a *models.Privilege, code string) error {
	
	query := `INSERT mr_privileges SET mrpr_code=?, mrpr_name=?, mrpr_type_id=? , mrpr_role_id=? , mrpr_menu_id=?, mrpr_api_id=? , mrpr_page_id=?, mrpr_element_id=?, mrpr_action_element_id=?, mrpr_created_by=? `
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, code, a.MrprName, a.MrprTypeID, a.MrprRoleID, a.MrprMenuID, a.MrprApiID, a.MrprPageID, a.MrprElementID, a.MrprActionElementID, a.MrprCreatedBy)
	fmt.Println(reflect.TypeOf(res))
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	a.MrprID = int(lastId)
	return nil
}

func (m *mysqlPrivilegeRepository) Update(ctx context.Context, a *models.Privilege, id int64) error {

	query := `UPDATE mr_privileges SET mrpr_type_id=? , mrpr_role_id=? , mrpr_menu_id=?, mrpr_api_id=? , mrpr_page_id=?, mrpr_element_id=?, mrpr_action_element_id=?, mrpr_status=?, mrpr_updated_by=? WHERE mrpr_id =?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, a.MrprTypeID, a.MrprRoleID, a.MrprMenuID, a.MrprApiID, a.MrprPageID, a.MrprElementID, a.MrprActionElementID, a.MrprStatus, a.MrprUpdatedBy, id)

	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlPrivilegeRepository) Count(ctx context.Context) (count int, err error) {
	query := `SELECT count(mrpr_id) as total from mr_privileges`

	total, err := m.count(ctx, query)

	if err != nil {
		return 0, err
	}

	return total, nil
}




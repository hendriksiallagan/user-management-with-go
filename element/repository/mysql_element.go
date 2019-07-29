package repository

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/user-management-with-go/element"
	"github.com/user-management-with-go/models"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" 
)

type mysqlElementRepository struct {
	Conn *sql.DB
}

func NewMysqlElementRepository(Conn *sql.DB) element.Repository {

	return &mysqlElementRepository{Conn}
}

func (m *mysqlElementRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Element, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Element, 0)
	for rows.Next() {
		s := new(models.Element)
		err = rows.Scan(
			&s.MeID,
			&s.MeCode,
			&s.MeName,
			&s.MeDescription,
			&s.MePageID,
			&s.MeXpath,
			&s.MeActionElement,
			&s.MeStatus,
			&s.MeCreatedBy,
			&s.MeUpdatedBy,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

func (m *mysqlElementRepository) count(ctx context.Context, query string, args ...interface{}) (count int, err error) {

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

func (m *mysqlElementRepository) Fetch(ctx context.Context, offset int64, limit int64, search string) (a []*models.Element, total int, err error) {
	query := `select me_id, me_code, me_name, me_description, me_page_id, me_xpath, me_action_element, me_status, me_created_by, me_updated_by from mr_elements where me_name like ? ORDER BY me_id desc limit ? offset ? `

	res, err := m.fetch(ctx, query, search, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	query2 := `select count(me_id) from mr_elements where me_name like ?  `

	total, err = m.count(ctx, query2, search)

	if err != nil {
		return nil, 0, err
	}

	return res, total, err

}

func (m *mysqlElementRepository) Store(ctx context.Context, a *models.Element, code string) error {

	query := `INSERT mr_elements SET me_code=?, me_name=?, me_description=? , me_page_id=? , me_xpath=?, me_action_element=? , me_created_by=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, code, a.MeName, a.MeDescription, a.MePageID, a.MeXpath, a.MeActionElement, a.MeCreatedBy)
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	a.MeID = lastId
	return nil
}

func (m *mysqlElementRepository) Update(ctx context.Context, a *models.Element, id int64) error {

	query := `UPDATE mr_elements SET me_name=?, me_description=? , me_page_id=? , me_xpath=?, me_action_element=?, me_status=?, me_updated_by=? WHERE me_id =?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, a.MeName, a.MeDescription, a.MePageID, a.MeXpath, a.MeActionElement, a.MeStatus, a.MeUpdatedBy, id)

	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlElementRepository) Count(ctx context.Context) (count int, err error) {
	query := `SELECT count(me_id) as total from mr_elements`

	total, err := m.count(ctx, query)

	if err != nil {
		return 0, err
	}

	return total, nil
}




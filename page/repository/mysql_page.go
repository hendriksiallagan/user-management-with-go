package repository

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/user-management-with-go/page"
	"github.com/user-management-with-go/models"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" 
)

type mysqlPageRepository struct {
	Conn *sql.DB
}

func NewMysqlPageRepository(Conn *sql.DB) page.Repository {

	return &mysqlPageRepository{Conn}
}

func (m *mysqlPageRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Page, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Page, 0)
	for rows.Next() {
		s := new(models.Page)
		err = rows.Scan(
			&s.MpID,
			&s.MpCode,
			&s.MpName,
			&s.MpDescription,
			&s.MpFilepath,
			&s.MpStatus,
			&s.MpCreatedBy,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

func (m *mysqlPageRepository) count(ctx context.Context, query string, args ...interface{}) (count int, err error) {

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

func (m *mysqlPageRepository) Fetch(ctx context.Context, offset int64, limit int64, search string) (a []*models.Page, total int, err error) {
	query := `select mp_id, mp_code, mp_name, mp_description, mp_filepath, mp_status, mp_created_by from mr_pages where mp_name like ? ORDER BY mp_id desc limit ? offset ? `

	res, err := m.fetch(ctx, query, search, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	query2 := `select count(mp_id) from mr_pages where mp_name like ?  `

	total, err = m.count(ctx, query2, search)

	if err != nil {
		return nil, 0, err
	}

	return res, total, err

}

func (m *mysqlPageRepository) Store(ctx context.Context, a *models.Page, code string) error {

	query := `INSERT mr_pages SET mp_code=?, mp_description=? , mp_filepath=? , mp_name=? , mp_created_by=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, code, a.MpDescription, a.MpFilepath, a.MpName, a.MpCreatedBy)
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	a.MpID = lastId
	return nil
}

func (m *mysqlPageRepository) Update(ctx context.Context, a *models.Page, id int64) error {

	query := `UPDATE mr_pages SET mp_name=? , mp_description=? , mp_filepath=? , mp_status=?, mp_updated_by=? WHERE mp_id =?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, a.MpName, a.MpDescription, a.MpFilepath, a.MpStatus, a.MpUpdatedBy, id)

	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlPageRepository) Count(ctx context.Context) (count int, err error) {
	query := `SELECT count(mp_id) as total from mr_pages`

	total, err := m.count(ctx, query)

	if err != nil {
		return 0, err
	}

	return total, nil
}




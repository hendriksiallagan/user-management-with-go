package repository

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/user-management-with-go/menu"
	"github.com/user-management-with-go/models"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" 
)

type mysqlMenuRepository struct {
	Conn *sql.DB
}

func NewMysqlMenuRepository(Conn *sql.DB) menu.Repository {

	return &mysqlMenuRepository{Conn}
}

func (m *mysqlMenuRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Menu, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Menu, 0)
	for rows.Next() {
		s := new(models.Menu)
		err = rows.Scan(
			&s.MmID,
			&s.MmCode,
			&s.MmName,
			&s.MmDescription,
			&s.MmUrl,
			&s.MmHeaderID,
			&s.MmStatus,
			&s.MmCreatedBy,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

func (m *mysqlMenuRepository) count(ctx context.Context, query string, args ...interface{}) (count int, err error) {

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

func (m *mysqlMenuRepository) Fetch(ctx context.Context, offset int64, limit int64, search string) (a []*models.Menu, total int, err error) {
	query := `select mm_id, mm_code, mm_name, mm_description, mm_url, mm_header_id, mm_status, mm_created_by from mr_menus where mm_name like ? ORDER BY mm_id desc limit ? offset ? `

	res, err := m.fetch(ctx, query, search, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	query2 := `select count(mm_id) from mr_menus where mm_name like ?  `

	total, err = m.count(ctx, query2, search)

	if err != nil {
		return nil, 0, err
	}

	return res, total, err

}

func (m *mysqlMenuRepository) Store(ctx context.Context, a *models.Menu, code string) error {

	query := `INSERT mr_menus SET mm_code=?, mm_description=? , mm_url=? , mm_header_id=?, mm_name=? , mm_created_by=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, code, a.MmDescription, a.MmUrl, a.MmHeaderID, a.MmName, a.MmCreatedBy)
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	a.MmID = lastId
	return nil
}

func (m *mysqlMenuRepository) Update(ctx context.Context, a *models.Menu, id int64) error {

	query := `UPDATE mr_menus SET mm_name=? , mm_description=? , mm_url=? , mm_header_id=?, mm_status=?, mm_updated_by=? WHERE mm_id =?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, a.MmName, a.MmDescription, a.MmUrl, a.MmHeaderID, a.MmStatus, a.MmUpdatedBy, id)

	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlMenuRepository) Count(ctx context.Context) (count int, err error) {
	query := `SELECT count(mm_id) as total from mr_menus`

	total, err := m.count(ctx, query)

	if err != nil {
		return 0, err
	}

	return total, nil
}




package repository

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/user-management-with-go/actionelement"
	"github.com/user-management-with-go/models"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" 
)

type mysqlActionelementRepository struct {
	Conn *sql.DB
}

func NewMysqlActionelementRepository(Conn *sql.DB) actionelement.Repository {

	return &mysqlActionelementRepository{Conn}
}

func (m *mysqlActionelementRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Actionelement, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Actionelement, 0)
	for rows.Next() {
		s := new(models.Actionelement)
		err = rows.Scan(
			&s.MaeID,
			&s.MaeCode,
			&s.MaeName,
			&s.MaeDescription,
			&s.MaeScript,
			&s.MaeStatus,
			&s.MaeCreatedBy,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

func (m *mysqlActionelementRepository) count(ctx context.Context, query string, args ...interface{}) (count int, err error) {

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

func (m *mysqlActionelementRepository) Fetch(ctx context.Context, offset int64, limit int64, search string) (a []*models.Actionelement, total int, err error) {
	query := `select mae_id, mae_code, mae_name, mae_description, mae_script, mae_status, mae_created_by from mr_action_elements where mae_name like ? ORDER BY mae_id desc limit ? offset ? `

	res, err := m.fetch(ctx, query, search, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	query2 := `select count(mae_id) from mr_action_elements where mae_name like ?  `

	total, err = m.count(ctx, query2, search)

	if err != nil {
		return nil, 0, err
	}

	return res, total, err

}

func (m *mysqlActionelementRepository) Store(ctx context.Context, a *models.Actionelement, code string) error {

	query := `INSERT mr_action_elements SET mae_code=?, mae_description=? , mae_script=? , mae_name=? , mae_created_by=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, code, a.MaeDescription, a.MaeScript, a.MaeName, a.MaeCreatedBy)
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	a.MaeID = lastId
	return nil
}

func (m *mysqlActionelementRepository) Update(ctx context.Context, a *models.Actionelement, id int64) error {

	query := `UPDATE mr_action_elements SET mae_name=? , mae_description=? , mae_script=? , mae_status=?, mae_updated_by=? WHERE mae_id =?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, a.MaeName, a.MaeDescription, a.MaeScript, a.MaeStatus, a.MaeUpdatedBy, id)

	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlActionelementRepository) Count(ctx context.Context) (count int, err error) {
	query := `SELECT count(mae_id) as total from mr_action_elements`

	total, err := m.count(ctx, query)

	if err != nil {
		return 0, err
	}

	return total, nil
}




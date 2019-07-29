package repository

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"github.com/user-management-with-go/user"
	"github.com/user-management-with-go/models"
	//"fmt"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

func NewMysqlUserRepository(Conn *sql.DB) user.Repository {

	return &mysqlUserRepository{Conn}
}

func (m *mysqlUserRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.User, 0)
	for rows.Next() {
		s := new(models.User)
		err = rows.Scan(
			&s.MuID,
			&s.MuCode,
			&s.MuName,
			&s.MuEmail,
			&s.MuDescription,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

func (m *mysqlUserRepository) Fetch(ctx context.Context, offset uint64, limit uint64, search string) (res []*models.User, total int, err error) {
	query := `select mu_id, mu_code, mu_name, mu_email, mu_description from mr_users where mu_name like ? ORDER BY mu_id desc limit ? offset ? `

	res, err = m.fetch(ctx, query, search, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	query2 := `select count(mu_id) from mr_users where mu_name like ?  `

	total, err = m.count(ctx, query2, search)

	if err != nil {
		return nil, 0, err
	}

	return res, total, err
}

func (m *mysqlUserRepository) fetchPassword(ctx context.Context, query string, args ...interface{}) ([]*models.UserPassword, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.UserPassword, 0)
	for rows.Next() {
		s := new(models.UserPassword)
		err = rows.Scan(
			&s.LupID,
			&s.LupUserID,
			&s.LupUserPassword,
			&s.LupExpPasswordLink,
			&s.LupLinkExpDuration,
			&s.LupIsExpired,
			&s.LupCreatedBy,
			&s.LupCreatedAt,
			&s.LupStatus,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

func (m *mysqlUserRepository) FetchPasswordByID(ctx context.Context, offset uint64, limit uint64, id uint64) (res []*models.UserPassword, total int, err error) {
	query := `select lup_id, lup_user_id, lup_user_password, lup_exp_password_link, lup_link_expired_duration, lup_is_expired, lup_created_by, lup_created_at, lup_status FROM logs_user_password WHERE lup_user_id = ? ORDER BY lup_id desc limit ? offset ? `

	res, err = m.fetchPassword(ctx, query, id, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	query2 := `select count(lup_id) from logs_user_password `

	total, err = m.count(ctx, query2)

	if err != nil {
		return nil, 0, err
	}

	return res, total, err
}

func (m *mysqlUserRepository) fetchPin(ctx context.Context, query string, args ...interface{}) ([]*models.UserPIN, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.UserPIN, 0)
	for rows.Next() {
		s := new(models.UserPIN)
		err = rows.Scan(
			&s.LopinID,
			&s.LopinUserID,
			&s.LopinPIN,
			&s.LopinExpiredDate,
			&s.LopinCreatedBy,
			&s.LopinCreatedAt,
			&s.LopinStatus,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

func (m *mysqlUserRepository) FetchPinByID(ctx context.Context, offset uint64, limit uint64, id uint64) (res []*models.UserPIN, total int, err error) {
	query := `select lopin_id, lopin_user_id, lopin_pin, lopin_expired_date, lopin_created_by, lopin_created_at, lopin_status FROM logs_user_pin WHERE lopin_user_id = ? ORDER BY lopin_id desc limit ? offset ? `

	res, err = m.fetchPin(ctx, query, id, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	query2 := `select count(lopin_id) from logs_user_pin `

	total, err = m.count(ctx, query2)

	if err != nil {
		return nil, 0, err
	}

	return res, total, err
}

func (m *mysqlUserRepository) fetchRole(ctx context.Context, query string, args ...interface{}) ([]*models.UserRole, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.UserRole, 0)
	for rows.Next() {
		s := new(models.UserRole)
		err = rows.Scan(
			&s.MurID,
			&s.MurCode,
			&s.MurName,
			&s.MurUserID,
			&s.MurStatus,
			&s.MurCreatedBy,
			&s.MurCreatedAt,
			&s.MurUpdatedBy,
			&s.MurUpdatedAt,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

func (m *mysqlUserRepository) FetchRoleByID(ctx context.Context, offset uint64, limit uint64, id uint64) (res []*models.UserRole, total int, err error) {
	query := `select mur_id, mur_code, mur_name, mur_user_id, mur_status, mur_created_by, mur_created_at, mur_updated_by, mur_updated_at FROM mr_user_roles WHERE mur_user_id = ? ORDER BY mur_id desc limit ? offset ? `

	res, err = m.fetchRole(ctx, query, id, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	query2 := `select count(mur_id) from mr_user_roles WHERE mur_user_id = ? `

	total, err = m.count(ctx, query2, id)

	if err != nil {
		return nil, 0, err
	}

	return res, total, err
}

func (m *mysqlUserRepository) fetchStatusInfo(ctx context.Context, query string, args ...interface{}) ([]*models.UserStatusInfo, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.UserStatusInfo, 0)
	for rows.Next() {
		s := new(models.UserStatusInfo)
		err = rows.Scan(
			&s.LusiID,
			&s.LusiReason,
			&s.LusiStatusBefore,
			&s.LusiStatusCurrent,
			&s.LusiDuration,
			&s.LusiUserID,
			&s.LusiCreatedBy,
			&s.LusiCreatedAt,
			&s.LusiStatus,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

func (m *mysqlUserRepository) FetchStatusInfoByID(ctx context.Context, offset uint64, limit uint64, id uint64) (res []*models.UserStatusInfo, total int, err error) {
	query := `select lusi_id, lusi_reason, lusi_status_before, lusi_status_current, lusi_duration, lusi_user_id, lusi_created_by, lusi_created_at, lusi_status_type FROM logs_user_status_info WHERE lusi_user_id = ? ORDER BY lusi_id desc limit ? offset ? `

	res, err = m.fetchStatusInfo(ctx, query, id, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	query2 := `select count(lusi_id) from logs_user_status_info WHERE lusi_user_id = ? `

	total, err = m.count(ctx, query2, id)

	if err != nil {
		return nil, 0, err
	}

	return res, total, err
}

func (m *mysqlUserRepository) fetchOtp(ctx context.Context, query string, args ...interface{}) ([]*models.UserOtp, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.UserOtp, 0)
	for rows.Next() {
		s := new(models.UserOtp)
		err = rows.Scan(
			&s.LuoID,
			&s.LuoUserID,
			&s.LuoOtp,
			&s.LuoStatus,
			&s.LuoExpiredDate,
			&s.LuoCreatedBy,
			&s.LuoCreatedAt,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

func (m *mysqlUserRepository) FetchOtpByID(ctx context.Context, offset uint64, limit uint64, id uint64) (res []*models.UserOtp, total int, err error) {
	query := `select luo_id, luo_user_id, luo_otp, luo_status, luo_expired_date, luo_created_by, luo_created_at FROM logs_user_otp WHERE luo_user_id = ? ORDER BY luo_id desc limit ? offset ? `

	res, err = m.fetchOtp(ctx, query, id, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	query2 := `select count(luo_id) from logs_user_otp WHERE luo_user_id = ? `

	total, err = m.count(ctx, query2, id)

	if err != nil {
		return nil, 0, err
	}

	return res, total, err
}

func (m *mysqlUserRepository) fetchToken(ctx context.Context, query string, args ...interface{}) ([]*models.UserToken, error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.UserToken, 0)
	for rows.Next() {
		s := new(models.UserToken)
		err = rows.Scan(
			&s.LutID,
			&s.LutUserID,
			&s.LutToken,
			&s.LutStatus,
			&s.LutCreatedBy,
			&s.LutCreatedAt,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

func (m *mysqlUserRepository) FetchTokenByID(ctx context.Context, offset uint64, limit uint64, id uint64) (res []*models.UserToken, total int, err error) {
	query := `select lut_id, lut_user_id, lut_token, lut_status, lut_created_by, lut_created_at FROM logs_user_token WHERE lut_user_id = ? ORDER BY lut_id desc limit ? offset ? `

	res, err = m.fetchToken(ctx, query, id, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	query2 := `select count(lut_id) from logs_user_token WHERE lut_user_id = ? `

	total, err = m.count(ctx, query2, id)

	if err != nil {
		return nil, 0, err
	}

	return res, total, err
}

func (m *mysqlUserRepository) FetchByID(ctx context.Context, id int64) (res *models.User, err error) {
	query := `SELECT mu_id, mu_code, mu_name, mu_email, mu_description from mr_users WHERE mu_id = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return nil, models.ErrNotFound
	}

	return res, err
}

func (m *mysqlUserRepository) Store(ctx context.Context, a *models.User, code string) error {

	query := `INSERT mr_users SET mu_code=? , mu_name=? , mu_email=?, mu_description=?, mu_created_by=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, code, a.MuName, a.MuEmail, a.MuDescription, a.MuCreatedBy)
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	a.MuID = uint64(lastId)
	return nil
}

func (m *mysqlUserRepository) StoreRole(ctx context.Context, a *models.UserRole, code string) error {

	query := `INSERT mr_user_roles SET mur_code=? , mur_name=?, mur_user_id=? , mur_created_by=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, code, a.MurName, a.MurUserID, a.MurCreatedBy)
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	a.MurID = uint64(lastId)
	return nil
}

func (m *mysqlUserRepository) Update(ctx context.Context, a *models.User, id int64) error {

	query := `UPDATE mr_users SET mu_name=? , mu_email=? , mu_description=?, mu_status=?, mu_updated_by=? WHERE mu_id =?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, a.MuName, a.MuEmail, a.MuDescription, a.MuStatus, a.MuUpdatedBy, id)

	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlUserRepository) UpdateRole(ctx context.Context, a *models.UserRole, id uint64) error {

	query := `UPDATE mr_user_roles SET mur_name=? , mur_status=? , mur_updated_by=? WHERE mur_id =?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, a.MurName, a.MurStatus, a.MurUpdatedBy, id)

	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlUserRepository) UpdatePassword(ctx context.Context, a *models.ResetPassword, id uint64) error {

	query := `UPDATE mr_users SET mu_password=? , mu_updated_by=?  WHERE mu_id =?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, a.Password, a.UpdatedBy, id)
	if err != nil {
		return err
	}

	query2 := `INSERT logs_user_password SET lup_user_id=? , lup_user_password=?, lup_exp_password_link=? , lup_link_expired_duration=?, lup_is_expired=?, lup_created_by=?`
	stmt, err = m.Conn.PrepareContext(ctx, query2)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, a.UserID, a.Password, a.ExpPasswordLink, a.ExpDuration, a.IsExpired, a.CreatedBy)
	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlUserRepository) count(ctx context.Context, query string, args ...interface{}) (count int, err error) {

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

func (m *mysqlUserRepository) Count(ctx context.Context) (count int, err error) {
	query := `SELECT count(mu_id) as total from mr_users`

	total, err := m.count(ctx, query)

	if err != nil {
		return 0, err
	}

	return total, nil
}

func (m *mysqlUserRepository) CountRole(ctx context.Context) (count int, err error) {
	query := `SELECT count(mur_id) as total from mr_user_roles`

	total, err := m.count(ctx, query)

	if err != nil {
		return 0, err
	}

	return total, nil
}






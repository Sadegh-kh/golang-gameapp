package mysql

import (
	"database/sql"
	"fmt"
	"gameapp/entity"
	"gameapp/pkg/richerror"
)

func (d *MySQLDB) Register(u entity.User) (entity.User, error) {
	res, err := d.DB.Exec("insert into users(name,phone_number,password) value(?,?,?)", u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("%w", err)
	}
	// error always nil
	id, _ := res.LastInsertId()
	u.ID = uint(id)
	return u, nil

}

func (d *MySQLDB) IsPhoneNumberUniq(phoneNumber string) (bool, error) {

	row := d.DB.QueryRow("select * from users where phone_number = ?", phoneNumber)
	_, err := userScan(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("%w", err)
	}
	return false, nil
}

func (d *MySQLDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	row := d.DB.QueryRow("SELECT * FROM users where phone_number = ?", phoneNumber)

	user, err := userScan(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.RichError{
				Operation:    "mysql.GetUserByPhoneNumber",
				WrappedError: nil,
				Message:      "user not exist",
				Kind:         richerror.NotFound,
				Meta:         nil,
			}
		}

		return entity.User{}, richerror.RichError{
			Operation:    "mysql.GetUserByPhoneNumber",
			WrappedError: err,
			Message:      "unexpected error",
			Kind:         richerror.Unexpected,
			Meta: map[string]interface{}{
				"message": err.Error(),
			},
		}
	}

	return user, nil
}

func (d *MySQLDB) GetUserByID(uid uint) (entity.User, error) {
	row := d.DB.QueryRow("select * from users where id=?", uid)
	user, err := userScan(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.RichError{
				Operation:    "mysql.GetUserByID",
				WrappedError: nil,
				Message:      "user not found",
				Kind:         richerror.NotFound,
				Meta:         nil,
			}
		}

		return entity.User{}, richerror.RichError{
			Operation:    "mysql.GetUserByID",
			WrappedError: err,
			Message:      "unexpected error",
			Kind:         richerror.Unexpected,
			Meta: map[string]interface{}{
				"message": err.Error(),
			},
		}
	}

	return user, nil
}

func userScan(row *sql.Row) (entity.User, error) {
	var user entity.User
	var create_at []uint8
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &create_at, &user.Password)

	return user, err
}

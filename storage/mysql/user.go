package mysql

import (
	"database/sql"
	"fmt"
	"gameapp/entity"
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

func (d *MySQLDB) CheckUserExistAndGet(phoneNumber string) (entity.User, bool, error) {
	row := d.DB.QueryRow("SELECT * FROM users where phone_number = ?", phoneNumber)

	user, err := userScan(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}

		return entity.User{}, true, err
	}

	return user, true, nil
}

func (d *MySQLDB) GetUserByID(uid uint) (entity.User, error) {
	row := d.DB.QueryRow("select * from users where id=?", uid)
	user, err := userScan(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, fmt.Errorf("user not found")
		}

		return entity.User{}, err
	}

	return user, nil
}

func userScan(row *sql.Row) (entity.User, error) {
	var user entity.User
	var create_at []uint8

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &create_at, &user.Password)

	return user, err
}

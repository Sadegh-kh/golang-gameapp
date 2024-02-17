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
	user := entity.User{}
	var create_at []uint8
	row := d.DB.QueryRow("select * from users where phone_number = ?", phoneNumber)
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &create_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("%w", err)
	}
	return false, nil
}

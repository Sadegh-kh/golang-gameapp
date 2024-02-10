package mysql

import (
	"fmt"
	"gameapp/entity"
)

func (d *MySQLDB) Register(u entity.User) (entity.User, error) {
	res, err := d.DB.Exec("insert into users(name,phone_number,password) value(?,?,?)", u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't excute on db in register bucuse: %w", err)
	}
	// error always nil
	id, _ := res.LastInsertId()
	u.ID = uint(id)
	return u, nil

}

func (d *MySQLDB) IsPhoneNumberUniq(phoneNumber string) (bool, error) {
	panic("not implmented")
}

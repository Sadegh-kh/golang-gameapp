package mysql

import "gameapp/entity"

func (d MySQLDB) SaveUser(u entity.User) (entity.User, error) {}

func (d MySQLDB) IsPhoneNumberUniq(phoneNumber string) (bool, error) {}

package userservice

import (
	"fmt"
	"gameapp/entity"
	"gameapp/pkg/hashpassword"
)

type Storage interface{
	Validator
	SaveUser(u entity.User)(entity.User,error)

}
type Validator interface{
	IsPhoneNumberUniq(phoneNumber string)(bool,error)
}

type Service struct {
	storage Storage
}
type RegisterRequest struct {
	Name        string
	PhoneNumber string
	Password    string
}

type RegisterResponse struct {
	User entity.User
}

func (s Service) Register(req RegisterRequest) (RegisterResponse,error) {

	// validator
	// uniq phone number
	// 11 len of phon number
	// more than 6 len of password
	if isUniq,err:=s.storage.IsPhoneNumberUniq(req.PhoneNumber);err!=nil || !isUniq{
		if err!=nil{
			return RegisterResponse{},err
		}
		if !isUniq	{
			return RegisterResponse{},fmt.Errorf("phone number is not uniq")
		}
	}
	if len(req.Password)<6{
		return RegisterResponse{},fmt.Errorf("password should be more than 6 character")
	}

	// hash passwrd
	passHash:=hashpassword.EncodePasword(req.Password)

	// save to storage
	newUser,err:=s.storage.SaveUser(entity.User{
		ID: 0,
		Name: req.Name,
		PhoneNumber: req.PhoneNumber,
		Password: passHash,
	})
	if err!=nil{
		return RegisterResponse{},err
	}

	// return new user
	return RegisterResponse{User: newUser},nil
}
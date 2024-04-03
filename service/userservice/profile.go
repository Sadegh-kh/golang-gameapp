package userservice

import "gameapp/param"

func (s Service) Profile(req param.ProfileRequest) (param.ProfileResponse, error) {
	user, err := s.storage.GetUserByID(req.UserID)
	if err != nil {
		return param.ProfileResponse{}, err
	}

	return param.ProfileResponse{Name: user.Name}, nil

}

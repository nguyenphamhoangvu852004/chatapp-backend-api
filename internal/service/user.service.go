package service

import "chapapp-backend-api/internal/reporitory"

type UserService struct {
	userRepo *reporitory.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: reporitory.NewUserRepo(),
	}
}

func (us *UserService) GetUserById() string {
	return us.userRepo.GetUserById()
}

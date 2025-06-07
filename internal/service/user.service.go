package service

import (
	"chapapp-backend-api/internal/reporitory"
	"chapapp-backend-api/pkg/response"
)

type IUserService interface {
	Register(email string, username string, password string) int
}

type userService struct {
	userRepo reporitory.IUserRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, username string, password string) int {
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeParamInvalid
	}
	return response.ErrCodeSuccess
}

func NewUserService(userRepo reporitory.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

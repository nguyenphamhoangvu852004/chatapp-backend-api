package controller

import (
	"chapapp-backend-api/internal/service"
	"chapapp-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("fdjk", "kdsajf", "123123")
	response.SuccessReponse(c, result, nil)
}

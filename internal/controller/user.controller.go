package controller

import (
	"chapapp-backend-api/internal/service"
	"chapapp-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	// response.SuccessReponse(c, 20001, uc.userService.GetUserById())
	response.ErrorReponse(c, 20003, "Loi roiii neee")
}

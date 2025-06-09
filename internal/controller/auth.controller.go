package controller

import (
	"chapapp-backend-api/internal/dto/auth"
	exception "chapapp-backend-api/internal/exeption"
	"chapapp-backend-api/internal/service"
	"chapapp-backend-api/pkg/response"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (uc *AuthController) Register(c *gin.Context) {
	var inputDto auth.RegisterInputDTO
	fmt.Println(inputDto)
	if err := c.BindJSON(&inputDto); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := uc.authService.Register(inputDto)
	if err != nil {
		var customErr *exception.CustomError
		if errors.As(err, &customErr) {
			response.ErrorReponse(c, customErr.Code, customErr.Message)
		} else {
			response.ErrorReponse(c, 500, "internal server error")
		}
		return
	}
	response.SuccessReponse(c, http.StatusCreated, result)
}
func (uc *AuthController) VerifyOTP(c *gin.Context) {
	var inputDto auth.VerifyOTPInputDTO
	if err := c.BindJSON(&inputDto); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := uc.authService.VerifyOTP(inputDto)
	if err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessReponse(c, http.StatusOK, result)
}
func (uc *AuthController) Login(c *gin.Context) {

	var inputDto auth.LoginInputDTO
	if err := c.BindJSON(&inputDto); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := uc.authService.Login(inputDto)
	if err != nil {
		var customErr *exception.CustomError
		if errors.As(err, &customErr) {
			response.ErrorReponse(c, customErr.Code, customErr.Message)
		} else {
			response.ErrorReponse(c, 500, "internal server error")
		}
		return
	}
	response.SuccessReponse(c, http.StatusOK, result)
}
func (uc *AuthController) SendOTP(c *gin.Context) {
	// tạo dt
	var inputDto auth.SendOTPInputDTO
	if err := c.BindJSON(&inputDto); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// gọi service
	result, err := uc.authService.SendOTP(inputDto)
	if err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// trả về client
	response.SuccessReponse(c, http.StatusAccepted, result)
}

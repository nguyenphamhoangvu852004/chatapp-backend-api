package controller

import (
	"chapapp-backend-api/internal/dto"
	exception "chapapp-backend-api/internal/exeption"
	"chapapp-backend-api/internal/service"
	"chapapp-backend-api/pkg/response"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService service.IAccountService
}

func (accountController *AccountController) ChangePassword(c *gin.Context) {
	var inputDTO dto.ChangePasswordInputDTO
	userRaw, exists := c.Get("user")
	if !exists {
		response.ErrorReponse(c, http.StatusUnauthorized, "No user in context")
		return
	}

	userMap, ok := userRaw.(map[string]interface{})
	if !ok {
		response.ErrorReponse(c, http.StatusInternalServerError, "User data format error")
	}

	idFloat, ok := userMap["id"].(float64)
	if !ok {
		response.ErrorReponse(c, http.StatusInternalServerError, "ID data format errror")
		return
	}

	// Convert float64 id to int, then to string
	inputDTO.Id = fmt.Sprintf("%.0f", idFloat)
	if err := c.BindJSON(&inputDTO); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Println(inputDTO)

	result, err := accountController.accountService.ChangePassword(inputDTO)
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

func (accountController *AccountController) GetRandomList(c *gin.Context) {
	var inputDTO dto.GetRamdonAccountInputDTO
	if err := c.BindQuery(&inputDTO); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := accountController.accountService.GetRandomList(inputDTO)
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

func (accountController *AccountController) GetDetail(c *gin.Context) {
	accountId := c.Param("id")
	result, err := accountController.accountService.GetDetail(accountId)
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

func (accountController *AccountController) GetList(c *gin.Context) {
	var inputDTO dto.GetListAccountInputDTO
	if err := c.BindQuery(&inputDTO); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(inputDTO)
	result, err := accountController.accountService.GetList(inputDTO)
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

func NewAccountController(accountService service.IAccountService) *AccountController {
	return &AccountController{
		accountService: accountService,
	}
}

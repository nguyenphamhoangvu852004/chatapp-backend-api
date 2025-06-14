package controller

import (
	"chapapp-backend-api/internal/dto"
	exception "chapapp-backend-api/internal/exeption"
	"chapapp-backend-api/internal/service"
	"chapapp-backend-api/pkg/response"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService service.IAccountService
}

func (accountController *AccountController) GetRandomList(c *gin.Context) {
	accountId := c.Param("id")
	result, err := accountController.accountService.GetRandomList(accountId)
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
	if err := c.ShouldBindQuery(&inputDTO); err != nil {
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

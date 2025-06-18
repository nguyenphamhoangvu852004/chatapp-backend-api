package controller

import (
	"chapapp-backend-api/internal/dto"
	exception "chapapp-backend-api/internal/exeption"
	"chapapp-backend-api/internal/service"
	"chapapp-backend-api/pkg/response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BanController struct {
	banService service.IBanService
}

func (b *BanController) GetList(c *gin.Context) {
	var inputDTO dto.GetListBanInputDTO
	if err := c.ShouldBindQuery(&inputDTO); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
	}
	if inputDTO.Page == 0 {
		inputDTO.Page = 1
	}
	if inputDTO.Limit == 0 {
		inputDTO.Limit = 10
	}
	res, err := b.banService.GetListBan(inputDTO)
	if err != nil {
		var customErr *exception.CustomError
		if errors.As(err, &customErr) {
			response.ErrorReponse(c, customErr.Code, customErr.Message)
		} else {
			response.ErrorReponse(c, 500, "internal server error")
		}
		return
	}
	response.SuccessReponse(c, http.StatusOK, res)
}

func (b *BanController) Delete(c *gin.Context) {
	var inputDTO dto.DeleteBanInputDTO
	if err := c.ShouldBindBodyWithJSON(&inputDTO); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
	}
	res, err := b.banService.Delete(inputDTO)
	if err != nil {
		var customErr *exception.CustomError
		if errors.As(err, &customErr) {
			response.ErrorReponse(c, customErr.Code, customErr.Message)
		} else {
			response.ErrorReponse(c, 500, "internal server error")
		}
		return
	}
	response.SuccessReponse(c, http.StatusOK, res)
}

func (b *BanController) Create(c *gin.Context) {
	var inputDTO dto.CreateBanInputDTO
	if err := c.ShouldBindBodyWithJSON(&inputDTO); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
	}
	res, err := b.banService.Create(inputDTO)
	if err != nil {
		var customErr *exception.CustomError
		if errors.As(err, &customErr) {
			response.ErrorReponse(c, customErr.Code, customErr.Message)
		} else {
			response.ErrorReponse(c, 500, "internal server error")
		}
		return
	}
	response.SuccessReponse(c, http.StatusOK, res)
}

func NewBanController(banService service.IBanService) *BanController {
	return &BanController{
		banService: banService,
	}
}

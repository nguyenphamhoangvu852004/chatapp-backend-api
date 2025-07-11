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

type MessageController struct {
	messageService service.IMessageService
}

func (m *MessageController) Delete(c *gin.Context) {
	var inputDTO dto.DeleteMessageInputDTO
	if err := c.ShouldBindBodyWithJSON(&inputDTO); err != nil {  
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	res, err := m.messageService.Delete(inputDTO)
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

func (m *MessageController) Create(c *gin.Context) {
	var inputDTO dto.CreateMessageInputDTO
	if err := c.ShouldBindBodyWithJSON(&inputDTO); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := m.messageService.Create(inputDTO)

	if err != nil {
		var customErr *exception.CustomError
		if errors.As(err, &customErr) {
			response.ErrorReponse(c, customErr.Code, customErr.Message)
		} else {
			response.ErrorReponse(c, 500, "internal server error")
		}
		return
	}
	response.SuccessReponse(c, http.StatusCreated, res)
}
func (c *MessageController) GetMessages(ctx *gin.Context) {
	var input dto.GetListMessageInputDTO
	if err := ctx.ShouldBindQuery(&input); err != nil {
		response.ErrorReponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result, err := c.messageService.GetList(input)
	if err != nil {
		response.ErrorReponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessReponse(ctx, http.StatusOK, result)
}

func NewMessageController(messageService service.IMessageService) *MessageController {
	return &MessageController{
		messageService: messageService,
	}
}

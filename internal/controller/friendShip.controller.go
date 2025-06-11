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

type FriendShipController struct {
	friendShipService service.IFriendShipService
}

func NewFriendShipController(friendShipService service.IFriendShipService) *FriendShipController {
	return &FriendShipController{friendShipService: friendShipService}

}

func (controllet *FriendShipController) GetListReceiveFriendShips(c *gin.Context) {
	id := c.Param("id")
	result, err := controllet.friendShipService.GetListReceiveFriendShips(id)
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

func (controller *FriendShipController) GetListFriendShipsOfAccount(c *gin.Context) {

	id := c.Param("id")

	result, err := controller.friendShipService.GetListFriendShipsOfAccount(id)
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
func (controller *FriendShipController) GetListSendFriendShips(c *gin.Context) {
	id := c.Param("id")
	result, err := controller.friendShipService.GetListSendFriendShips(id)
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
func (controller *FriendShipController) GetList(c *gin.Context) {
	var inputDTO dto.GetListFriendShipInputDTO
	if err := c.ShouldBindQuery(&inputDTO); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := controller.friendShipService.GetList(inputDTO)
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

func (controller *FriendShipController) Update(c *gin.Context) {
	var inputDTO dto.UpdateFriendShipInputDTO
	if err := c.ShouldBindBodyWithJSON(&inputDTO); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := controller.friendShipService.Update(inputDTO)
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

func (controller *FriendShipController) Create(c *gin.Context) {
	var inputDTO dto.CreateFriendShipInputDTO
	if err := c.ShouldBindBodyWithJSON(&inputDTO); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := controller.friendShipService.Create(inputDTO)
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

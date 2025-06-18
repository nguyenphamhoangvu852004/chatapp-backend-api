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

type ConversationController struct {
	ConversationService service.IConversationSerivce
}

func (controller *ConversationController) GetMembers(c *gin.Context) {
	conversationId := c.Param("id")
	if conversationId == "" {
		response.ErrorReponse(c, http.StatusBadRequest, "Missing conversationId")
		return
	}

	res, err := controller.ConversationService.GetConversationMembers(conversationId)
	if err != nil {
		response.ErrorReponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessReponse(c, http.StatusOK, res)
}

func (controller *ConversationController) ModifyConversation(c *gin.Context) {
	var inputDTO dto.ModifyConversationInputDTO

	// ✅ Bind form-data text (ownerId, conversationId, name)
	if err := c.ShouldBind(&inputDTO); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// ✅ Lấy avatarUrl đã được middleware upload lên cloud
	if avatarURL, ok := c.Get("avatarUrl"); ok {
		if avatarStr, ok := avatarURL.(string); ok {
			inputDTO.AvatarURL = &avatarStr
		}
	}

	// ✅ Gọi service
	res, err := controller.ConversationService.ModifyConversation(inputDTO)
	if err != nil {
		response.ErrorReponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessReponse(c, http.StatusOK, res)
}

func (controller *ConversationController) GetGroupsJoinedByMe(c *gin.Context) {
	res, err := controller.ConversationService.GetGroupsJoinedByMe(c.Param("id"))
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
func (controller *ConversationController) DeleteMembers(c *gin.Context) {
	var inputDTO dto.RemoveMembersInputDTO
	if err := c.ShouldBindBodyWithJSON(&inputDTO); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
	}
	res, err := controller.ConversationService.RemoveMembers(inputDTO)
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

func (controller *ConversationController) Delete(c *gin.Context) {
	var inputDTO dto.DeleteMessageGroupInputDTO
	if err := c.ShouldBindBodyWithJSON(&inputDTO); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, err.Error())
	}
	res, err := controller.ConversationService.Delete(inputDTO)
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

func (controller *ConversationController) AddMembers(c *gin.Context) {
	var input dto.AddMemberInputDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorReponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}
	result, err := controller.ConversationService.AddMembers(input)
	if err != nil {
		var customErr *exception.CustomError
		if errors.As(err, &customErr) {
			response.ErrorReponse(c, customErr.Code, customErr.Message)
		} else {
			response.ErrorReponse(c, http.StatusInternalServerError, "Internal server error")
		}
		return
	}
	response.SuccessReponse(c, http.StatusOK, result)
}

func (controller *ConversationController) GetListOwnedByMe(c *gin.Context) {
	var ownerId = c.Param("id")
	result, err := controller.ConversationService.GetGroupListWhereUserIsAdmin(ownerId)
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

func (controller *ConversationController) Create(c *gin.Context) {
	var inputDTO = dto.CreateConversationInputDTO{
		Name:    c.PostForm("name"),
		OwnerId: c.PostForm("ownerId"),
		// AvatarURL: "hihihihihi",
		AvatarURL: c.GetString("avatarUrl"),
	}

	result, err := controller.ConversationService.Create(inputDTO)
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

func NewConversationController(conversationService service.IConversationSerivce) *ConversationController {
	return &ConversationController{ConversationService: conversationService}
}

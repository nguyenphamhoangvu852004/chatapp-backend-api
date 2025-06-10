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

type ProfileController struct {
	profileService service.IProfileService
}

func (controller *ProfileController) Update(c *gin.Context) {

	coverUrl := c.GetString("coverUrl")   // lấy từ middleware
	avatarUrl := c.GetString("avatarUrl") // lấy từ middleware

	// coverUrl := "test cover thoi nhe"   // lấy từ middleware á nha
	// avatarUrl := "test avatar thoi nhe" // lấy từ middleware á nha

	var inputDto = dto.UpdateProfileInputDTO{
		ProfileId: c.Param("id"),
		Bio:       c.PostForm("bio"),
		FullName:  c.PostForm("fullname"),
		CoverURL:  coverUrl,
		AvatarURL: avatarUrl,
	}

	result, err := controller.profileService.Update(inputDto)
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

func NewProfileController(profileService service.IProfileService) *ProfileController {
	return &ProfileController{profileService: profileService}
}

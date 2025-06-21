package dto

type (
	GetRamdonAccountInputDTO struct {
		Me       string `form:"me"`
		GetBlock bool   `form:"getBlock"`
	}
)

type (
	GetListAccountInputDTO struct {
		Me       string `form:"me"`
		GetBlock bool   `form:"getBlock"`
		Phone    string `form:"phone"`
	}
)

type (
	GetAccountDetailOutputDTO struct {
		Id          string                    `json:"id"`
		Username    string                    `json:"username"`
		Email       string                    `json:"email"`
		PhoneNumber string                    `json:"phoneNumber"`
		IsBanned    bool                      `json:"isBanned"`
		Profile     GetProfileDetailOutputDTO `json:"profile"`
	}
)

type (
	ChangePasswordInputDTO struct {
		Id              string `json:"id"`
		OldPassword     string `json:"oldPassword"`
		NewPassword     string `json:"newPassword"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	ChangePasswordOutputDTO struct {
		Id string `json:"id"`
		IsSuccess bool `json:"isSuccess"`
	}
)

package dto

type GetAccountDetailOutputDTO struct {
	Id          string                            `json:"id"`
	Username    string                            `json:"username"`
	Email       string                            `json:"email"`
	PhoneNumber string                            `json:"phoneNumber"`
	Profile     GetProfileDetailOutputDTO `json:"profile"`
}

package dto

type (
	CreateBlockInputDTO struct {
		BlockerId string `json:"blockerId"`
		BlockedId string `json:"blockedId"`
	}
	CreateBlockOutputDTO struct {
		BlockerId string `json:"blockerId"`
		BlockedId string `json:"blockedId"`
		IsSuccess bool   `json:"isSuccess"`
	}
)

type (
	DeleteBlockInputDTO struct {
		BlockerId string `json:"blockerId"`
		BlockedId string `json:"blockedId"`
	}
	DeleteBlockOutputDTO struct {
		BlockerId string `json:"blockerId"`
		BlockedId string `json:"blockedId"`
		IsSuccess bool   `json:"isSuccess"`
	}
)

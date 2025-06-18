package dto

type (
	CreateBanInputDTO struct {
		AdminId   string `json:"adminId"`
		AccountId string `json:"accountId"`
	}
	CreateBanOutputDTO struct {
		AccountId string `json:"accountId"`
		IsBanned  bool   `json:"isBanned"`
	}
)

type (
	DeleteBanInputDTO struct {
		AdminId   string `json:"adminId"`
		AccountId string `json:"accountId"`
	}
	DeleteBanOutputDTO struct {
		AccountId string `json:"accountId"`
		IsBanned  bool   `json:"isBanned"`
	}
)

type (
	GetListBanInputDTO struct {
		PaginationInputDTO
	}
	GetListBanOutpuDTO struct {
		List []GetAccountDetailOutputDTO `json:"list"`
	}
)

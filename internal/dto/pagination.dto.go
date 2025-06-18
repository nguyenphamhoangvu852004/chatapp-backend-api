package dto

type (
	PaginationInputDTO struct {
		Page  int `form:"page"`
		Limit int `form:"limit"`
	}
)

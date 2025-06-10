package service

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/reporitory"
	"fmt"
	"time"
)

type IAccountService interface {
	GetDetail(id string) (dto.GetAccountDetailOutputDTO, error)
	GetList() ([]dto.GetAccountDetailOutputDTO, error)
}

type accountService struct {
	accountRepo reporitory.IAccountRepository
}

// GetList implements IAccountService.
func (a *accountService) GetList() ([]dto.GetAccountDetailOutputDTO, error) {
	res, err := a.accountRepo.GetList()
	if err != nil {
		return nil, err
	}
	var result []dto.GetAccountDetailOutputDTO
	for _, v := range res {
		result = append(result, dto.GetAccountDetailOutputDTO{
			Id:          fmt.Sprintf("%d", v.ID),
			Username:    v.Username,
			Email:       v.Email,
			PhoneNumber: v.PhoneNumber,
			Profile: dto.GetProfileDetailOutputDTO{
				Id:        fmt.Sprintf("%d", v.Profile.ID),
				FullName:  v.Profile.FullName,
				Bio:       v.Profile.Bio,
				AvatarURL: v.Profile.AvatarURL,
				CoverURL:  v.Profile.CoverURL,
				CreatedAt: v.Profile.CreatedAt.Format(time.RFC3339),
				UpdatedAt: v.Profile.UpdatedAt.Format(time.RFC3339),
			},
		})
	}
	return result, nil
}

// GetDetail implements IAccountService.
func (a *accountService) GetDetail(id string) (dto.GetAccountDetailOutputDTO, error) {
	res, err := a.accountRepo.GetUserByAccountId(id)
	if err != nil {
		return dto.GetAccountDetailOutputDTO{}, err
	}
	return dto.GetAccountDetailOutputDTO{
		Id:          fmt.Sprintf("%d", res.ID),
		Username:    res.Username,
		Email:       res.Email,
		PhoneNumber: res.PhoneNumber,
		Profile: dto.GetProfileDetailOutputDTO{
			Id:        fmt.Sprintf("%d", res.Profile.ID),
			FullName:  res.Profile.FullName,
			Bio:       res.Profile.Bio,
			AvatarURL: res.Profile.AvatarURL,
			CoverURL:  res.Profile.CoverURL,
			CreatedAt: res.Profile.CreatedAt.Format(time.RFC3339),
			UpdatedAt: res.Profile.UpdatedAt.Format(time.RFC3339),
		},
	}, nil
}

func NewAccountService(accountRepo reporitory.IAccountRepository) IAccountService {
	return &accountService{accountRepo: accountRepo}
}

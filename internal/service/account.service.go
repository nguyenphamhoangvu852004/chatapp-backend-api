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
	GetRandomList(accountId string) ([]dto.GetAccountDetailOutputDTO, error)
}

type accountService struct {
	accountRepo reporitory.IAccountRepository
	blockRepo   reporitory.IBlockRepository
}

// GetRandomList implements IAccountService.
func (a *accountService) GetRandomList(accountId string) ([]dto.GetAccountDetailOutputDTO, error) {
	accounts, err := a.accountRepo.GetRandomFive()
	if err != nil {
		return nil, err
	}

	var accountIdUint uint
	fmt.Sscanf(accountId, "%d", &accountIdUint)

	var listOutDTO []dto.GetAccountDetailOutputDTO
	for _, acc := range accounts {
		if acc.ID == accountIdUint {
			continue
		}
		listOutDTO = append(listOutDTO, dto.GetAccountDetailOutputDTO{
			Id:          fmt.Sprintf("%d", acc.ID),
			Username:    acc.Username,
			Email:       acc.Email,
			PhoneNumber: acc.PhoneNumber,
			Profile: dto.GetProfileDetailOutputDTO{
				Id:        fmt.Sprintf("%d", acc.Profile.ID),
				FullName:  acc.Profile.FullName,
				Bio:       acc.Profile.Bio,
				AvatarURL: acc.Profile.AvatarURL,
				CoverURL:  acc.Profile.CoverURL,
			},
		})
	}

	return listOutDTO, nil
}

// GetList implements IAccountService.
func (a *accountService) GetList() ([]dto.GetAccountDetailOutputDTO, error) {
	res, err := a.accountRepo.GetList()
	if err != nil {
		return []dto.GetAccountDetailOutputDTO{}, err
	}

	var outputDTO []dto.GetAccountDetailOutputDTO
	for _, entity := range res {
		var account dto.GetAccountDetailOutputDTO
		account.Id = fmt.Sprintf("%d", entity.ID)
		account.Username = entity.Username
		account.Email = entity.Email
		account.PhoneNumber = entity.PhoneNumber
		account.Profile = dto.GetProfileDetailOutputDTO{
			Id:        fmt.Sprintf("%d", entity.Profile.ID),
			FullName:  entity.Profile.FullName,
			Bio:       entity.Profile.Bio,
			AvatarURL: entity.Profile.AvatarURL,
			CoverURL:  entity.Profile.CoverURL,
		}
		outputDTO = append(outputDTO, account)
	}
	return outputDTO, nil
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

func NewAccountService(accountRepo reporitory.IAccountRepository, blockRepo reporitory.IBlockRepository) IAccountService {
	return &accountService{accountRepo: accountRepo,
		blockRepo: blockRepo}
}

package service

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/entity"
	"chapapp-backend-api/internal/reporitory"
	"fmt"
	"time"
)

type IAccountService interface {
	GetDetail(id string) (dto.GetAccountDetailOutputDTO, error)
	GetList(data dto.GetListAccountInputDTO) ([]dto.GetAccountDetailOutputDTO, error)
}

type accountService struct {
	accountRepo reporitory.IAccountRepository
	blockRepo   reporitory.IBlockRepository
}

// GetList implements IAccountService.
func (a *accountService) GetList(data dto.GetListAccountInputDTO) ([]dto.GetAccountDetailOutputDTO, error) {
	accounts, err := a.accountRepo.GetList()
	if err != nil {
		return nil, err
	}

	// Get danh sách đã block hoặc bị block
	var currentUserIDUint uint
	fmt.Sscanf(data.CurrentUserId, "%d", &currentUserIDUint)
	blockedIDs, _ := a.blockRepo.GetListBlocker(currentUserIDUint)
	blockedMeIDs, _ := a.blockRepo.GetListBlocked(currentUserIDUint)

	// Dùng map cho nhanh
	blockMap := map[string]bool{}
	for _, block := range blockedIDs {
		blockMap[fmt.Sprintf("%v", block.BlockedID)] = true
	}
	for _, block := range blockedMeIDs {
		blockMap[fmt.Sprintf("%v", block.BlockerID)] = true
	}

	// Filter
	var result []entity.Account
	for _, acc := range accounts {
		if !blockMap[fmt.Sprintf("%v", acc.ID)] {
			result = append(result, acc)
		}
	}

	var listOutDTO []dto.GetAccountDetailOutputDTO
	for _, acc := range result {
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
				CreatedAt: acc.Profile.CreatedAt.Format(time.RFC3339),
				UpdatedAt: acc.Profile.UpdatedAt.Format(time.RFC3339),
			},
		})
	}

	return listOutDTO, nil
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

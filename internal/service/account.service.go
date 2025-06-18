package service

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/reporitory"
	"fmt"
	"time"
)

type IAccountService interface {
	GetDetail(id string) (dto.GetAccountDetailOutputDTO, error)
	GetList(data dto.GetListAccountInputDTO) ([]dto.GetAccountDetailOutputDTO, error)
	GetRandomList(data dto.GetRamdonAccountInputDTO) ([]dto.GetAccountDetailOutputDTO, error)
}

type accountService struct {
	accountRepo reporitory.IAccountRepository
	blockRepo   reporitory.IBlockRepository
}

// GetRandomList implements IAccountService.
func (a *accountService) GetRandomList(data dto.GetRamdonAccountInputDTO) ([]dto.GetAccountDetailOutputDTO, error) {
	accounts, err := a.accountRepo.GetRandomFive(data.Me)
	if err != nil {
		return nil, err
	}

	var myID uint
	fmt.Sscanf(data.Me, "%d", &myID)

	// Tập hợp các ID cần loại bỏ
	blockedIDs := make(map[uint]bool)
	if data.GetBlock {
		// Người đã block mình
		blockedList, err := a.blockRepo.GetListBlocked(myID)
		if err != nil {
			return nil, err
		}
		for _, b := range blockedList {
			blockedIDs[b.BlockerID] = true
		}

		// Người mình đã block
		blockerList, err := a.blockRepo.GetListBlocker(myID)
		if err != nil {
			return nil, err
		}
		for _, b := range blockerList {
			blockedIDs[b.BlockedID] = true
		}
	}

	var listOutDTO []dto.GetAccountDetailOutputDTO
	for _, acc := range accounts {
		if acc.ID == myID || blockedIDs[acc.ID] {
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
func (a *accountService) GetList(data dto.GetListAccountInputDTO) ([]dto.GetAccountDetailOutputDTO, error) {
	res, err := a.accountRepo.GetList(data)
	if err != nil {
		return nil, err
	}

	var myID uint
	fmt.Sscanf(data.Me, "%d", &myID)

	// Map để kiểm tra nhanh các ID cần loại bỏ
	blockedMap := make(map[uint]bool)

	if data.GetBlock {
		// Người mình đã block
		blockedList, err := a.blockRepo.GetListBlocked(myID)
		if err != nil {
			return nil, err
		}
		for _, b := range blockedList {
			blockedMap[b.BlockerID] = true
		}

		// Người đã block mình
		blockerList, err := a.blockRepo.GetListBlocker(myID)
		if err != nil {
			return nil, err
		}
		for _, b := range blockerList {
			blockedMap[b.BlockedID] = true
		}
	}

	var outputDTO []dto.GetAccountDetailOutputDTO
	for _, entity := range res {
		if blockedMap[entity.ID] {
			continue // loại bỏ người bị mình block hoặc đã block mình
		}

		outputDTO = append(outputDTO, dto.GetAccountDetailOutputDTO{
			Id:          fmt.Sprintf("%d", entity.ID),
			Username:    entity.Username,
			Email:       entity.Email,
			PhoneNumber: entity.PhoneNumber,
			Profile: dto.GetProfileDetailOutputDTO{
				Id:        fmt.Sprintf("%d", entity.Profile.ID),
				FullName:  entity.Profile.FullName,
				Bio:       entity.Profile.Bio,
				AvatarURL: entity.Profile.AvatarURL,
				CoverURL:  entity.Profile.CoverURL,
			},
		})
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

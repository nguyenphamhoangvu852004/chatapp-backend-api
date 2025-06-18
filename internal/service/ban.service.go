package service

import (
	"chapapp-backend-api/internal/dto"
	exception "chapapp-backend-api/internal/exeption"
	"chapapp-backend-api/internal/reporitory"
	"fmt"
	"net/http"
)

type IBanService interface {
	Create(data dto.CreateBanInputDTO) (dto.CreateBanOutputDTO, error)
	Delete(data dto.DeleteBanInputDTO) (dto.DeleteBanOutputDTO, error)
	GetListBan(data dto.GetListBanInputDTO) (dto.GetListBanOutpuDTO, error)
}

type banService struct {
	banRepo     reporitory.IBanRepository
	accountRepo reporitory.IAccountRepository
}

// GetListBan implements IBanService.
func (b *banService) GetListBan(data dto.GetListBanInputDTO) (dto.GetListBanOutpuDTO, error) {
	res, err := b.accountRepo.GetListBan(data)
	if err != nil {
		return dto.GetListBanOutpuDTO{}, err
	}

	var outputDTO dto.GetListBanOutpuDTO
	for _, account := range res {
		outputDTO.List = append(outputDTO.List, dto.GetAccountDetailOutputDTO{
			Id:          fmt.Sprintf("%d", account.ID),
			Username:    account.Username,
			Email:       account.Email,
			PhoneNumber: account.PhoneNumber,
			IsBanned:    account.IsBanned,
			Profile: dto.GetProfileDetailOutputDTO{
				Id:        fmt.Sprintf("%d", account.Profile.ID),
				FullName:  account.Profile.FullName,
				AvatarURL: account.Profile.AvatarURL,
				CoverURL:  account.Profile.CoverURL,
			},
		})
	}
	return outputDTO, nil
}

// Delete implements IBanService.
func (b *banService) Delete(data dto.DeleteBanInputDTO) (dto.DeleteBanOutputDTO, error) {
	// kiem ra coi no co bi block hong, neu khong thi return

	accountEntity, err := b.accountRepo.GetUserByAccountId(data.AccountId)

	if err != nil {
		return dto.DeleteBanOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "Not found account")
	}
	if !accountEntity.IsBanned {
		return dto.DeleteBanOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "Account is not banned")
	} else {
		accountEntity.IsBanned = false
		updatedAccount, err := b.accountRepo.Update(accountEntity)
		if err != nil {
			return dto.DeleteBanOutputDTO{}, err
		}
		return dto.DeleteBanOutputDTO{
			AccountId: fmt.Sprintf("%d", updatedAccount.ID),
			IsBanned:  updatedAccount.IsBanned,
		}, nil
	}

}

// Create implements IBanService.
func (b *banService) Create(data dto.CreateBanInputDTO) (dto.CreateBanOutputDTO, error) {
	// kiểm tra coi nó có bị block chưa, nếu rồi thì return luôn
	accountEntity, err := b.accountRepo.GetUserByAccountId(data.AccountId)
	if err != nil {
		return dto.CreateBanOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "Not found account")
	}
	if accountEntity.IsBanned {
		return dto.CreateBanOutputDTO{}, exception.NewCustomError(http.StatusConflict, "Account is already banned")
	} else {
		accountEntity.IsBanned = true
		updatedAccount, err := b.accountRepo.Update(accountEntity)
		if err != nil {
			return dto.CreateBanOutputDTO{}, err
		}
		return dto.CreateBanOutputDTO{
			AccountId: fmt.Sprintf("%d", updatedAccount.ID),
			IsBanned:  updatedAccount.IsBanned,
		}, nil
	}

}

func NewBanService(banRepo reporitory.IBanRepository, accountRepo reporitory.IAccountRepository) IBanService {
	return &banService{
		banRepo:     banRepo,
		accountRepo: accountRepo,
	}
}

package reporitory

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/entity"
)

type IAccountRepository interface {
	GetUserByEmail(email string) (entity.Account, error)
	GetUserByUsername(username string) (entity.Account, error)
	Create(account entity.Account) (entity.Account, error)
}

type accountRepository struct {
}

// GetUserByUsername implements IAccountRepository.
func (a *accountRepository) GetUserByUsername(username string) (entity.Account, error) {
	var account entity.Account
	err := global.Mdb.Where("username = ?", username).First(&account).Error
	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}

func (a *accountRepository) Create(account entity.Account) (entity.Account, error) {
	result := global.Mdb.Create(&account)
	if result.Error != nil {
		return entity.Account{}, result.Error
	}
	return account, nil
}

func (a *accountRepository) GetUserByEmail(email string) (entity.Account, error) {
	var account entity.Account
	err := global.Mdb.Where("email = ?", email).First(&account).Error
	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}

func NewAccountRepository() IAccountRepository {
	return &accountRepository{}
}

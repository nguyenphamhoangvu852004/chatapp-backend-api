package reporitory

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/entity"
)

type IAccountRepository interface {
	GetList() ([]entity.Account, error)
	GetUserByEmail(email string) (entity.Account, error)
	GetUserByUsername(username string) (entity.Account, error)
	GetUserByAccountId(id string) (entity.Account, error)
	Create(account entity.Account) (entity.Account, error)
	Update(account entity.Account) (entity.Account, error)
}

type accountRepository struct {
}

// GetList implements IAccountRepository.
func (a *accountRepository) GetList() ([]entity.Account, error) {
	var accounts []entity.Account
	err := global.Mdb.Preload("Profile").Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// GetUserByAccountId implements IAccountRepository.
func (a *accountRepository) GetUserByAccountId(id string) (entity.Account, error) {
	var account entity.Account
	err := global.Mdb.Preload("Profile").Where("id = ?", id).First(&account).Error

	if err != nil {
		return entity.Account{}, err
	}
	return account, nil
}

// Update implements IAccountRepository.
func (a *accountRepository) Update(account entity.Account) (entity.Account, error) {
	result := global.Mdb.Save(&account)
	if result.Error != nil {
		return entity.Account{}, result.Error
	}
	return account, nil
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

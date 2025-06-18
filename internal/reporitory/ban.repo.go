package reporitory

import "chapapp-backend-api/internal/entity"

type IBanRepository interface {
	Create(account entity.Account) (entity.Account, error)
	Delete(account entity.Account) (entity.Account, error)
	Update(account entity.Account) (entity.Account, error)
	GetList() ([]entity.Account, error)
	GetListBan() ([]entity.Account, error)
}

type banRepository struct{}

// GetListBan implements IBanRepository.
func (b *banRepository) GetListBan() ([]entity.Account, error) {
	panic("unimplemented")
}

// Create implements IBanRepository.
func (b *banRepository) Create(account entity.Account) (entity.Account, error) {
	panic("unimplemented")
}

// Delete implements IBanRepository.
func (b *banRepository) Delete(account entity.Account) (entity.Account, error) {
	panic("unimplemented")
}

// GetList implements IBanRepository.
func (b *banRepository) GetList() ([]entity.Account, error) {
	panic("unimplemented")
}

// Update implements IBanRepository.
func (b *banRepository) Update(account entity.Account) (entity.Account, error) {
	panic("unimplemented")
}

func NewBanRepository() IBanRepository {
	return &banRepository{}
}

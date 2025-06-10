package reporitory

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/entity"
)

type IProfileRepository interface {
	Create(profile entity.Profile) (entity.Profile, error)
	Update(profile entity.Profile) (entity.Profile, error)
	GetByID(id uint) (entity.Profile, error)
	GetByAccountID(id uint) (entity.Profile, error)
	DeleteByID(id uint) (entity.Profile, error)
}

type profileRepository struct {
}

// Create implements IProfileRepository.
func (p *profileRepository) Create(profile entity.Profile) (entity.Profile, error) {
	result := global.Mdb.Create(&profile)
	if result.Error != nil {
		return entity.Profile{}, result.Error
	}
	return profile, nil
}

// DeleteByID implements IProfileRepository.
func (p *profileRepository) DeleteByID(id uint) (entity.Profile, error) {
	panic("unimplemented")
}

// GetByAccountID implements IProfileRepository.
func (p *profileRepository) GetByAccountID(id uint) (entity.Profile, error) {
	panic("unimplemented")
}

// GetByID implements IProfileRepository.
func (p *profileRepository) GetByID(id uint) (entity.Profile, error) {
	var profile entity.Profile
	err := global.Mdb.First(&profile, id).Error
	if err != nil {
		return entity.Profile{}, err
	}
	return profile, nil
}

// Update implements IProfileRepository.
func (p *profileRepository) Update(profile entity.Profile) (entity.Profile, error) {

	result := global.Mdb.Save(&profile)
	if result.Error != nil {
		return entity.Profile{}, result.Error
	}
	return profile, nil
}

func NewProfileRepository() IProfileRepository {
	return &profileRepository{}
}

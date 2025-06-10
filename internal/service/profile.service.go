package service

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/reporitory"
	"strconv"
	"time"
)

type IProfileService interface {
	// Create(data profile.CreateProfileInputDTO) (profile.CreateProfileOutputDTO, error)
	Update(data dto.UpdateProfileInputDTO) (dto.UpdateProfileOutputDTO, error)
}

type profileService struct {
	profileRepo reporitory.IProfileRepository
}

func (p *profileService) Update(data dto.UpdateProfileInputDTO) (dto.UpdateProfileOutputDTO, error) {
	// Convert ProfileId from string to uint
	profileIDUint, err := strconv.ParseUint(data.ProfileId, 10, 32)
	if err != nil {
		return dto.UpdateProfileOutputDTO{}, err
	}
	entity, err := p.profileRepo.GetByID(uint(profileIDUint))
	if err != nil {
		return dto.UpdateProfileOutputDTO{}, err
	}
	entity.FullName = data.FullName
	entity.Bio = data.Bio
	entity.AvatarURL = data.AvatarURL
	entity.CoverURL = data.CoverURL
	entity.UpdatedAt = time.Now()
	updatedEntity, err := p.profileRepo.Update(entity)
	if err != nil {
		return dto.UpdateProfileOutputDTO{}, err
	}
	return dto.UpdateProfileOutputDTO{
		ProfileId: strconv.FormatUint(uint64(updatedEntity.ID), 10),
		FullName:  updatedEntity.FullName,
		Bio:       updatedEntity.Bio,
		AvatarURL: updatedEntity.AvatarURL,
		CoverURL:  updatedEntity.CoverURL,
		CreatedAt: updatedEntity.CreatedAt.Format(time.RFC3339),
		UpdatedAt: updatedEntity.UpdatedAt.Format(time.RFC3339),
	}, nil

}

func NewProfileService(profileRepo reporitory.IProfileRepository) IProfileService {
	return &profileService{profileRepo: profileRepo}
}

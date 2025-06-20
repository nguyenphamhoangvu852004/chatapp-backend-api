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
	profileIDUint, err := strconv.ParseUint(data.ProfileId, 10, 32)
	if err != nil {
		return dto.UpdateProfileOutputDTO{}, err
	}

	entity, err := p.profileRepo.GetByID(uint(profileIDUint))
	if err != nil {
		return dto.UpdateProfileOutputDTO{}, err
	}

	// Chỉ cập nhật nếu trường có giá trị khác nil
	if data.FullName != "" {
		entity.FullName = data.FullName
	}
	if data.Bio != "" {
		entity.Bio = data.Bio
	}
	if data.AvatarURL != "" {
		entity.AvatarURL = data.AvatarURL
	}
	if data.CoverURL != "" {
		entity.CoverURL = data.CoverURL
	}

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

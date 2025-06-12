package reporitory

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/entity"
)

type IParticipantRepository interface {
	Create(participant entity.Participant) (entity.Participant, error)
}

type participantRepository struct {
}

// Create implements IParticipantRepository.
func (p *participantRepository) Create(participant entity.Participant) (entity.Participant, error) {
	result := global.Mdb.Create(&participant)
	if result.Error != nil {
		return entity.Participant{}, result.Error
	}
	return participant, nil
}

func NewParticiapntRepository() IParticipantRepository {
	return &participantRepository{}
}

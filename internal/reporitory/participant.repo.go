package reporitory

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type IParticipantRepository interface {
	DeleteMany(conversationId uint, accountIds []uint) error
	AddMembers(participants []entity.Participant) error
	GetListOwnedByMe(data string) ([]entity.Participant, error)
	Create(participant entity.Participant) (entity.Participant, error)
	CheckIsAdmin(uint, uint) (bool, error)
	GetGroupListWhereUserIsAdmin(accountId string) ([]entity.Participant, error)
	FindGroupsByAccountId(accountId uint) ([]entity.Participant, error)
	FindMembersByConversationID(conversationId uint) ([]entity.Participant, error)
}

type participantRepository struct {
}

func (r *participantRepository) FindMembersByConversationID(conversationId uint) ([]entity.Participant, error) {
	var participants []entity.Participant
	err := global.Mdb.
		Preload("Account.Profile").
		Where("conversation_id = ?", conversationId).
		Find(&participants).Error
	return participants, err
}
func (r *participantRepository) FindGroupsByAccountId(accountId uint) ([]entity.Participant, error) {
	var participants []entity.Participant
	err := global.Mdb.
		Preload("Conversation").
		Where("account_id = ?", accountId).
		Find(&participants).Error

	if err != nil {
		return nil, err
	}

	// Lọc conversation là group
	var result []entity.Participant
	for _, p := range participants {
		if p.Conversation.IsGroup {
			result = append(result, p)
		}
	}
	return result, nil
}

func (r *participantRepository) DeleteMany(conversationId uint, accountIds []uint) error {
	if len(accountIds) == 0 {
		return nil
	}
	return global.Mdb.
		Where("conversation_id = ? AND account_id IN ?", conversationId, accountIds).
		Delete(&entity.Participant{}).Error
}

// AddMembers implements IParticipantRepository.
func (r *participantRepository) AddMembers(participants []entity.Participant) error {
	for _, p := range participants {
		var existing entity.Participant
		err := global.Mdb.Where("account_id = ? AND conversation_id = ?", p.AccountID, p.ConversationID).First(&existing).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := global.Mdb.Create(&p).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// CheckIsAdmin implements IParticipantRepository.
func (r *participantRepository) CheckIsAdmin(accountId, conversationId uint) (bool, error) {
	var participant entity.Participant
	err := global.Mdb.
		Where("account_id = ? AND conversation_id = ? AND role = ?", accountId, conversationId, "admin").
		First(&participant).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return err == nil, err
}

func (r *participantRepository) GetGroupListWhereUserIsAdmin(accountId string) ([]entity.Participant, error) {
	var participants []entity.Participant
	err := global.Mdb.
		Joins("JOIN conversations ON participants.conversation_id = conversations.id").
		Preload("Conversation.Participants.Account.Profile").
		Preload("Account.Profile").
		Where("participants.account_id = ? AND participants.role = ? AND conversations.is_group = ?", accountId, "admin", true).
		Find(&participants).Error

	if err != nil {
		return nil, err
	}
	return participants, nil
}

// GetListOwnedByMe implements IParticipantRepository.
func (p *participantRepository) GetListOwnedByMe(data string) ([]entity.Participant, error) {
	var participants []entity.Participant
	err := global.Mdb.Preload("Conversation").Preload("Account.Profile").Where("account_id = ? and role = ?", data, "admin").Find(&participants).Error
	if err != nil {
		return nil, err
	}
	return participants, nil
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

package reporitory

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/entity"
)

type IConversationRepository interface {
	DeleteById(id uint) error
	FindById(data uint) (entity.Conversation, error)
	GetListOwnedByMe(data string) ([]dto.GetConversationOutputDTO, error)
	FindConversationBetweenTwo(user1ID, user2ID uint) (*entity.Conversation, error)
	Create(convention entity.Conversation) (entity.Conversation, error)
}

type conversationRepository struct{}

func (r *conversationRepository) DeleteById(id uint) error {
	return global.Mdb.Unscoped().Delete(&entity.Conversation{}, id).Error
}

// FindById implements IConversationRepository.
func (r *conversationRepository) FindById(data uint) (entity.Conversation, error) {
	var conversation entity.Conversation
	err := global.Mdb.Where("id = ?", data).First(&conversation).Error
	if err != nil {
		return entity.Conversation{}, err
	}
	return conversation, nil
}

// GetListOwnedByMe implements IConversationRepository.
func (r *conversationRepository) GetListOwnedByMe(data string) ([]dto.GetConversationOutputDTO, error) {
	var conversations []dto.GetConversationOutputDTO
	err := global.Mdb.
		Joins("JOIN participants p1 ON p1.conversation_id = conversations.id").
		Where("p1.account_id = ?", data).
		Where("conversations.is_group = true").
		Where("conversations. = false").
		Find(&conversations).Error
	if err != nil {
		return nil, err
	}
	return conversations, nil
}

// FindConversationBetweenTwo implements IConversationRepository.
func (r *conversationRepository) FindConversationBetweenTwo(accountID1, accountID2 uint) (*entity.Conversation, error) {
	var conversation entity.Conversation
	err := global.Mdb.
		Joins("JOIN participants p1 ON p1.conversation_id = conversations.id").
		Joins("JOIN participants p2 ON p2.conversation_id = conversations.id").
		Where("p1.account_id = ? AND p2.account_id = ? AND conversations.is_group = false", accountID1, accountID2).
		First(&conversation).Error

	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

// Create implements IConversationRepository.
func (c *conversationRepository) Create(convention entity.Conversation) (entity.Conversation, error) {
	result := global.Mdb.Create(&convention)
	if result.Error != nil {
		return entity.Conversation{}, result.Error
	}
	return convention, nil
}

func NewConversationRepository() IConversationRepository {
	return &conversationRepository{}
}

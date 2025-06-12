package reporitory

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/entity"
)

type IConversationRepository interface {
	FindConversationBetweenTwo(user1ID, user2ID uint) (*entity.Conversation, error)
	Create(convention entity.Conversation) (entity.Conversation, error)
}

type conversationRepository struct{}

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

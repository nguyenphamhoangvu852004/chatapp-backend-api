package reporitory

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/entity"
)

type IMessageRepository interface {
	GetMessagesByConversation(conversationId string) ([]entity.Message, error)
	Create(message entity.Message) (entity.Message, error)
}

type messageRepository struct {
}

// GetMessagesByConversation implements IMessageRepository.
func (m *messageRepository) GetMessagesByConversation(conversationId string) ([]entity.Message, error) {
var messages []entity.Message

	err := global.Mdb.
		Where("conversation_id = ?", conversationId).
		Order("created_at ASC").
		Find(&messages).Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}

// Create implements IMessageRepository.
func (m *messageRepository) Create(message entity.Message) (entity.Message, error) {
	result := global.Mdb.Create(&message)
	if result.Error != nil {
		return entity.Message{}, result.Error
	}
	return message, nil
}

func NewMessageRepository() IMessageRepository {
	return &messageRepository{}
}

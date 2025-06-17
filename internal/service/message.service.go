package service

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/entity"
	"chapapp-backend-api/internal/reporitory"
	"fmt"
	"strconv"
	"time"
)

type IMessageService interface {
	GetList(data dto.GetListMessageInputDTO) (dto.GetListMessageOutputDTO, error)
	Create(data dto.CreateMessageInputDTO) (dto.CreateMessageOutputDTO, error)
}

type messageService struct {
	messageRepo reporitory.IMessageRepository
}

// Create implements IMessageService.
func (s *messageService) Create(data dto.CreateMessageInputDTO) (dto.CreateMessageOutputDTO, error) {
	var conversationID uint
	if id, err := strconv.ParseUint(data.ConversationId, 10, 64); err == nil {
		conversationID = uint(id)
	} else {
		return dto.CreateMessageOutputDTO{}, err
	}
	var senderID uint
	if id, err := strconv.ParseUint(data.SenderId, 10, 64); err == nil {
		senderID = uint(id)
	} else {
		return dto.CreateMessageOutputDTO{}, err
	}
	var messageEntity = entity.Message{
		SenderID:       senderID,
		ConversationID: conversationID,
		Content:        data.Content,
		OriginFilename: *data.OriginFilename,
		MessageType:    "text",
	}

	createdMessage, err := s.messageRepo.Create(messageEntity)
	if err != nil {
		return dto.CreateMessageOutputDTO{}, err
	}

	return dto.CreateMessageOutputDTO{
		MessageId:      strconv.FormatUint(uint64(createdMessage.ID), 10),
		Content:        createdMessage.Content,
		OriginFilename: &createdMessage.OriginFilename,
	}, nil

}

func (s *messageService) GetList(data dto.GetListMessageInputDTO) (dto.GetListMessageOutputDTO, error) {
	messages, err := s.messageRepo.GetMessagesByConversation(data.ConversationId)
	if err != nil {
		return dto.GetListMessageOutputDTO{}, err
	}

	var result []dto.Message
	for _, m := range messages {
		result = append(result, dto.Message{
			ID:        fmt.Sprintf("%d", m.ID),
			SenderId:  fmt.Sprintf("%d", m.SenderID),
			Content:   m.Content,
			Type:      m.MessageType,
			CreatedAt: m.CreatedAt.Format(time.RFC3339),
		})
	}

	return dto.GetListMessageOutputDTO{
		ConversationId: data.ConversationId,
		Messages:       result,
	}, nil
}

func NewMessageService(messageRepo reporitory.IMessageRepository) IMessageService {
	return &messageService{messageRepo: messageRepo}
}

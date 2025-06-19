package service

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/entity"
	exception "chapapp-backend-api/internal/exeption"
	"chapapp-backend-api/internal/reporitory"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type IMessageService interface {
	GetList(data dto.GetListMessageInputDTO) (dto.GetListMessageOutputDTO, error)
	Create(data dto.CreateMessageInputDTO) (dto.CreateMessageOutputDTO, error)
	Delete(data dto.DeleteMessageInputDTO) (dto.DeleteMessageOutputDTO, error)
}

type messageService struct {
	messageRepo      reporitory.IMessageRepository
	accountRepo      reporitory.IAccountRepository
	conversationRepo reporitory.IConversationRepository
}

// Delete implements IMessageService.
func (s *messageService) Delete(data dto.DeleteMessageInputDTO) (dto.DeleteMessageOutputDTO, error) {
	// tìm cái account
	owner, err := s.accountRepo.GetUserByAccountId(data.OwnerId)
	if err != nil {
		return dto.DeleteMessageOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "account not found")
	}

	// tim cai conversation

	conversationID, err := strconv.ParseUint(data.ConversationId, 10, 64)
	if err != nil {
		return dto.DeleteMessageOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "invalid conversation id")
	}
	conversation, err := s.conversationRepo.FindById(uint(conversationID))
	if err != nil {
		return dto.DeleteMessageOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "conversation not found")
	}

	// tim cai message

	message, err := s.messageRepo.FindById(data.MessageId)
	if err != nil {
		return dto.DeleteMessageOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "message not found")
	}

	// kiem tra coi no co bi xoa hay chua

	if message.DeletedAt.Valid {
		return dto.DeleteMessageOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "message has been deleted")
	}

	// kiem tra coi co phai message do cua nguoi do hay khong

	if message.SenderID != owner.ID {
		return dto.DeleteMessageOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "you are not the owner of this message")
	}

	// xoa no di

	res, err := s.messageRepo.Delete(message)
	if err != nil {
		return dto.DeleteMessageOutputDTO{}, err
	}

	return dto.DeleteMessageOutputDTO{
		MessageId:      strconv.FormatUint(uint64(res.ID), 10),
		OwnerId:        strconv.FormatUint(uint64(owner.ID), 10),
		ConversationId: strconv.FormatUint(uint64(conversation.ID), 10),
		IsDeleted:      true,
	}, nil

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
		Size:           *data.Size,
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
		Size:           &createdMessage.Size,
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
			ID:             fmt.Sprintf("%d", m.ID),
			SenderId:       fmt.Sprintf("%d", m.SenderID),
			Content:        m.Content,
			Type:           m.MessageType,
			OriginFilename: &m.OriginFilename,
			Size:           &m.Size,
			IsDeleted:      m.DeletedAt.Valid,
			CreatedAt:      m.CreatedAt.Format(time.RFC3339),
		})
	}

	return dto.GetListMessageOutputDTO{
		ConversationId: data.ConversationId,
		Messages:       result,
	}, nil
}

func NewMessageService(messageRepo reporitory.IMessageRepository, accountRepo reporitory.IAccountRepository, conversationRepo reporitory.IConversationRepository) IMessageService {
	return &messageService{messageRepo: messageRepo, accountRepo: accountRepo, conversationRepo: conversationRepo}

}

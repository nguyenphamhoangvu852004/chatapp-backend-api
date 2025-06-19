package service

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/entity"
	"chapapp-backend-api/internal/service"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 🧪 Mock IMessageRepository
type MockMessageRepository struct {
	mock.Mock
}

func (m *MockMessageRepository) Create(msg entity.Message) (entity.Message, error) {
	args := m.Called(msg)
	return args.Get(0).(entity.Message), args.Error(1)
}

func (m *MockMessageRepository) GetMessagesByConversation(conversationID string) ([]entity.Message, error) {
	args := m.Called(conversationID)
	return args.Get(0).([]entity.Message), args.Error(1)
}

func TestCreateMessage_Success(t *testing.T) {
	mockRepo := new(MockMessageRepository)
	svc := service.NewMessageService(mockRepo)

	originFile := "test.txt"
	size := "1234"

	input := dto.CreateMessageInputDTO{
		SenderId:       "1",
		ConversationId: "2",
		Content:        "Hello",
		OriginFilename: &originFile,
		Size:           &size,
	}

	expectedEntity := entity.Message{
		BaseEntity: entity.BaseEntity{
			ID:        10,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		SenderID:       1,
		ConversationID: 2,
		Content:        "Hello",
		OriginFilename: "test.txt",
		Size:           "1234",
		MessageType:    "text",
	}

	mockRepo.On("Create", mock.MatchedBy(func(msg entity.Message) bool {
		return msg.SenderID == 1 && msg.ConversationID == 2
	})).Return(entity.Message{
		BaseEntity: entity.BaseEntity{
			ID:        10,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Content:        "Hello",
		OriginFilename: "test.txt",
		Size:           "1234",
	}, nil)
	output, err := svc.Create(input)
	assert.NoError(t, err)
	assert.Equal(t, "10", output.MessageId)
	assert.Equal(t, "Hello", output.Content)
	assert.Equal(t, "test.txt", *output.OriginFilename)
assert.Equal(t, expectedEntity.Content, output.Content)
// hoặc
assert.Equal(t, expectedEntity.OriginFilename, *output.OriginFilename)

	mockRepo.AssertExpectations(t)
}

func TestCreateMessage_InvalidID(t *testing.T) {
	mockRepo := new(MockMessageRepository)
	svc := service.NewMessageService(mockRepo)

	input := dto.CreateMessageInputDTO{
		SenderId:       "abc", // invalid ID
		ConversationId: "2",
		Content:        "Hello",
		OriginFilename: nil,
		Size:           nil,
	}

	_, err := svc.Create(input)
	assert.Error(t, err)
}

func TestGetList_Success(t *testing.T) {
	mockRepo := new(MockMessageRepository)
	svc := service.NewMessageService(mockRepo)

	conversationID := "2"
	fakeTime := time.Now()

	mockRepo.On("GetMessagesByConversation", conversationID).Return([]entity.Message{
		{
			SenderID:       1,
			Content:        "Hello",
			MessageType:    "text",
			OriginFilename: "file.jpg",
			Size:           "999",
			BaseEntity: entity.BaseEntity{
				CreatedAt: fakeTime,
				UpdatedAt: fakeTime,
			},
		},
	}, nil)

	input := dto.GetListMessageInputDTO{
		ConversationId: conversationID,
	}
	output, err := svc.GetList(input)

	assert.NoError(t, err)
	assert.Equal(t, conversationID, output.ConversationId)
	assert.Len(t, output.Messages, 1)
	assert.Equal(t, "Hello", output.Messages[0].Content)
	assert.Equal(t, "file.jpg", *output.Messages[0].OriginFilename)
	assert.Equal(t, "1", output.Messages[0].SenderId)

	mockRepo.AssertExpectations(t)
}

func TestGetList_Error(t *testing.T) {
	mockRepo := new(MockMessageRepository)
	svc := service.NewMessageService(mockRepo)

mockRepo.On("GetMessagesByConversation", "2").Return([]entity.Message(nil), errors.New("db error"))


	_, err := svc.GetList(dto.GetListMessageInputDTO{ConversationId: "2"})
	assert.Error(t, err)
}

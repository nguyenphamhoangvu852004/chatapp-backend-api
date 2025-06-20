package service

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/entity"
	"chapapp-backend-api/internal/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTestService() (*MockMessageRepository, *MockAccountRepo, *MockConversationRepo, service.IMessageService) {
	mockMsgRepo := new(MockMessageRepository)
	mockAccRepo := new(MockAccountRepo)
	mockConvRepo := new(MockConversationRepo)

	svc := service.NewMessageService(mockMsgRepo, mockAccRepo, mockConvRepo)
	return mockMsgRepo, mockAccRepo, mockConvRepo, svc
}

type MockMessageRepository struct {
	mock.Mock
}

// Delete implements reporitory.IMessageRepository.
func (m *MockMessageRepository) Delete(message entity.Message) (entity.Message, error) {
	args := m.Called(message)
	return args.Get(0).(entity.Message), args.Error(1)
}

// FindById implements reporitory.IMessageRepository.
func (m *MockMessageRepository) FindById(messageId string) (entity.Message, error) {
	args := m.Called(messageId)
	return args.Get(0).(entity.Message), args.Error(1)
}

func (m *MockMessageRepository) Create(msg entity.Message) (entity.Message, error) {
	args := m.Called(msg)
	return args.Get(0).(entity.Message), args.Error(1)
}

func (m *MockMessageRepository) GetMessagesByConversation(conversationID string) ([]entity.Message, error) {
	args := m.Called(conversationID)
	return args.Get(0).([]entity.Message), args.Error(1)
}

type MockConversationRepo struct {
	mock.Mock
}

// Create implements reporitory.IConversationRepository.
func (m *MockConversationRepo) Create(conversation entity.Conversation) (entity.Conversation, error) {
	args := m.Called(conversation)
	return args.Get(0).(entity.Conversation), args.Error(1)
}

// DeleteById implements reporitory.IConversationRepository.
func (m *MockConversationRepo) DeleteById(id uint) error {
	panic("unimplemented")
}

// FindById implements reporitory.IConversationRepository.

func (m *MockConversationRepo) FindById(data uint) (entity.Conversation, error) {
	args := m.Called(data)
	return args.Get(0).(entity.Conversation), args.Error(1)
}

// FindConversationBetweenTwo implements reporitory.IConversationRepository.
func (m *MockConversationRepo) FindConversationBetweenTwo(user1ID uint, user2ID uint) (*entity.Conversation, error) {
	panic("unimplemented")
}

// GetListOwnedByMe implements reporitory.IConversationRepository.
func (m *MockConversationRepo) GetListOwnedByMe(data string) ([]dto.GetConversationOutputDTO, error) {
	panic("unimplemented")
}

// GetMembersByConversationID implements reporitory.IConversationRepository.
func (m *MockConversationRepo) GetMembersByConversationID(conversationID uint) ([]entity.Participant, error) {
	panic("unimplemented")
}

// Update implements reporitory.IConversationRepository.
func (m *MockConversationRepo) Update(convention entity.Conversation) (entity.Conversation, error) {
	panic("unimplemented")
}

func TestCreateMessage_Success(t *testing.T) {
	msgRepo, _, _, svc := setupTestService()

	originFile := "test.txt"
	size := "1234"

	input := dto.CreateMessageInputDTO{
		SenderId:       "1",
		ConversationId: "2",
		Content:        "Hello",
		OriginFilename: &originFile,
		Size:           &size,
	}

	expected := entity.Message{
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

	msgRepo.On("Create", mock.MatchedBy(func(msg entity.Message) bool {
		return msg.SenderID == 1 && msg.ConversationID == 2
	})).Return(expected, nil)

	output, err := svc.Create(input)

	assert.NoError(t, err)
	assert.Equal(t, "10", output.MessageId)
	assert.Equal(t, "Hello", output.Content)
	assert.Equal(t, "test.txt", *output.OriginFilename)

	msgRepo.AssertExpectations(t)
}

func TestCreateMessage_InvalidID(t *testing.T) {
	_, _, _, svc := setupTestService()

	input := dto.CreateMessageInputDTO{
		SenderId:       "abc",
		ConversationId: "2",
		Content:        "Hello",
	}

	_, err := svc.Create(input)
	assert.Error(t, err)
}

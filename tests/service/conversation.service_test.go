package service

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/entity"
	"chapapp-backend-api/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateConversation_Success(t *testing.T) {
	// Arrange
	conversationRepo := new(MockConversationRepo)
	participantRepo := new(MockParticipantRepo)

	service := service.NewConversationService(conversationRepo, participantRepo)

	input := dto.CreateConversationInputDTO{
		Name:      "Test Group",
		OwnerId:   "1",
		AvatarURL: "avatar.png",
	}

	mockConversation := entity.Conversation{
		BaseEntity:  entity.BaseEntity{ID: 1},
		Name:        input.Name,
		IsGroup:     true,
		GroupAvatar: input.AvatarURL,
	}
	mockParticipant := entity.Participant{
		AccountID:      1,
		ConversationID: 1,
		Role:           "admin",
	}

	conversationRepo.On("Create", mock.AnythingOfType("entity.Conversation")).Return(mockConversation, nil)
	participantRepo.On("Create", mock.Anything).Return(mockParticipant, nil)

	// Act
	result, err := service.Create(input)

	// Assert
	assert.NoError(t, err)
	assert.True(t, result.IsGroup)
	assert.Equal(t, "1", result.ConversationId)
	assert.Equal(t, "1", result.OwnerId)
	assert.Equal(t, "Test Group", result.Name)
}

func TestAddMembers_Success(t *testing.T) {
	mockConversationRepo := new(MockConversationRepo)
	mockParticipantRepo := new(MockParticipantRepo)
	service := service.NewConversationService(mockConversationRepo, mockParticipantRepo)

	// Fake input
	input := dto.AddMemberInputDTO{
		ConversationId: "1",
		OwnerId:        "10",
		MemberIds:      []string{"2", "3"},
	}

	conversationID := uint(1)
	ownerID := uint(10)

	// Mock: FindById trả về 1 conversation tồn tại
	mockConversationRepo.On("FindById", conversationID).Return(entity.Conversation{
		BaseEntity: entity.BaseEntity{ID: conversationID},
		Name:       "Test Group",
		IsGroup:    true,
	}, nil)

	// Mock: CheckIsAdmin trả về true
	mockParticipantRepo.On("CheckIsAdmin", ownerID, conversationID).Return(true, nil)

	// Mock: AddMembers không trả lỗi
	mockParticipantRepo.On("AddMembers", mock.Anything).Return(nil)

	// Call service
	output, err := service.AddMembers(input)

	// Assert
	assert.NoError(t, err)
	assert.True(t, output.IsSuccess)
	assert.Equal(t, input.ConversationId, output.ConversationId)
	assert.Equal(t, input.OwnerId, output.OwnerId)
	assert.Equal(t, input.MemberIds, output.MemberIds)

	// Verify mock
	mockConversationRepo.AssertExpectations(t)
	mockParticipantRepo.AssertExpectations(t)
}

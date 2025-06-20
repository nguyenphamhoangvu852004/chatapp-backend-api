package service

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/entity"
	"chapapp-backend-api/internal/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTestFriendShipService() (
	*MockFriendShipRepo,
	*MockAccountRepo,
	*MockProfileRepo,
	*MockParticipantRepo,
	*MockConversationRepo,
	*MockBlockRepo,
	service.IFriendShipService,
) {
	friendRepo := new(MockFriendShipRepo)
	accRepo := new(MockAccountRepo)
	profileRepo := new(MockProfileRepo)
	participantRepo := new(MockParticipantRepo)
	convRepo := new(MockConversationRepo)
	blockRepo := new(MockBlockRepo)

	svc := service.NewFriendShipService(friendRepo, accRepo, profileRepo, participantRepo, convRepo, blockRepo)

	return friendRepo, accRepo, profileRepo, participantRepo, convRepo, blockRepo, svc
}

type MockParticipantRepo struct {
	mock.Mock
}

// AddMembers implements reporitory.IParticipantRepository.
func (m *MockParticipantRepo) AddMembers(participants []entity.Participant) error {
	args := m.Called(participants)
	return args.Error(0)
}

// CheckIsAdmin implements reporitory.IParticipantRepository.
func (m *MockParticipantRepo) CheckIsAdmin(accountID uint, conversationID uint) (bool, error) {
	args := m.Called(accountID, conversationID)
	return args.Bool(0), args.Error(1)
}

// Create implements reporitory.IParticipantRepository.
func (m *MockParticipantRepo) Create(participant entity.Participant) (entity.Participant, error) {
	args := m.Called(participant)
	return args.Get(0).(entity.Participant), args.Error(1)
}

// DeleteMany implements reporitory.IParticipantRepository.
func (m *MockParticipantRepo) DeleteMany(conversationId uint, accountIds []uint) error {
	panic("unimplemented")
}

// FindGroupsByAccountId implements reporitory.IParticipantRepository.
func (m *MockParticipantRepo) FindGroupsByAccountId(accountId uint) ([]entity.Participant, error) {
	panic("unimplemented")
}

// FindMembersByConversationID implements reporitory.IParticipantRepository.
func (m *MockParticipantRepo) FindMembersByConversationID(conversationId uint) ([]entity.Participant, error) {
	panic("unimplemented")
}

// GetGroupListWhereUserIsAdmin implements reporitory.IParticipantRepository.
func (m *MockParticipantRepo) GetGroupListWhereUserIsAdmin(accountId string) ([]entity.Participant, error) {
	panic("unimplemented")
}

// GetListOwnedByMe implements reporitory.IParticipantRepository.
func (m *MockParticipantRepo) GetListOwnedByMe(data string) ([]entity.Participant, error) {
	panic("unimplemented")
}

type MockProfileRepo struct {
	mock.Mock
}

// Create implements reporitory.IProfileRepository.
func (m *MockProfileRepo) Create(profile entity.Profile) (entity.Profile, error) {
	panic("unimplemented")
}

// DeleteByID implements reporitory.IProfileRepository.
func (m *MockProfileRepo) DeleteByID(id uint) (entity.Profile, error) {
	panic("unimplemented")
}

// GetByAccountID implements reporitory.IProfileRepository.
func (m *MockProfileRepo) GetByAccountID(id uint) (entity.Profile, error) {
	panic("unimplemented")
}

// GetByID implements reporitory.IProfileRepository.
func (m *MockProfileRepo) GetByID(id uint) (entity.Profile, error) {
	panic("unimplemented")
}

// Update implements reporitory.IProfileRepository.
func (m *MockProfileRepo) Update(profile entity.Profile) (entity.Profile, error) {
	panic("unimplemented")
}

type MockBlockRepo struct {
	mock.Mock
}

// CreateBlock implements reporitory.IBlockRepository.
func (m *MockBlockRepo) CreateBlock(block entity.Block) (entity.Block, error) {
	panic("unimplemented")
}

// DeleteBlock implements reporitory.IBlockRepository.
func (m *MockBlockRepo) DeleteBlock(block entity.Block) (entity.Block, error) {
	panic("unimplemented")
}

// GetListBlocked implements reporitory.IBlockRepository.
func (m *MockBlockRepo) GetListBlocked(accountID uint) ([]entity.Block, error) {
	panic("unimplemented")
}

// GetListBlocker implements reporitory.IBlockRepository.
func (m *MockBlockRepo) GetListBlocker(accountID uint) ([]entity.Block, error) {
	panic("unimplemented")
}

// IsBlocked implements reporitory.IBlockRepository.
func (m *MockBlockRepo) IsBlocked(blockerId uint, blockedId uint) (bool, error) {
	panic("unimplemented")
}

type MockFriendShipRepo struct {
	mock.Mock
}

// Create implements reporitory.IFriendShipRepository.
func (m *MockFriendShipRepo) Create(friendShip entity.FriendShip) (entity.FriendShip, error) {
	args := m.Called(friendShip)
	return args.Get(0).(entity.FriendShip), args.Error(1)
}

// DeleteByID implements reporitory.IFriendShipRepository.
func (m *MockFriendShipRepo) DeleteByID(id uint) (entity.FriendShip, error) {
	panic("unimplemented")
}

// FindAllFriendOfAccount implements reporitory.IFriendShipRepository.
func (m *MockFriendShipRepo) FindAllFriendOfAccount(id string) ([]entity.Account, error) {
	panic("unimplemented")
}

// FindAllReceivedFriendRequests implements reporitory.IFriendShipRepository.
func (m *MockFriendShipRepo) FindAllReceivedFriendRequests(accountID string) ([]entity.Account, error) {
	panic("unimplemented")
}

// FindAllSendFriendShips implements reporitory.IFriendShipRepository.
func (m *MockFriendShipRepo) FindAllSendFriendShips(id string) ([]entity.FriendShip, error) {
	panic("unimplemented")
}

// FindBySenderAndReceiver implements reporitory.IFriendShipRepository.
func (m *MockFriendShipRepo) FindBySenderAndReceiver(senderId uint, receiverId uint) (entity.FriendShip, error) {
	args := m.Called(senderId, receiverId)
	return args.Get(0).(entity.FriendShip), args.Error(1)
}

// GetByAccountID implements reporitory.IFriendShipRepository.
func (m *MockFriendShipRepo) GetByAccountID(id uint) (entity.FriendShip, error) {
	panic("unimplemented")
}

// GetByID implements reporitory.IFriendShipRepository.
func (m *MockFriendShipRepo) GetByID(id uint) (entity.FriendShip, error) {
	panic("unimplemented")
}

// GetList implements reporitory.IFriendShipRepository.
func (m *MockFriendShipRepo) GetList(data dto.GetListFriendShipInputDTO) ([]entity.FriendShip, error) {
	panic("unimplemented")
}

// Update implements reporitory.IFriendShipRepository.
func (m *MockFriendShipRepo) Update(fs entity.FriendShip) (entity.FriendShip, error) {
	args := m.Called(fs)
	return args.Get(0).(entity.FriendShip), args.Error(1)
}

func TestCreateFriendShip_Success(t *testing.T) {
	friendRepo, accRepo, _, _, _, _, svc := setupTestFriendShipService()

	input := dto.CreateFriendShipInputDTO{
		SenderID:   "1",
		ReceiverID: "2",
	}

	// Mock sender và receiver account
	sender := entity.Account{
		BaseEntity: entity.BaseEntity{ID: 1},
		Username:   "sender",
	}
	receiver := entity.Account{
		BaseEntity: entity.BaseEntity{ID: 2},
		Username:   "receiver",
	}

	// Mock không tồn tại quan hệ trước đó
	friendRepo.On("FindBySenderAndReceiver", uint(1), uint(2)).
		Return(entity.FriendShip{}, errors.New("not found"))

	// Mock accountRepo trả về sender và receiver
	accRepo.On("GetUserByAccountId", "1").Return(sender, nil)

	accRepo.On("GetUserByAccountId", "2").Return(receiver, nil)

	// Mock tạo mới friendShip
	friendRepo.On("Create", mock.MatchedBy(func(fs entity.FriendShip) bool {
		return fs.SenderID == 1 && fs.ReceiverID == 2 && fs.Status == entity.PENDING
	})).Return(entity.FriendShip{
		SenderID:   1,
		ReceiverID: 2,
		Status:     entity.PENDING,
	}, nil)

	// Gọi hàm tạo
	output, err := svc.Create(input)

	// Kiểm tra kết quả
	assert.NoError(t, err)
	assert.Equal(t, "1", output.SenderID)
	assert.Equal(t, "2", output.ReceiverID)
	assert.Equal(t, entity.PENDING, output.Status)

	// Đảm bảo mock được gọi đúng
	friendRepo.AssertExpectations(t)
	accRepo.AssertExpectations(t)
}

func TestDeleteFriendShip_NotFound(t *testing.T) {
	friendRepo, accRepo, _, _, _, _, svc := setupTestFriendShipService()

	// Input DTO
	input := dto.DeleteFriendShipInputDTO{
		SenderID:   1,
		ReceiverID: 2,
	}

	// Mock accountRepo trả về sender và receiver hợp lệ
	accRepo.On("GetUserByAccountId", "1").Return(entity.Account{
		BaseEntity: entity.BaseEntity{ID: 1},
		Username:   "sender",
	}, nil)
	accRepo.On("GetUserByAccountId", "2").Return(entity.Account{
		BaseEntity: entity.BaseEntity{ID: 2},
		Username:   "receiver",
	}, nil)

	// Mock không tìm thấy mối quan hệ theo cả 2 chiều
	friendRepo.On("FindBySenderAndReceiver", uint(1), uint(2)).
		Return(entity.FriendShip{}, errors.New("not found"))
	friendRepo.On("FindBySenderAndReceiver", uint(2), uint(1)).
		Return(entity.FriendShip{}, errors.New("not found"))

	// Gọi hàm
	_, err := svc.Delete(input)

	// Kỳ vọng lỗi "friendship not found"
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "friendship not found")

	// Đảm bảo mock được gọi
	friendRepo.AssertExpectations(t)
	accRepo.AssertExpectations(t)
}

func TestUpdateFriendShip_ToAccepted_CreatesConversationAndParticipants(t *testing.T) {
	// Setup mock & service
	friendRepo, _, _, partRepo, convRepo, _, svc := setupTestFriendShipService()

	input := dto.UpdateFriendShipInputDTO{
		SenderID:   "1",
		ReceiverID: "2",
		Status:     "ACCEPTED",
	}

	// Mocks for conversion
	senderID := uint(1)
	receiverID := uint(2)

	// Mock FindBySenderAndReceiver → trả về friendship tồn tại
	friendEntity := entity.FriendShip{
		BaseEntity: entity.BaseEntity{ID: 99},
		SenderID:   senderID,
		ReceiverID: receiverID,
		Status:     entity.PENDING,
	}
	friendRepo.On("FindBySenderAndReceiver", senderID, receiverID).
		Return(friendEntity, nil)

	// Mock Update → sau khi đổi trạng thái thành ACCEPTED
	friendEntity.Status = entity.ACCEPTED
	friendRepo.On("Update", mock.MatchedBy(func(f entity.FriendShip) bool {
		return f.Status == entity.ACCEPTED
	})).Return(friendEntity, nil)

	// Mock ConversationRepo.Create → tạo mới conversation
	convRepo.On("Create", mock.MatchedBy(func(c entity.Conversation) bool {
		return !c.IsGroup
	})).Return(entity.Conversation{
		BaseEntity: entity.BaseEntity{ID: 123},
		IsGroup:    false,
	}, nil)

	// Mock ParticipantRepo.Create → tạo 2 participants
	partRepo.On("Create", mock.Anything).Return(entity.Participant{}, nil).Twice()

	// Call service
	output, err := svc.Update(input)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "1", output.SenderID)
	assert.Equal(t, "2", output.ReceiverID)

	friendRepo.AssertExpectations(t)
	convRepo.AssertExpectations(t)
	partRepo.AssertExpectations(t)
}

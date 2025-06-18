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

type IFriendShipService interface {
	Delete(data dto.DeleteFriendShipInputDTO) (dto.DeleteFriendShipOutputDTO, error)
	Create(data dto.CreateFriendShipInputDTO) (dto.CreateFriendShipOutputDTO, error)
	Update(data dto.UpdateFriendShipInputDTO) (dto.UpdateFriendShipOutputDTO, error)
	GetList(data dto.GetListFriendShipInputDTO) (dto.GetListFriendShipOutputDTO, error)
	GetListSendFriendShips(id string) (dto.GetFriendShipOutputDTO, error)
	GetListFriendShipsOfAccount(id string) (dto.GetFriendShipOutputDTO, error)
	GetListReceiveFriendShips(id string) (dto.GetFriendShipOutputDTO, error)
}

type friendShipService struct {
	friendShipRepo   reporitory.IFriendShipRepository
	accountRepo      reporitory.IAccountRepository
	profileRepo      reporitory.IProfileRepository
	conversationRepo reporitory.IConversationRepository
	participantRepo  reporitory.IParticipantRepository
	blockRepo        reporitory.IBlockRepository
}

// Delete implements IFriendShipService.
func (s *friendShipService) Delete(data dto.DeleteFriendShipInputDTO) (dto.DeleteFriendShipOutputDTO, error) {
	// 1. Lấy thông tin sender và receiver
	sender, err := s.accountRepo.GetUserByAccountId(strconv.FormatInt(data.SenderID, 10))
	if err != nil {
		return dto.DeleteFriendShipOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "sender not found")
	}

	receiver, err := s.accountRepo.GetUserByAccountId(strconv.FormatInt(data.ReceiverID, 10))
	if err != nil {
		return dto.DeleteFriendShipOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "receiver not found")
	}

	// 2. Tìm mối quan hệ theo cả hai chiều
	friendShip, err := s.friendShipRepo.FindBySenderAndReceiver(sender.ID, receiver.ID)
	if err != nil || friendShip.ID == 0 {
		friendShip, err = s.friendShipRepo.FindBySenderAndReceiver(receiver.ID, sender.ID)
		if err != nil || friendShip.ID == 0 {
			return dto.DeleteFriendShipOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "friendship not found")
		}
	}

	// 3. Xóa quan hệ bạn bè
	friendShipDeleted, err := s.friendShipRepo.DeleteByID(friendShip.ID)
	if err != nil {
		return dto.DeleteFriendShipOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "failed to delete friendship")
	}

	// Xoá hết trong cái bảng conversation luôn

	// 4. Trả response
	return dto.DeleteFriendShipOutputDTO{
		SenderID:   int64(friendShipDeleted.SenderID),
		ReceiverID: int64(friendShipDeleted.ReceiverID),
		IsSuccess:  true,
	}, nil
}

// GetListReceiveFriendShips implements IFriendShipService.
func (s *friendShipService) GetListReceiveFriendShips(id string) (dto.GetFriendShipOutputDTO, error) {
	// Tìm thông tin người dùng
	account, err := s.accountRepo.GetUserByAccountId(id)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "account not found")
	}

	// Convert id sang uint
	var myID uint
	fmt.Sscanf(id, "%d", &myID)

	// Lấy danh sách người đã block mình
	blockedList, err := s.blockRepo.GetListBlocked(myID)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, err
	}
	// Lấy danh sách người mình đã block
	blockerList, err := s.blockRepo.GetListBlocker(myID)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, err
	}

	// Tập hợp các ID cần loại ra
	blockedIDs := make(map[uint]bool)
	for _, b := range blockedList {
		blockedIDs[b.BlockerID] = true
	}
	for _, b := range blockerList {
		blockedIDs[b.BlockedID] = true
	}

	// Lấy danh sách sender (gửi lời mời)
	senderAccounts, err := s.friendShipRepo.FindAllReceivedFriendRequests(id)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, err
	}

	var receivers []dto.Receiver
	for _, sender := range senderAccounts {
		if blockedIDs[sender.ID] {
			continue
		}

		receiver := dto.Receiver{
			ID:       sender.ID,
			Username: sender.Username,
			Email:    sender.Email,
			ImageURL: "",
			Status:   "PENDING",
		}
		if sender.Profile != nil {
			receiver.ImageURL = sender.Profile.AvatarURL
		}
		receivers = append(receivers, receiver)
	}

	output := &dto.GetFriendShipOutputDTO{
		Me: dto.Sender{
			ID:       account.ID,
			Username: account.Username,
			Email:    account.Email,
		},
		Others: receivers,
	}

	return *output, nil
}

// GetListFriendShipsOfAccount implements IFriendShipService.
func (s *friendShipService) GetListFriendShipsOfAccount(id string) (dto.GetFriendShipOutputDTO, error) {
	// Tìm account
	account, err := s.accountRepo.GetUserByAccountId(id)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "account not found")
	}

	// Convert id sang uint
	var myID uint
	fmt.Sscanf(id, "%d", &myID)

	// Lấy danh sách người đã block mình
	blockedList, err := s.blockRepo.GetListBlocked(myID)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, err
	}
	// Lấy danh sách người mình đã block
	blockerList, err := s.blockRepo.GetListBlocker(myID)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, err
	}

	// Tổng hợp ID cần loại bỏ
	blockedIDs := make(map[uint]bool)
	for _, b := range blockedList {
		blockedIDs[b.BlockerID] = true
	}
	for _, b := range blockerList {
		blockedIDs[b.BlockedID] = true
	}

	// Tìm bạn bè
	friendAccounts, err := s.friendShipRepo.FindAllFriendOfAccount(id)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, err
	}

	// Tạo danh sách receivers (chỉ bạn bè không bị block hoặc block mình)
	var receivers []dto.Receiver
	for _, acc := range friendAccounts {
		if blockedIDs[acc.ID] {
			continue
		}

		conversation, _ := s.conversationRepo.FindConversationBetweenTwo(account.ID, acc.ID)
		receiver := dto.Receiver{
			ID:             acc.ID,
			Username:       acc.Username,
			Email:          acc.Email,
			ImageURL:       "",
			Status:         "ACCEPTED",
			ConversationID: conversation.ID,
		}
		if acc.Profile != nil {
			receiver.ImageURL = acc.Profile.AvatarURL
		}
		receivers = append(receivers, receiver)
	}

	output := &dto.GetFriendShipOutputDTO{
		Me: dto.Sender{
			ID:       account.ID,
			Username: account.Username,
			Email:    account.Email,
		},
		Others: receivers,
	}

	return *output, nil
}

// GetListSendFriendShips implements IFriendShipService.
func (f *friendShipService) GetListSendFriendShips(id string) (dto.GetFriendShipOutputDTO, error) {
	// Tìm thông tin tài khoản
	account, err := f.accountRepo.GetUserByAccountId(id)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "Not found account")
	}

	// Convert id sang uint
	var myID uint
	fmt.Sscanf(id, "%d", &myID)

	// Lấy danh sách block (2 chiều)
	blockedList, err := f.blockRepo.GetListBlocked(myID)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, err
	}
	blockerList, err := f.blockRepo.GetListBlocker(myID)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, err
	}

	// Tổng hợp các ID bị block
	blockedIDs := make(map[uint]bool)
	for _, b := range blockedList {
		blockedIDs[b.BlockerID] = true
	}
	for _, b := range blockerList {
		blockedIDs[b.BlockedID] = true
	}

	// Lấy danh sách yêu cầu đã gửi
	list, err := f.friendShipRepo.FindAllSendFriendShips(id)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to get list send friendship")
	}

	var listReceiver []dto.Receiver
	for _, fs := range list {
		// Nếu người nhận nằm trong danh sách bị block thì bỏ qua
		if blockedIDs[fs.ReceiverID] {
			continue
		}

		listReceiver = append(listReceiver, dto.Receiver{
			ID:       fs.ReceiverID,
			Username: fs.Receiver.Username,
			Email:    fs.Receiver.Email,
			ImageURL: fs.Receiver.Profile.AvatarURL,
			Status:   fs.Status,
		})
	}

	var outputDTO = dto.GetFriendShipOutputDTO{
		Me: dto.Sender{
			ID:       account.ID,
			Username: account.Username,
			Email:    account.Email,
		},
		Others: listReceiver,
	}

	return outputDTO, nil
}

// GetList implements IFriendShipService.
func (f *friendShipService) GetList(data dto.GetListFriendShipInputDTO) (dto.GetListFriendShipOutputDTO, error) {

	list, err := f.friendShipRepo.GetList(data)
	if err != nil {
		return dto.GetListFriendShipOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to get list friendship")
	}
	var outDTO dto.GetListFriendShipOutputDTO
	var listFriendShipItemDTO []dto.FriendShipItemDTO
	for _, v := range list {
		var friendShipItemDTO dto.FriendShipItemDTO
		meUint, err := strconv.ParseUint(data.Me, 10, 64)
		if err != nil {
			return dto.GetListFriendShipOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Invalid Me ID")
		}
		if v.SenderID == uint(meUint) {
			friendShipItemDTO = dto.FriendShipItemDTO{
				ID:        v.ID,
				Friend:    dto.FriendDTO{ID: v.SenderID, Fullname: v.Sender.Profile.FullName, AvatarURL: v.Sender.Profile.AvatarURL},
				Status:    v.Status,
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
			}
		} else {
			friendShipItemDTO = dto.FriendShipItemDTO{
				ID:        v.ID,
				Friend:    dto.FriendDTO{ID: v.SenderID, Fullname: v.Sender.Profile.FullName, AvatarURL: v.Sender.Profile.AvatarURL},
				Status:    v.Status,
				CreatedAt: v.CreatedAt.Format(time.RFC3339),
			}
		}
		listFriendShipItemDTO = append(listFriendShipItemDTO, friendShipItemDTO)
	}

	meUint, err := strconv.ParseUint(data.Me, 10, 64)
	if err != nil {
		return dto.GetListFriendShipOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Invalid Me ID")
	}
	outDTO.Me = uint(meUint)
	outDTO.Data = listFriendShipItemDTO
	return outDTO, nil

}

// Update implements IFriendShipService.
func (f *friendShipService) Update(data dto.UpdateFriendShipInputDTO) (dto.UpdateFriendShipOutputDTO, error) {

	senderIDUint, err := strconv.ParseUint(data.SenderID, 10, 64)
	if err != nil {
		return dto.UpdateFriendShipOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Invalid sender ID")
	}
	receiverIDUint, err := strconv.ParseUint(data.ReceiverID, 10, 64)
	if err != nil {
		return dto.UpdateFriendShipOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Invalid receiver ID")
	}

	friendShipEntity, err := f.friendShipRepo.FindBySenderAndReceiver(uint(senderIDUint), uint(receiverIDUint))
	if err != nil {
		return dto.UpdateFriendShipOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "Friendship not found")
	}
	if data.Status == "ACCEPTED" {
		friendShipEntity.Status = entity.ACCEPTED
	} else if data.Status == "REJECTED" {
		friendShipEntity.Status = entity.REJECTED
	}

	_, err = f.friendShipRepo.Update(friendShipEntity)
	if err != nil {
		return dto.UpdateFriendShipOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to update friendship")
	}

	// tao 1 row converation o day
	var conversationEntity = entity.Conversation{
		IsGroup: false,
	}

	conversationCreated, err := f.conversationRepo.Create(conversationEntity)
	if err != nil {
		return dto.UpdateFriendShipOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to create conversation")
	}

	// toa 1 row để 2 người chat ne
	// 2 người sẽ giao tiếp qua cái conversationID hết á
	receiverIDUint, err = strconv.ParseUint(data.ReceiverID, 10, 64)
	if err != nil {
		return dto.UpdateFriendShipOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Invalid receiver ID")
	}
	var participantEntity1 = entity.Participant{
		AccountID:      uint(receiverIDUint),
		ConversationID: conversationCreated.ID,
		Role:           "member",
	}
	_, err = f.participantRepo.Create(participantEntity1)
	senderIDUint, err = strconv.ParseUint(data.SenderID, 10, 64)
	if err != nil {
		return dto.UpdateFriendShipOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Invalid sender ID")
	}
	var participantEntity2 = entity.Participant{
		AccountID:      uint(senderIDUint),
		ConversationID: conversationCreated.ID,
		Role:           "member",
	}
	_, err = f.participantRepo.Create(participantEntity2)
	if err != nil {
		return dto.UpdateFriendShipOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to create participant")
	}

	return dto.UpdateFriendShipOutputDTO{
		SenderID:   data.SenderID,
		ReceiverID: data.ReceiverID,
	}, nil
}

// Create implements IFriendShipService.
func (f *friendShipService) Create(data dto.CreateFriendShipInputDTO) (dto.CreateFriendShipOutputDTO, error) {
	// Convert SenderID and ReceiverID from string to uint
	senderIDUint, err := strconv.ParseUint(data.SenderID, 10, 64)
	if err != nil {
		return dto.CreateFriendShipOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Invalid sender ID")
	}
	receiverIDUint, err := strconv.ParseUint(data.ReceiverID, 10, 64)
	if err != nil {
		return dto.CreateFriendShipOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Invalid receiver ID")
	}

	_, err = f.friendShipRepo.FindBySenderAndReceiver(uint(senderIDUint), uint(receiverIDUint))
	if err == nil {
		return dto.CreateFriendShipOutputDTO{}, exception.NewCustomError(http.StatusConflict, "Friendship already exists")
	}
	// tìm ra account của 2 cái id senderId va reveicerId
	if _, err := f.accountRepo.GetUserByAccountId(data.SenderID); err != nil {
		return dto.CreateFriendShipOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "Not found sender account")
	}
	if _, err := f.accountRepo.GetUserByAccountId(data.ReceiverID); err != nil {
		return dto.CreateFriendShipOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "Not found receiver account")
	}

	var friendShipEntity = entity.FriendShip{
		SenderID:   uint(senderIDUint),
		ReceiverID: uint(receiverIDUint),
		Status:     entity.PENDING,
	}

	friendShipEntity, err = f.friendShipRepo.Create(friendShipEntity)
	if err != nil {
		return dto.CreateFriendShipOutputDTO{}, err
	}

	return dto.CreateFriendShipOutputDTO{
		SenderID:   strconv.FormatUint(uint64(friendShipEntity.SenderID), 10),
		ReceiverID: strconv.FormatUint(uint64(friendShipEntity.ReceiverID), 10),
		Status:     friendShipEntity.Status,
	}, nil
}

func NewFriendShipService(friendShipRepo reporitory.IFriendShipRepository, accountRepo reporitory.IAccountRepository, profileRepo reporitory.IProfileRepository, participantRepo reporitory.IParticipantRepository, conversationRepo reporitory.IConversationRepository, blockRepo reporitory.IBlockRepository) IFriendShipService {
	return &friendShipService{friendShipRepo: friendShipRepo, accountRepo: accountRepo, profileRepo: profileRepo, participantRepo: participantRepo, conversationRepo: conversationRepo, blockRepo: blockRepo}
}

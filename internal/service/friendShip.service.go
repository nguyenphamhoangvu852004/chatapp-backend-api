package service

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/entity"
	exception "chapapp-backend-api/internal/exeption"
	"chapapp-backend-api/internal/reporitory"
	"net/http"
	"strconv"
	"time"
)

type IFriendShipService interface {
	Create(data dto.CreateFriendShipInputDTO) (dto.CreateFriendShipOutputDTO, error)
	Update(data dto.UpdateFriendShipInputDTO) (dto.UpdateFriendShipOutputDTO, error)
	GetList(data dto.GetListFriendShipInputDTO) (dto.GetListFriendShipOutputDTO, error)
	GetListSendFriendShips(id string) (dto.GetFriendShipOutputDTO, error)
	GetListFriendShipsOfAccount(id string) (dto.GetFriendShipOutputDTO, error)
	GetListReceiveFriendShips(id string) (dto.GetFriendShipOutputDTO, error)
}

type friendShipService struct {
	friendShipRepo reporitory.IFriendShipRepository
	accountRepo    reporitory.IAccountRepository
	profileRepo    reporitory.IProfileRepository
}

// GetListReceiveFriendShips implements IFriendShipService.
func (s *friendShipService) GetListReceiveFriendShips(id string) (dto.GetFriendShipOutputDTO, error) {
	account, err := s.accountRepo.GetUserByAccountId(id)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "account not found")
	}

	senderAccounts, err := s.friendShipRepo.FindAllReceivedFriendRequests(id)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, err
	}

	var receivers []dto.Receiver
	for _, sender := range senderAccounts {
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

	// Tìm bạn bè
	friendAccounts, err := s.friendShipRepo.FindAllFriendOfAccount(id)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, err
	}

	var receivers []dto.Receiver
	for _, acc := range friendAccounts {
		receiver := dto.Receiver{
			ID:       acc.ID,
			Username: acc.Username,
			Email:    acc.Email,
			ImageURL: "",
			Status:   "ACCEPTED",
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
	account, err := f.accountRepo.GetUserByAccountId(id)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, exception.NewCustomError(http.StatusNotFound, "Not found account")
	}
	list, err := f.friendShipRepo.FindAllSendFriendShips(id)
	if err != nil {
		return dto.GetFriendShipOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to get list send friendship")
	}
	var listReceiver []dto.Receiver
	for i := range list {
		listReceiver = append(listReceiver, dto.Receiver{
			ID:       list[i].ReceiverID,
			Username: list[i].Receiver.Username,
			Email:    list[i].Receiver.Email,
			ImageURL: list[i].Receiver.Profile.AvatarURL,
			Status:   list[i].Status,
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

func NewFriendShipService(friendShipRepo reporitory.IFriendShipRepository, accountRepo reporitory.IAccountRepository, profileRepo reporitory.IProfileRepository) IFriendShipService {
	return &friendShipService{friendShipRepo: friendShipRepo, accountRepo: accountRepo, profileRepo: profileRepo}
}

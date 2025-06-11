package reporitory

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/entity"
	"strconv"
)

type IFriendShipRepository interface {
	FindAllFriendOfAccount(id string) ([]entity.Account, error)
	FindAllSendFriendShips(id string) ([]entity.FriendShip, error)
	GetList(data dto.GetListFriendShipInputDTO) ([]entity.FriendShip, error)
	Create(friendShip entity.FriendShip) (entity.FriendShip, error)
	Update(friendShip entity.FriendShip) (entity.FriendShip, error)
	GetByID(id uint) (entity.FriendShip, error)
	GetByAccountID(id uint) (entity.FriendShip, error)
	DeleteByID(id uint) (entity.FriendShip, error)
	FindBySenderAndReceiver(senderId uint, receiverId uint) (entity.FriendShip, error)
	FindAllReceivedFriendRequests(accountID string) ([]entity.Account, error) 
}

type friendShipRepository struct {
}
func (r *friendShipRepository) FindAllReceivedFriendRequests(accountID string) ([]entity.Account, error) {
	var friendShips []entity.FriendShip
	var senders []entity.Account

	err := global.Mdb.
		Preload("Sender.Profile").
		Where("receiver_id = ? AND status = ?", accountID, "PENDING").
		Find(&friendShips).Error
	if err != nil {
		return nil, err
	}

	for _, fs := range friendShips {
		senders = append(senders, fs.Sender)
	}

	return senders, nil
}

// FindAllFriendOfAccount implements IFriendShipRepository.
func (r *friendShipRepository) FindAllFriendOfAccount(accountID string) ([]entity.Account, error) {
	var friendShips []entity.FriendShip
	var friends []entity.Account

	err := global.Mdb.
		Preload("Sender.Profile").
		Preload("Receiver.Profile").
		Where("status = ? AND (sender_id = ? OR receiver_id = ?)", "ACCEPTED", accountID, accountID).
		Find(&friendShips).Error
	if err != nil {
		return nil, err
	}

	// Convert accountID from string to uint
	parsedAccountID, err := strconv.ParseUint(accountID, 10, 64)
	if err != nil {
		return nil, err
	}
	accountIDUint := uint(parsedAccountID)

	for _, fs := range friendShips {
		if fs.SenderID == accountIDUint {
			friends = append(friends, fs.Receiver)
		} else {
			friends = append(friends, fs.Sender)
		}
	}

	return friends, nil
}



// FindAllSendFriendShips implements IFriendShipRepository.
func (r *friendShipRepository) FindAllSendFriendShips(id string) ([]entity.FriendShip, error) {
	var friendShips []entity.FriendShip

	// Convert string id to uint
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	idUint := uint(parsedId)

	err = global.Mdb.
		Preload("Receiver").         // Preload Receiver (Account)
		Preload("Receiver.Profile"). // Preload Profile trong Receiver
		Where("sender_id = ? AND status = ?", idUint, "PENDING").
		Find(&friendShips).Error

	if err != nil {
		return nil, err
	}
	return friendShips, nil
}

// GetList implements IFriendShipRepository.
func (p *friendShipRepository) GetList(data dto.GetListFriendShipInputDTO) ([]entity.FriendShip, error) {
	var friendShipList []entity.FriendShip
	query := global.Mdb.Preload("Sender").Preload("Receiver").Preload("Sender.Profile").Preload("Receiver.Profile").Model(&entity.FriendShip{})

	if data.Status != "" {
		query = query.Where("status = ?", data.Status)
	}

	if data.Me != "" {
		query = query.Where("sender_id = ? or receiver_id = ?", data.Me, data.Me)
	}

	result := query.Find(&friendShipList)
	if result.Error != nil {
		return []entity.FriendShip{}, result.Error
	}
	return friendShipList, nil
}

// FindBySenderAndReceiver implements IFriendShipRepository.
func (p *friendShipRepository) FindBySenderAndReceiver(senderId uint, receiverId uint) (entity.FriendShip, error) {
	var friendShip entity.FriendShip
	result := global.Mdb.Where("sender_id = ? AND receiver_id = ?", senderId, receiverId).First(&friendShip)
	if result.Error != nil {
		return entity.FriendShip{}, result.Error
	}
	return friendShip, nil
}

// Create implements IFriendShipRepository.
func (p *friendShipRepository) Create(profile entity.FriendShip) (entity.FriendShip, error) {
	result := global.Mdb.Create(&profile)
	if result.Error != nil {
		return entity.FriendShip{}, result.Error
	}
	return profile, nil
}

// DeleteByID implements IFriendShipRepository.
func (p *friendShipRepository) DeleteByID(id uint) (entity.FriendShip, error) {
	panic("unimplemented")
}

// GetByAccountID implements IFriendShipRepository.
func (p *friendShipRepository) GetByAccountID(id uint) (entity.FriendShip, error) {
	panic("unimplemented")
}

// GetByID implements IFriendShipRepository.
func (p *friendShipRepository) GetByID(id uint) (entity.FriendShip, error) {
	var profile entity.FriendShip
	err := global.Mdb.First(&profile, id).Error
	if err != nil {
		return entity.FriendShip{}, err
	}
	return profile, nil
}

// Update implements IFriendShipRepository.
func (p *friendShipRepository) Update(profile entity.FriendShip) (entity.FriendShip, error) {
	result := global.Mdb.Save(&profile)
	if result.Error != nil {
		return entity.FriendShip{}, result.Error
	}
	return profile, nil
}

func NewFriendShipRepository() IFriendShipRepository {
	return &friendShipRepository{}
}

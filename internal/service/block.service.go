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

type IBlockService interface {
	IsBlocked(id1 string, id2 string) (bool, error)
	Create(data dto.CreateBlockInputDTO) (dto.CreateBlockOutputDTO, error)
	Delete(data dto.DeleteBlockInputDTO) (dto.DeleteBlockOutputDTO, error)
	GetList(data string) (dto.GetBlockListOutputDTO, error)
}

type blockService struct {
	blockRepo      reporitory.IBlockRepository
	friendShipRepo reporitory.IFriendShipRepository
}

// GetList implements IBlockService.
func (b *blockService) GetList(data string) (dto.GetBlockListOutputDTO, error) {
	var outputDTO dto.GetBlockListOutputDTO

	userID, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		return dto.GetBlockListOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Invalid user ID")
	}

	// Lấy danh sách các block entity mà user là blocker
	blocks, err := b.blockRepo.GetListBlocker(uint(userID))
	if err != nil {
		return dto.GetBlockListOutputDTO{}, err
	}

	// Nếu không có block nào thì return rỗng
	if len(blocks) == 0 {
		return outputDTO, nil
	}

	// Gán thông tin bản thân từ Blocker (lấy từ phần tử đầu tiên)
	blocker := blocks[0].Blocker
	outputDTO.Me = dto.GetAccountDetailOutputDTO{
		Id:          fmt.Sprintf("%d", blocker.ID),
		Username:    blocker.Username,
		Email:       blocker.Email,
		PhoneNumber: blocker.PhoneNumber,
		Profile: dto.GetProfileDetailOutputDTO{
			Id:        fmt.Sprintf("%d", blocker.Profile.ID),
			FullName:  blocker.Profile.FullName,
			Bio:       blocker.Profile.Bio,
			AvatarURL: blocker.Profile.AvatarURL,
			CoverURL:  blocker.Profile.CoverURL,
			CreatedAt: blocker.Profile.CreatedAt.Format(time.RFC3339),
			UpdatedAt: blocker.Profile.UpdatedAt.Format(time.RFC3339),
		},
	}

	// Gán danh sách những người bị block
	for _, block := range blocks {
		blocked := block.Blocked
		outputDTO.BlockedList = append(outputDTO.BlockedList, dto.GetAccountDetailOutputDTO{
			Id:          fmt.Sprintf("%d", blocked.ID),
			Username:    blocked.Username,
			Email:       blocked.Email,
			PhoneNumber: blocked.PhoneNumber,
			Profile: dto.GetProfileDetailOutputDTO{
				Id:        fmt.Sprintf("%d", blocked.Profile.ID),
				FullName:  blocked.Profile.FullName,
				Bio:       blocked.Profile.Bio,
				AvatarURL: blocked.Profile.AvatarURL,
				CoverURL:  blocked.Profile.CoverURL,
				CreatedAt: blocked.Profile.CreatedAt.Format(time.RFC3339),
				UpdatedAt: blocked.Profile.UpdatedAt.Format(time.RFC3339),
			},
		})
	}

	return outputDTO, nil
}

// Delete implements IBlockService.
func (b *blockService) Delete(data dto.DeleteBlockInputDTO) (dto.DeleteBlockOutputDTO, error) {
	// kiểm tra coi tụi nó có bị block thật không, không có thì return khoong ton tai
	checkFlag, _ := b.IsBlocked(data.BlockerId, data.BlockedId)
	if checkFlag == false {
		return dto.DeleteBlockOutputDTO{BlockerId: data.BlockerId, BlockedId: data.BlockedId, IsSuccess: false}, exception.NewCustomError(http.StatusNotFound, "Block not found")
	}

	// Convert string IDs to uint
	blockerID, err := strconv.ParseUint(data.BlockerId, 10, 64)
	if err != nil {
		return dto.DeleteBlockOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Invalid BlockerID")
	}
	blockedID, err := strconv.ParseUint(data.BlockedId, 10, 64)
	if err != nil {
		return dto.DeleteBlockOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Invalid BlockedID")
	}
	blockDeleted, err := b.blockRepo.DeleteBlock(entity.Block{
		BlockerID: uint(blockerID),
		BlockedID: uint(blockedID),
	})
	if err != nil {
		return dto.DeleteBlockOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to delete block")
	}
	return dto.DeleteBlockOutputDTO{
		BlockerId: strconv.FormatUint(uint64(blockDeleted.BlockerID), 10),
		BlockedId: strconv.FormatUint(uint64(blockDeleted.BlockedID), 10),
		IsSuccess: true,
	}, nil
}

// Create implements IBlockService.
func (b *blockService) Create(data dto.CreateBlockInputDTO) (dto.CreateBlockOutputDTO, error) {
	// Kiểm tra đã block nhau chưa
	checkedFlag, _ := b.IsBlocked(data.BlockerId, data.BlockedId)
	if checkedFlag {
		return dto.CreateBlockOutputDTO{}, exception.NewCustomError(http.StatusConflict, "Users are already blocked")
	}

	// Convert string -> uint
	blockerID, err := strconv.ParseUint(data.BlockerId, 10, 64)
	if err != nil {
		return dto.CreateBlockOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Invalid BlockerID")
	}
	blockedID, err := strconv.ParseUint(data.BlockedId, 10, 64)
	if err != nil {
		return dto.CreateBlockOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Invalid BlockedID")
	}

	newBlock := entity.Block{
		BlockerID: uint(blockerID),
		BlockedID: uint(blockedID),
	}

	// Kiểm tra quan hệ bạn bè (2 chiều)
	friendEntity, err := b.friendShipRepo.FindBySenderAndReceiver(uint(blockerID), uint(blockedID))
	if err != nil {
		friendEntity, err = b.friendShipRepo.FindBySenderAndReceiver(uint(blockedID), uint(blockerID))
	}

	// Nếu tìm thấy, thì xoá mối quan hệ bạn bè
	if err == nil {
		_,err = b.friendShipRepo.DeleteByID(friendEntity.ID)
		if err != nil {
			return dto.CreateBlockOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to delete friendship")
		}
	}

	// Tạo bản ghi block mới
	blockCreated, err := b.blockRepo.CreateBlock(newBlock)
	if err != nil {
		return dto.CreateBlockOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to create block")
	}

	return dto.CreateBlockOutputDTO{
		BlockerId: strconv.FormatUint(uint64(blockCreated.BlockerID), 10),
		BlockedId: strconv.FormatUint(uint64(blockCreated.BlockedID), 10),
		IsSuccess: true,
	}, nil
}

// IsBlocked implements IBlockService.
func (b *blockService) IsBlocked(id1 string, id2 string) (bool, error) {
	blockerID, err := strconv.ParseUint(id1, 10, 64)
	if err != nil {
		return false, exception.NewCustomError(http.StatusBadRequest, "Invalid BlockerID")
	}
	blockedID, err := strconv.ParseUint(id2, 10, 64)
	if err != nil {
		return false, exception.NewCustomError(http.StatusBadRequest, "Invalid BlockedID")
	}
	return b.blockRepo.IsBlocked(uint(blockerID), uint(blockedID))
}

func NewBlockService(blockRepo reporitory.IBlockRepository, friendShipRepo reporitory.IFriendShipRepository) IBlockService {
	return &blockService{
		blockRepo:      blockRepo,
		friendShipRepo: friendShipRepo,
	}
}

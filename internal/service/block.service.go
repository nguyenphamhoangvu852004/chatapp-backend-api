package service

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/entity"
	exception "chapapp-backend-api/internal/exeption"
	"chapapp-backend-api/internal/reporitory"
	"net/http"
	"strconv"
)

type IBlockService interface {
	IsBlocked(id1 string, id2 string) (bool, error)
	Create(data dto.CreateBlockInputDTO) (dto.CreateBlockOutputDTO, error)
	Delete(data dto.DeleteBlockInputDTO) (dto.DeleteBlockOutputDTO, error)
}

type blockService struct {
	blockRepo reporitory.IBlockRepository
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
	// kiem tra coi 2 thang do block nhau chua
	checkedFlag, _ := b.IsBlocked(data.BlockerId, data.BlockedId)
	if checkedFlag == true {
		return dto.CreateBlockOutputDTO{}, exception.NewCustomError(http.StatusConflict, "Users are already blocked")
	} else {
		// Convert string IDs to uint
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

func NewBlockService(blockRepo reporitory.IBlockRepository) IBlockService {
	return &blockService{
		blockRepo: blockRepo,
	}
}

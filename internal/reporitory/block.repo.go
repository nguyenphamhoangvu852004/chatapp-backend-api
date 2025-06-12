package reporitory

import (
	"chapapp-backend-api/global"
	"chapapp-backend-api/internal/entity"
)

type IBlockRepository interface {
	IsBlocked(blockerId uint, blockedId uint) (bool, error)
	CreateBlock(block entity.Block) (entity.Block, error)
	DeleteBlock(block entity.Block) (entity.Block, error)
	GetListBlocked(accountID uint) ([]entity.Block, error)
	GetListBlocker(accountID uint) ([]entity.Block, error)
}

type blockRepository struct{}

// CreateBlock implements IBlockRepository.
func (b *blockRepository) CreateBlock(block entity.Block) (entity.Block, error) {
	result := global.Mdb.Create(&block)
	if result.Error != nil {
		return entity.Block{}, result.Error
	}
	return block, nil
}

// DeleteBlock implements IBlockRepository.
func (b *blockRepository) DeleteBlock(block entity.Block) (entity.Block, error) {
	result := global.Mdb.Unscoped().Where("blocker_id = ? AND blocked_id = ?", block.BlockerID, block.BlockedID).Delete(&block)
	if result.Error != nil {
		return entity.Block{}, result.Error
	}
	return block, nil
}

// GetListBlocked implements IBlockRepository.
func (b *blockRepository) GetListBlocked(accountID uint) ([]entity.Block, error) {
	var blocks []entity.Block
	result := global.Mdb.Where("blocked_id = ?", accountID).Find(&blocks)
	if result.Error != nil {
		return nil, result.Error
	}
	return blocks, nil
}

// GetListBlocker implements IBlockRepository.
func (b *blockRepository) GetListBlocker(accountID uint) ([]entity.Block, error) {
	var blockers []entity.Block
	result := global.Mdb.Where("blocker_id = ?", accountID).Find(&blockers)
	if result.Error != nil {
		return nil, result.Error
	}
	return blockers, nil
}

// IsBlocked implements IBlockRepository.
func (b *blockRepository) IsBlocked(blockerID uint, blockedID uint) (bool, error) {
	var block entity.Block
	result := global.Mdb.Where("blocker_id = ? AND blocked_id = ?", blockerID, blockedID).First(&block)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func NewBlockRepository() IBlockRepository {
	return &blockRepository{}
}

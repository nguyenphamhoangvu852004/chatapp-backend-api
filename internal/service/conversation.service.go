package service

import (
	"chapapp-backend-api/internal/dto"
	"chapapp-backend-api/internal/entity"
	exception "chapapp-backend-api/internal/exeption"
	"chapapp-backend-api/internal/reporitory"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type IConversationSerivce interface {
	Create(data dto.CreateConversationInputDTO) (dto.CreateConversationOutputDTO, error)
	GetGroupListWhereUserIsAdmin(accountId string) ([]dto.GetConversationOutputDTO, error)
	AddMembers(data dto.AddMemberInputDTO) (dto.AddMemberOutputDTO, error)
	Delete(data dto.DeleteMessageGroupInputDTO) (dto.DeleteMessageGroupOutputDTO, error)
	RemoveMembers(data dto.RemoveMembersInputDTO) (dto.RemoveMembersOutputDTO, error)
	GetGroupsJoinedByMe(accountId string) ([]dto.GetJoinedGroupsOutputDTO, error)
}

type conversationService struct {
	conversationRepo reporitory.IConversationRepository

	participantRepo reporitory.IParticipantRepository
}

func (s *conversationService) GetGroupsJoinedByMe(accountId string) ([]dto.GetJoinedGroupsOutputDTO, error) {
	uid, err := strconv.ParseUint(accountId, 10, 64)
	if err != nil {
		return nil, errors.New("invalid account ID")
	}

	participants, err := s.participantRepo.FindGroupsByAccountId(uint(uid))
	if err != nil {
		return nil, err
	}

	var result []dto.GetJoinedGroupsOutputDTO
	for _, p := range participants {
		// Lấy tất cả members của conversation này
		members, err := s.participantRepo.FindMembersByConversationID(p.ConversationID)
		if err != nil {
			continue // bỏ qua nếu lỗi
		}

		var memberDTOs []dto.Member
		for _, m := range members {
			memberDTOs = append(memberDTOs, dto.Member{
				AccountId:   fmt.Sprintf("%d", m.Account.ID),
				Email:       m.Account.Email,
				PhoneNumber: m.Account.PhoneNumber,
				Username:    m.Account.Username,
				AvatarURL:   m.Account.Profile.AvatarURL,
				Role:        m.Role,
			})
		}

		result = append(result, dto.GetJoinedGroupsOutputDTO{
			ConversationId: fmt.Sprintf("%d", p.Conversation.ID),
			Name:           p.Conversation.Name,
			GroupAvatar:    p.Conversation.GroupAvatar,
			IsGroup:        p.Conversation.IsGroup,
			Role:           p.Role,
			Members:        memberDTOs,
		})
	}
	return result, nil
}

// RemoveMembers implements IConversationSerivce.
func (s *conversationService) RemoveMembers(data dto.RemoveMembersInputDTO) (dto.RemoveMembersOutputDTO, error) {
	// Parse conversation ID và owner ID
	conversationIdUint, err := strconv.ParseUint(data.ConversationId, 10, 64)
	if err != nil {
		return dto.RemoveMembersOutputDTO{IsSuccess: false}, errors.New("invalid conversation ID")
	}
	ownerIdUint, err := strconv.ParseUint(data.OwnerId, 10, 64)
	if err != nil {
		return dto.RemoveMembersOutputDTO{IsSuccess: false}, errors.New("invalid owner ID")
	}

	// 1. Kiểm tra xem conversation có tồn tại không
	conversation, err := s.conversationRepo.FindById(uint(conversationIdUint))
	if err != nil || conversation.ID == 0 {
		return dto.RemoveMembersOutputDTO{IsSuccess: false}, errors.New("conversation not found")
	}

	// 2. Kiểm tra owner có phải admin không
	isAdmin, err := s.participantRepo.CheckIsAdmin(uint(ownerIdUint), uint(conversationIdUint))
	if err != nil {
		return dto.RemoveMembersOutputDTO{IsSuccess: false}, err
	}
	if !isAdmin {
		return dto.RemoveMembersOutputDTO{IsSuccess: false}, errors.New("you are not admin of this group")
	}

	// 3. Convert member IDs từ []string → []uint
	var memberIDs []uint
	for _, idStr := range data.Ids {
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err == nil {
			memberIDs = append(memberIDs, uint(id))
		}
	}

	if len(memberIDs) == 0 {
		return dto.RemoveMembersOutputDTO{IsSuccess: false}, errors.New("no valid member IDs")
	}

	// 4. Xoá các participant khỏi nhóm
	err = s.participantRepo.DeleteMany(uint(conversationIdUint), memberIDs)
	if err != nil {
		return dto.RemoveMembersOutputDTO{IsSuccess: false}, err
	}

	return dto.RemoveMembersOutputDTO{
		ConversationId: data.ConversationId,
		OwnerId:        data.OwnerId,
		Ids:            data.Ids,
		IsSuccess:      true,
	}, nil
}

// Delete implements IConversationSerivce.
// Delete implements IConversationSerivce.
func (s *conversationService) Delete(data dto.DeleteMessageGroupInputDTO) (dto.DeleteMessageGroupOutputDTO, error) {
	// Parse conversation ID và owner ID từ string sang uint
	conversationIdUint, err := strconv.ParseUint(data.ConversationId, 10, 64)
	if err != nil {
		return dto.DeleteMessageGroupOutputDTO{IsSuccess: false}, exception.NewCustomError(http.StatusBadRequest, "invalid conversation ID")
	}

	ownerIdUint, err := strconv.ParseUint(data.OwnerId, 10, 64)
	if err != nil {
		return dto.DeleteMessageGroupOutputDTO{IsSuccess: false}, exception.NewCustomError(http.StatusBadRequest, "invalid owner ID")
	}

	// 1. Kiểm tra xem conversation có tồn tại không
	conversation, err := s.conversationRepo.FindById(uint(conversationIdUint))
	if err != nil || conversation.ID == 0 {
		return dto.DeleteMessageGroupOutputDTO{IsSuccess: false}, exception.NewCustomError(http.StatusNotFound, "conversation not found")
	}

	// 2. Kiểm tra owner có phải admin trong conversation không
	isAdmin, err := s.participantRepo.CheckIsAdmin(uint(ownerIdUint), uint(conversationIdUint))
	if err != nil {
		return dto.DeleteMessageGroupOutputDTO{IsSuccess: false}, err
	}
	if !isAdmin {
		return dto.DeleteMessageGroupOutputDTO{IsSuccess: false}, exception.NewCustomError(http.StatusBadRequest, "you are not admin of this group")
	}

	// 3. Xoá conversation (các participant sẽ tự động bị xoá do CASCADE)
	err = s.conversationRepo.DeleteById(uint(conversationIdUint))
	if err != nil {
		return dto.DeleteMessageGroupOutputDTO{IsSuccess: false}, err
	}

	// 4. Trả về kết quả
	return dto.DeleteMessageGroupOutputDTO{
		ConversationId: data.ConversationId,
		OwnerId:        data.OwnerId,
		IsSuccess:      true,
	}, nil
}

// AddMember implements IConversationSerivce.
func (s *conversationService) AddMembers(data dto.AddMemberInputDTO) (dto.AddMemberOutputDTO, error) {
	// Parse conversation ID và owner ID từ string sang uint
	conversationIdUint, err := strconv.ParseUint(data.ConversationId, 10, 64)
	if err != nil {
		return dto.AddMemberOutputDTO{IsSuccess: false}, errors.New("invalid conversation ID")
	}

	ownerIdUint, err := strconv.ParseUint(data.OwnerId, 10, 64)
	if err != nil {
		return dto.AddMemberOutputDTO{IsSuccess: false}, errors.New("invalid owner ID")
	}

	// 1. Kiểm tra xem conversation có tồn tại không
	conversation, err := s.conversationRepo.FindById(uint(conversationIdUint))
	if err != nil || conversation.ID == 0 {
		return dto.AddMemberOutputDTO{IsSuccess: false}, errors.New("conversation not found")
	}

	// 2. Kiểm tra owner có phải admin trong conversation không
	isAdmin, err := s.participantRepo.CheckIsAdmin(uint(ownerIdUint), uint(conversationIdUint))
	if err != nil {
		return dto.AddMemberOutputDTO{IsSuccess: false}, err
	}
	if !isAdmin {
		return dto.AddMemberOutputDTO{IsSuccess: false}, errors.New("you are not admin of this group")
	}

	// 3. Tạo danh sách participant cần thêm
	var participants []entity.Participant
	for _, idStr := range data.MemberIds {
		memberIdUint, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			continue // skip nếu lỗi convert ID
		}
		participants = append(participants, entity.Participant{
			AccountID:      uint(memberIdUint),
			ConversationID: uint(conversationIdUint),
			Role:           "member",
		})
	}

	// 4. Thêm vào DB (bulk insert)
	err = s.participantRepo.AddMembers(participants)
	if err != nil {
		return dto.AddMemberOutputDTO{IsSuccess: false}, err
	}

	// 5. Trả về kết quả
	return dto.AddMemberOutputDTO{
		ConversationId: data.ConversationId,
		OwnerId:        data.OwnerId,
		MemberIds:      data.MemberIds,
		IsSuccess:      true,
	}, nil
}

func (s *conversationService) GetGroupListWhereUserIsAdmin(accountId string) ([]dto.GetConversationOutputDTO, error) {
	participants, err := s.participantRepo.GetGroupListWhereUserIsAdmin(accountId)
	if err != nil {
		return nil, err
	}

	var results []dto.GetConversationOutputDTO
	for _, participant := range participants {
		conversation := participant.Conversation

		var members []dto.Member
		for _, p := range conversation.Participants {
			account := p.Account
			members = append(members, dto.Member{
				AccountId:   fmt.Sprint(account.ID),
				Email:       account.Email,
				PhoneNumber: account.PhoneNumber,
				Username:    account.Username,
				AvatarURL:   account.Profile.AvatarURL,
				Role:        p.Role,
			})
		}

		results = append(results, dto.GetConversationOutputDTO{
			ConversationId: fmt.Sprint(conversation.ID),
			OwnerID:        fmt.Sprint(participant.AccountID),
			Name:           conversation.Name,
			GroupAvatar:    conversation.GroupAvatar,
			IsGroup:        conversation.IsGroup,
			Members:        members,
		})
	}

	return results, nil
}

// GetListOwnedByMe implements IConversationSerivce.

// Create implements IConversationSerivce.
func (c *conversationService) Create(data dto.CreateConversationInputDTO) (dto.CreateConversationOutputDTO, error) {
	var conversationEntity = entity.Conversation{
		Name:        data.Name,
		IsGroup:     true,
		GroupAvatar: data.AvatarURL,
	}
	conversation, err := c.conversationRepo.Create(conversationEntity)
	if err != nil {
		return dto.CreateConversationOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to create conversation")
	}
	conversationId := conversation.ID
	// thêm cái người owner vào
	// Convert OwnerId from string to uint
	ownerIdUint, err := strconv.ParseUint(data.OwnerId, 10, 64)
	if err != nil {
		return dto.CreateConversationOutputDTO{}, exception.NewCustomError(http.StatusBadRequest, "Invalid owner ID")
	}
	var participantEntity = entity.Participant{
		ConversationID: conversationId,
		AccountID:      uint(ownerIdUint),
		Role:           "admin",
	}
	participant, err := c.participantRepo.Create(participantEntity)
	if err != nil {
		return dto.CreateConversationOutputDTO{}, exception.NewCustomError(http.StatusInternalServerError, "Failed to create participant")
	}
	return dto.CreateConversationOutputDTO{
		ConversationId: strconv.FormatUint(uint64(conversationId), 10),
		Name:           conversation.Name,
		IsGroup:        conversation.IsGroup,
		OwnerId:        strconv.FormatUint(uint64(participant.AccountID), 10),
		GroupAvatar:    conversation.GroupAvatar,
	}, nil
}

func NewConversationService(conversationRepo reporitory.IConversationRepository, participateRepo reporitory.IParticipantRepository) IConversationSerivce {
	return &conversationService{conversationRepo: conversationRepo, participantRepo: participateRepo}
}

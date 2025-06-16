package dto

type (
	CreateConversationInputDTO struct {
		Name      string `json:"name"`
		AvatarURL string `json:"avatarUrl"`
		OwnerId   string `json:"ownerId"`
	}
	CreateConversationOutputDTO struct {
		OwnerId        string `json:"ownerId"`
		ConversationId string `json:"conversationId"`
		Name           string `json:"name"`
		IsGroup        bool   `json:"isGroup"`
		GroupAvatar    string `json:"groupAvatar"`
	}
)

type (
	GetConversationOutputDTO struct {
		ConversationId string   `json:"conversationId"`
		OwnerID        string   `json:"ownerId"`
		Name           string   `json:"name"`
		GroupAvatar    string   `json:"groupAvatar"`
		IsGroup        bool     `json:"isGroup"`
		Members        []Member `json:"members"`
	}
	Member struct {
		AccountId   string `json:"accountId"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phoneNumber"`
		Username    string `json:"username"`
		AvatarURL   string `json:"avatarUrl"`
		Role        string `json:"role"`
	}
)

type (
	AddMemberInputDTO struct {
		ConversationId string   `json:"conversationId"`
		OwnerId        string   `json:"ownerId"`
		MemberIds      []string `json:"ids"`
	}

	AddMemberOutputDTO struct {
		ConversationId string   `json:"conversationId"`
		OwnerId        string   `json:"ownerId"`
		MemberIds      []string `json:"ids"`
		IsSuccess      bool     `json:"isSuccess"`
	}
)

type (
	DeleteMessageGroupInputDTO struct {
		ConversationId string `json:"conversationId"`
		OwnerId        string `json:"ownerId"`
	}

	DeleteMessageGroupOutputDTO struct {
		ConversationId string `json:"conversationId"`
		OwnerId        string `json:"ownerId"`
		IsSuccess      bool   `json:"isSuccess"`
	}
)

type (
	RemoveMembersInputDTO struct {
		ConversationId string   `json:"conversationId"`
		OwnerId        string   `json:"ownerId"`
		Ids            []string `json:"ids"`
	}

	RemoveMembersOutputDTO struct {
		ConversationId string   `json:"conversationId"`
		OwnerId        string   `json:"ownerId"`
		Ids            []string `json:"ids"`
		IsSuccess      bool     `json:"isSuccess"`
	}
)

type (
	GetJoinedGroupsOutputDTO struct {
		ConversationId string   `json:"conversationId"`
		Name           string   `json:"name"`
		GroupAvatar    string   `json:"groupAvatar"`
		IsGroup        bool     `json:"isGroup"`
		Role           string   `json:"role"` // 'admin' hoặc 'member'
		Members        []Member `json:"members"`
	}
)

type (
	ModifyConversationInputDTO struct {
		OwnerId        string  `form:"ownerId"`
		ConversationId string  `form:"conversationId"`
		Name           *string `form:"name,omitempty"` // optional
		AvatarURL      *string `form:"-"`
	}

	ModifyConversationOutputDTO struct {
		ConversationId string `json:"conversationId"`
		OwnerId        string `json:"ownerId"`
		Name           string `json:"name"`
		AvatarURL      string `json:"avatarUrl"`
		IsSuccess      bool   `json:"isSuccess"`
	}
)

type (
	GetListMembersOuputDTO struct {
		ConversationId string   `json:"conversationId"`
		Members        []Member `json:"members"`
	}
)

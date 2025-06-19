package dto

type (
	CreateMessageInputDTO struct {
		SenderId       string  `json:"senderId"`
		ConversationId string  `json:"conversationId"`
		Content        string  `json:"content"`
		OriginFilename *string `json:"originFilename"`
		Size           *string `json:"size"`
	}
	CreateMessageOutputDTO struct {
		MessageId      string  `json:"messageId"`
		Content        string  `json:"content"`
		OriginFilename *string `json:"originFilename"`
		Size           *string `json:"size"`
	}
)

type (
	GetListMessageInputDTO struct {
		Me             string `form:"me"`
		ConversationId string `form:"conversationId"`
	}
	GetListMessageOutputDTO struct {
		ConversationId string    `json:"conversationId"`
		Messages       []Message `json:"messages"`
	}
	Message struct {
		ID             string  `json:"id"`
		SenderId       string  `json:"senderId"`
		Content        string  `json:"content"`
		Type           string  `json:"type"`
		OriginFilename *string `json:"originFilename"`
		Size           *string `json:"size"`
		CreatedAt      string  `json:"createdAt"`
		IsDeleted      bool    `json:"isDeleted"`
	}
)

type (
	DeleteMessageInputDTO struct {
		OwnerId        string `json:"ownerId"`
		ConversationId string `json:"conversationId"`
		MessageId      string `json:"messageId"`
	}

	DeleteMessageOutputDTO struct {
		OwnerId        string `json:"ownerId"`
		ConversationId string `json:"conversationId"`
		MessageId      string `json:"messageId"`
		IsDeleted      bool   `json:"isDeleted"`
	}
)
